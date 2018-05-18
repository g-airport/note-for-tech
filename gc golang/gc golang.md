#Golang 垃圾回收

##策略

1.引用计数(reference counting) \
&emsp;&emsp;每个对象维护一个引用计数器，
当引用该对象的对象被销毁或者更新的时候，
被引用对象的引用计数器自动减1，
当被应用的对象被创建，或者赋值给其他对象时，引用+1，引用为0的时候回收 \
&emsp;&emsp;缺点:频繁更新引用计数器降低性能

2.标记-清除(mark and sweep) \ (根查找-->标记-->清除)

&emsp;&emsp;从根变量来时遍历所有被引用对象，标记之后进行清除操作，对未标记对象进行回收 

&emsp;&emsp;缺点:每次垃圾回收的时候都会暂停所有的正常运行的代码，系统的响应能力会大大降低

######三色标记法:GC中用三种颜色标记不同的对象

(1)黑色:本身强引用,并已处理对象中的子引用

(2)灰色:本身强引用,还没处理对象中的子引用

(3)白色:不可达对象

Mark扫描时根据状态进行标记(缓解性能问题)

3.分代收集(generation)

&emsp;&emsp;将堆划分为两个或多个称为代（generation）的空间 \
&emsp;&emsp;新创建的对象存放在称为新生代（young generation）中 （一般来说，新生代的大小会比 老年代小很多） 
随着垃圾回收的重复执行，生命周期较长的对象会被提升（promotion）到老年代中（分类思路）

###Tips
1.`减少对象分配`

1.>`func(r *Reader) Read() ([]byte, error)` \
2.>`func(r *Reader) Read(buf []byte) (int, error)`

第一个函数没有形参，每次调用的时候返回一个[]byte \
第二个函数在每次调用的时候，形参是一个buf []byte 类型的对象，之后返回读入的byte的数目

(第一个函数在每次调用的时候都会分配一段空间 这就给gc造成额外压力 
 第二个函数在每次迪调用的时候 会重用形参声明)
 
2.`string 与 []byte 转换`

1.> `type = struct []uint8 {    uint8 *array;    int len;    int cap;}`

2.> `type = struct string {    uint8 *str;    int len;}`

解决策略: \
&emsp;&emsp;一种方式是一直使用[]byte 特别是在数据传输方面 \
[]byte中也包含着许多string会常用到的有效的操作 \
另一种是使用更为底层的操作直接进行转化

优化策略:主要是使用unsafe.Pointer直接进行转化

3. `+ string 操作`
 
&emsp;&emsp;由于采用+来进行string的连接会生成新的对象，降低gc的效率，最好采用方式是通过append函数来进行

弊端：`b := make([]int, 1024)   b = append(b, 99)`  
      `fmt.Println("len:", len(b), "cap:", cap(b))` \
&emsp;&emsp;append操作之后，数组的空间由1024增长到了1312 
######避免这种情况发生: 最初分配空间的时候完成好空间规划操作


