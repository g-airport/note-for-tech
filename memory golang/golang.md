package memory

使用 pprof 来检测内存泄漏 \
内存泄漏一般发生在 堆 

1.内存池 \
2.垃圾回收 \
3.大小切分 向上取整 页内存 块内存 避免内存碎片 \
4.多线程加锁 线程通信 （消息队列） 全局性的分配链 
 Golang 内存结构 \
 MHeap MCentral MCache \
 1.MHeap 分配堆 适用场景 申请大块内存 为下层 MCentral MCache 提供内存服务 \
 基本单位 : MSpan(若干连续内存页的数据结构 \
 MSpan 双端链表 \
 type MSpan struct { \
  	MSpan   *next \
  	MSpan   *prev \
  	PageId  start  //起始页号 \
  	unitptr npages //页数 \
 } 
 
2. MCache 运行分配池 每个线程都有自己的局部内存缓存MCache（局部） \
实现goroutine高并发的重要因素 分配小对象可直接从MCache中分配 不用加锁 提升了并发效

3. MCentral:作为MHeap和MCache的承上启下的连接 \
   承上:MHeap申请MSpan \ 
   启下:MSpan划分为各种大小的对象------>MCache使用 \
   type MCentral struct { \
		lock mutex;        //因为会有多个 P 过来竞争 \
		sizeClass int32; \
		noempty mSpanList; //mspan 双向链表 当前 mcentral 中可用的 mSpan list \
		empty mSpanList;   //已经被使用的 一种对所有 mSpan 的 track \
		int32 nfree; \
		…… \
 	} 

	type mSpanList struct { \
		first *mSpan \
		last  *mSpan \
	} 

分配流程 \
1.object size > 32K 则使用 mheap 直接分配 

2.object size < 16 byte 使用 mcache 的小对象分配器 tiny 直接分配 （tiny 是一个指针）

3.object size > 16 byte && size <=32K byte 时 先使用 mcache 中对应的 size class 分配

4.如果 mcache 对应的 size class 的 span 已经没有可用的块 则向 mcentral 请求

5.如果 mcentral 也没有可用的块 则向 mheap 申请 并切分

6.如果 mheap 也没有合适的 span 则想操作系统申请

Summary 

Golang (tcmalloc(thread-caching mallo)) 全局缓存堆 进程的私有缓存 

 MHeap就是全局缓存堆 \
 MCache作为线程私有缓
 
1.MHeap是一个全局变量 负责向系统申请内存 mallocinit()函数进行初始化 \
 如果分配内存对象大于32K直接向MHeap申请
 
2.MCache线程级别管理内存池 \
 关联结构体P \
 主要是负责线程内部内存申请
 
3.MCentral连接MHeap与MCache \
  MCache内存不够则向MCentral申请 \
  MCentral不够时向MHeap申请内存