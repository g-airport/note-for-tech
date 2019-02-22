package main

import "C"
import "unsafe"

//-------------------------------------------------------------------*
// go调用c 阻塞 后 导致线程 无法释放，出现递增
// top -pid pid 查看线程情况
//-------------------------------------------------------------------*



//runtime handoffp
//1.主动解绑
//2.非正常解绑 runtime sysmon retake (preempted)
//case : block io
//retake preempt (exception handoffp) -> thread increase promptly

////1.rawsyscall是调用压根不会阻塞的系统调用，比如getpid, getuid, time。
//因为vdso机制，直接把调用打到你的进程方法映射空间里
////2.runtime调度的syscall, 这里的syscall又分为enterSyscall和enterSyscallBlock
////3.go把mutex锁抽象成可cas + waitqueue + futex的阻塞, entersyscallBlock因为知道
//可能会阻塞，所以直接就handoffp

//调用handoffp会寻找空闲线程，如果没有就创建新线程
//sysmon线程会循环检测syscall阻塞，并发起解绑抢占

/*
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
void output(char *str) {
    usleep(1000000);
    printf("%s\n", str);
}
*/
import "C"

func main() {
	for i := 0; i < 50; i++ {
		go func() {
			str := "hello cgo"
			//change to char*s
			cStr := C.CString(str)
			C.output(cStr)
			C.free(unsafe.Pointer(cStr))

		}()
	}
	select {}
}





//--------------------------------------------------------------------*
////golang mutex bug : os_linux.go
//--------------------------------------------------------------------*

//futex 描述

//#include <linux/futex.h>
//
//#include <sys/time.h>
//
//int futex(int *uaddr, int futex_op, int val,
//
//　　　　　 const struct timespec *timeout, /* or: uint32_t val2 */
//
//　　　　　 int *uaddr2, int val3);

//futex()系统调用提供了一种方法用于等待某个特定条件的发生。
//一种典型的应用是在共享内存同步中作为阻塞装置。当使用futex
//时，绝大部分同步操作都在用户态完成，一个用户态程序只有在有
//可能阻塞一段较长的时间等待条件的发生时才使用futex。其他的
//futex操作可以用来唤醒等待特定条件发生的任何进程或线程。
//一个futex是一个32位的值，它的地址通过futex()传入，futexs
//在所有的平台上都是32位的(包括64位系统上)，所有的futex操作都
//是在这个值上。为了在多个进程间共享futex，futex位于由mmap或
//者shmat创建的一块共享内存上(这种情况下，futex值在不同的进程
//上可能位于不同的虚拟地址，但这些虚拟地址都指向相同的物理地址)。
//在多线程程序中，将futex值放到一个所有线程共享的全局变量中就
//可以了。当执行一个futex操作请求阻塞一个线程时，只有在 *uaddr == val
//时，内核才会执行阻塞操作，这个操作中的所有步骤：

//1. 导入*uaddr的值;
//2.比较;
//3.阻塞线程

//将原子的执行，并且当其他线程在同一个futex值上并行
//操作时所有的步骤不会乱序。
//futex的一种使用方式是用来实现锁，锁的状态(acquired or not acquired)
//可以用在共享内存中的原子标志表示。在非竞争的情况下，线程可以通过原子操作访问
//和修改锁的状态(这些操作全部在用户态操作，内核不会保存任何关于锁的状态)。
//从另一方面来说，另一个线程可能无法获取锁(因为锁已经被某个线程获取)，
//这个线程将通过如下这种方式执行futex()等待操作：


//atomic<int> lock; // lock ： 0. 锁未被获取 1.锁被获取
//futex(&lock, FUTEX_WAIT, 1, NULL, NULL, NULL);
//futex(FUTEX_WAIT)将会检测lock的值，只有在等于1时才阻塞线程。
//当线程释放锁时，该线程必须首先重置锁的状态，然后执行futex操作唤醒阻塞在lock上的线程。



//注意使用futex并没有显示的初始化和销毁操作，
//内核只有在一个指定的futex值上执行futex操作时(例如FUTEX_WAIT)才会维护futex数据。

//futex 参数

//uaddr指向futex值，在所有的平台上，futex值是一个4字节的整数并且必须4字节对齐。
//对于某些阻塞操作，timeout参数是一个指向timespec结构的指针，表明了操作的超时时间。
//然而，在其他的某些操作下，它的最低4字节被作为一个整数值，这个整数值的含义因futex
//操作的不同而不同，对于这些操作来说，内核会将timeout值转换为unsigned long，
//然后转换为uint32_t，在接下来的说明中，它将表示为val2。
//在需要的时候，uaddr2是指向第二个futex值的指针，val3的解释将依赖于具体的操作。


//futex 操作

//futex_op参数包含两个部分：
//1.操作类型
//2.标志选项(影响操作行为)。

//FUTEX_PRIVATE_FLAG (since Linux 2.6.22)
//这个标志选项可以用在所有的futex操作中，它告诉内核这个futex是进程内的，
//不和其他进程共享(它只能用来同步同一个进程内的线程)，这样内核可以做一些
//额外的优化

//FUTEX_CLOCK_REALTIME (since Linux 2.6.28)
//这个标志选项只能用在FUTEX_WAIT_BITSET 和 FUTEX_WAIT_REQUEUE_PI
//操作中，如果设置了这个标志选项，内核会将timeout视作基于CLOCK_REALTIME
//的绝对时间。如果没有设置这个标志，则内核将它视作基于CLOCK_MONOTONIC的相对时间。



//futex操作包含下面的几种:

//FUTEX_WAIT (since Linux 2.6.0)

//本操作会检查*uaddr是否等于val，如果相等的话，则睡眠等待在uaddr上的FUTEX_WAKE
//操作，如果线程开始睡眠，则它被认为是这个futex值上一个等待者。如果两者不相等，
//则操作失败并返回 error EAGAIN。比较*uaddr和val的目的是为了避免丢失唤醒。
//如果timeout参数不为NULL，则它指定了等待的超时时间(根据CLOCK_MONOTONIC测量得出)。



//FUTEX_WAKE (since Linux 2.6.0)

//本操作会唤醒最多val个等待者，绝大部分情况下，val的值为1(只唤醒一个等待者)或
//INT_MAX(唤醒所有的等待者) 。注意没有任何机制保证一定唤醒某些特定的等待者
//(比如被认为是高优先级的等待者)
//参数timeout(ts timespec), uaddr2, val3将被忽略。



//FUTEX_REQUEUE (since Linux 2.6.0)

//本操作和下面的 FUTEX_CMP_REQUEUE 完成相同的功能，除了不检查val3 。(参数val3被忽略)



//FUTEX_CMP_REQUEUE (since Linux 2.6.7)

//本操作首先检查*uaddr是否等于val3，如果不相等，则返回error EAGAIN。否则的话，
//将唤醒最多val个等待者，如果等待者的数量大于val，则剩下的等待者则从uaddr的等待队列
//中移除，添加到uaddr2的等待队列中。val2参数指定了移动到uaddr2的等待者的最大数量。
//val的典型值为0或1，指定为INT_MAX是没有任何用处的，这将会使得FUTEX_CMP_REQUEUE
//操作和 FUTEX_WAKE 操作相同。同样，val2的值通常为1或INT_MAX，指定为0是没有用的，
//这样将使得 FUTEX_CMP_REQUEUE  操作和  FUTEX_WAIT 操作相同。

//FUTEX_CMP_REQUEUE

//代替 FUTEX_REQUEUE 的，两者的不同点在于是否检查uaddr的值，这个检查操作可用来确保
//移除添加操作只在特定的条件下发生，这样就避免了某些 race condition。

//FUTEX_CMP_REQUEUE 和 FUTEX_REQUEUE

//用来避免 FUTEX_WAKE 可能产生的“线程惊群”现象。
//考虑下面的情况，多个线程在等待B(使用futex实现的等待队列):

//代码

//lock(A)
//while (!check_value(V)) {
//      unlock(A);
//      block_on(B);
//      lock(A);
//};
//unlock(A);

//如果一个线程使用 FUTEX_WAKE 唤醒 B上的所有线程，则它们会全部尝试获取锁 A，
//然而在这种情况下唤醒所有的线程时徒劳的，因为除了一个线程之外其他的线程又全部
//阻塞在 A 上，相比之下，REQUEUE 操作可以唤醒一个等待线程然后将剩下的等待线程
//移动到 A 下，当线程解锁 A时其他的线程可以继续执行。

//--------------------------------------------------------------------*
//#cgo CFLAGS: -I./
//#cgo LDFLAGS: -L./ -lhi

//CFLAGS中的-I（大写的i）参数表示.h头文件所在的路径
//LDFLAGS中的-L(大写) 表示.so文件所在的路径 -l(小写的L) 表示指定该路径下的库名称，
//比如要使用libhi.so，则只需用-lhi
//省略了libhi.so中的lib和.so字符，
//关于这些字符所代表的具体含义请自行google表示
