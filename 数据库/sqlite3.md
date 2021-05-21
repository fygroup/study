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


```
