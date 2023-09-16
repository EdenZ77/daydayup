参考：https://darjun.github.io/2020/03/13/godailylib/copier/
https://cloud.tencent.com/developer/article/1870553

先安装：go get github.com/jinzhu/copier

test06
浅拷贝是指对地址的拷贝
浅拷贝的是数据地址，只复制指向的对象的指针，此时新对象和老对象指向的内存地址是一样的，新对象值修改时老对象也会变化，释放内存地址时，同时释放内存地址
引用类型的都是浅拷贝：Go 语言官方说的引用类型有三个：slice切片、map映射、chan管道


深拷贝是指将地址指向的值进行拷贝
深拷贝的是数据本身，创造一个一样的新对象，新创建的对象与原对象不共享内存，新创建的对象在内存中开辟一个新的内存地址，新对象值修改时不会影响原对象值。既然内存地址不同，释放内存地址时，可分别释放
值类似的都是深拷贝：int、float、bool、array、struct

