### 锁和并发控制
```
http://shouce.jb51.net/sqlite/14.html
https://huili.github.io/lockandimplement/machining.html

粗粒度，文件级别锁，读写锁

五种方式的文件锁状态
1) UNLOCKED
2) SHARED
3) RESERVED
4) PENDING
5) EXCLUSIVE
```

### lock
```
http://www.chiark.greenend.org.uk/doc/sqlite3-doc/sharedcache.html

事务级别锁
    读事务、写事务
    一个事务默认是读事务，直到发生写操作，这是就变成写事务

表锁
    1) serialize(默认)
        每张表允许一个写锁和多个读锁共存，读数据要读锁，写数据要写锁
    2) read-uncommitted
        PRAGMA read_uncommitted = true;

Shared-Cache Locking (v3.3.0)

file:aux.db?cache=shared

```

### sqlite3_bind_text
```c++
int sqlite3_bind_text(sqlite3_stmt*, int, const char*, int n,void(*)(void*));

// 第二个参数为序号（从1开始）
// 第三个参数为字符串值
// 第四个参数为字符串长度
// 第五个参数为一个函数指针，SQLITE3执行完操作后回调此函数，通常用于释放字符串占用的内存
// 第五个参数还有两个常数
    // SQLITE_STATIC告诉sqlite3_bind_text函数字符串为常量，可以放心使用
    // SQLITE_TRANSIENT会使得sqlite3_bind_text函数对字符串做一份拷贝
```

### sqlite3_step 返回状态
```
#define SQLITE_ROW         100  /* sqlite3_step() has another row ready */
sql执行完成，有结果

#define SQLITE_DONE        101  /* sqlite3_step() has finished executing */
sql执行完成，没有结果


while (sqlite3_step(stmt) == SQLITE_ROW) {
    
}
```
