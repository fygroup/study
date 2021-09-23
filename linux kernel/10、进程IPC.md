

### linux IPC API
1、概念
```
https://zhuanlan.zhihu.com/p/37872762

1、管道：是第一个广泛使用的IPC形式，既可以在程序中使用，也可以在shell中使用。管道存在的问题在于他们只能在具有共同祖先（指父子进程之间）的进程间使用，不过该问题已经被有名管道（named pipe）即FIFO消息队列解决了。
2、信号量
3、消息队列：消息队列是在两个不相关进程间传递数据的一种简单、高效方式，她独立于发送进程、接受进程而存在。消息队列是数据结构，存放在内存，访问速度快。但管道是文件，存放在磁盘上，访问速度慢。管道是数据流式存取，消息队列是数据块式存取。（rpc也是信号量的一种）
4、共享内存（同一机器下最快）
```

### 管道
```
(1) pipe(无名管道)
    用于父子进程互相传递信号

    #include <unistd.h>

    int fd[2]; //0:读  1:写
    int ret = pipe(fd);
    if (ret==-1)perror();
    pid_t pt = fork();
    if (pt>0){
        close(fd[0]);       //关闭读
        write(fd[1]...);    
    }else{
        close(fd[1]);       //关闭写
        read(fd[0]...);
    }

(2) FIFO(有名管道)
    不同进程互相传递信号
    #include <sys/stat.h>
    mkfifo("file",0755);
    int fd = open("file",O_RDONLY); //O_WRONLY
    close(fd);

    linux保证了写管道的原子性，但是每次写不能大于pipe_buf

```

### 信号量
```c++
(1) POSIX信号量与System V信号量
    用于共享内存的同步，注意区别于线程的互斥量！！！（下面有详解）
    Posix信号量是基于内存的，即信号量值是放在共享内存中的，与文件系统中的路径名对应的名字来标识的。性能更优越
    System v信号量测试基于内核的，它放在内核里面。

(2)POSIX信号量
    > 一个进程创建POSIX信号量
        #include <semaphore>
        #define FILE_MODE (S_IRUSR|S_IWUSR|S_IRGRP|S_IROTH)

        int main(){
            sem_unlink("file");  //防止所需的信号量已存在
            sem_t* mutex;
            if (mutex = sem_open("file",O_CREAT|O_EXCL,FILE_MODE,1) == SEM_FAILED){
                error("mutex");
                exit(-1);
            }
            sem_close(mutex);   //关闭
        }

    > 另外一个进程运用POSIX信号量
        #include <semaphore.h>

        int main(){
            sem_t* mutex;
            if ((mutex = sem_open("file",0)) == SEM_FAILED){ //打开信号量
                error;
            }
            sem_wait(mutex);        //加锁
            ...
            sem_post(mutex);        //释放锁
        }

(3) System V信号量
    #include <sys/sem.h>
    #include <sys/types.h>
    #include <sys/stat.h>
    #include <fcntl.h>

    int sem_id;
    int set_semvalue(){
        union semun sem_union;
        sem_union.val = 1;
        if (semctl(sem_id,0,SETVAL,sem_union) == -1) return(0);
        return(1);
    }
    void del_semvalue(){
        union semun sem_union;
        if (semctl(sem_id,0,IPC_RMID,sem_union) == -1) perror();
    }
    int semaphore_p(){
        struct sembuf sem_b;
        sem_b.sem_num = 0;
        sem_b.sem_op = -1; //P()
        sem_b.sem_flg = SEM_UNDO;
        if (semop(sem_id,&sem_b,1) == -1)return(0);
        return(1);
    }
    int semaphore_v(){
        struct sembuf sem_b;
        sem_b.sem_num = 0;
        sem_b.sem_op = 1; //V()
        sem_b.sem_flg = SEM_UNDO;
        if (semop(sem_id,&sem_b,1) == -1)return(0);
        return(1);
    }

    int main(){
        key_t key = ftok("file",3);
        sem_id = semget(key,1,0666|IPC_CREAT); //创建信号量
        if (!set_semvalue()) perror();  //初始化信号量

        if (!semaphore_p()) perror;  //进入临界区
        ...
        if (!semaphore_v()) perror; //离开临界区

        del_semvalue();
    }
```

### 共享内存
```
分两种
System V的shmget()得到一个共享内存对象的id，用shmat()映射到进程自己的内存地址
POSIX的shm_open()打开一个文件，用mmap映射到自己的内存地址
注意：以上两种方式要用信号量同步
```
<img src="../picture/7.png" alt="shm_open+mmap" height=300 width=500/>

```c++
(1) System V共享内存
    shmget() shmat() shmdt()
    > 进程一 read
        #include<sys/shm.h>
        #include <sys/types.h>
        #include <sys/ipc.h>
        #define MEM_KEY (1234)

        typedef struct _shared{
            int text[10];
        }shared;

        int main(){

            //proj_id是一个1－255之间的一个整数值，典型的值是一个ASCII值
            key_t key = ftok("file", 3); 

            //创建共享内存,如果存在则报错
            //int shmid = shmget(key,sizeof(shared,IPC_CREAT|0666));
            int shmid = shmget(key, sizeof(shared),0666|IPC_CREAT|IPC_EXCL); 
            
            if (shmid == -1) perror();

            //连接当前进程地址空间
            void* shm = shmat(shmid,0,0); 
            if（shm == (void*)-1）perror();
            shared* my = (shared*)shm;
            printf("%d\n",my->text[1]);
            if (shmdt(shm) == -1) perror();    //把共享内存从当前进程分离
            if (shmctl(shmid,IPC_RMID, 0) == -1) perror //删除共享内存
        }

    > 进程二 write
        #include<sys/shm.h>
        #define MEM_KEY (1234)

        typedef struct _shared{
            int text[10];
        }shared;

        int main(){
            key_t key = ftok("file", 3);

            int shmid = shmget(key, sizeof(shared),0666|IPC_CREAT); //创建共享内存
            if (shmid == -1) perror();
            shm = shmat(shmid, 0, 0);
            if（shm == (void*)-1）perror();
            shared* my = (shared*)shm;
            my->text[1] = 5;
            if (shmdt(shm) == -1) perror();
        }

(2) POSIX共享内存
    shm_open() mmap() shm_unlink()
    > server
        #include <sys/mmap.h>
        #include <sys/shm.h>
        #include <sys/stat.h>
        #include <fcntl.h>
        #define  FILE_MODE (S_IRUSR | S_IWUSR | S_IRGRP | S_IROTH)

        typedef struct __ST{
            char text[5];
        }ST;

        int main(){
            shm_unlink("file");  //防止file已存在
            int fd = shm_open("file",O_RDWR|O_CREAT,FILE_MODE);
            if (fd == -1) perror();
            ftruncate(fd,sizeof(ST));
            ST* ptr;
            ptr = mmap(NULL,sizeof(ST),PROT_READ|PROT_WRITE,MAP_SHARED,fd,0);
            if (ptr == SEM_FAILED) perror();
            ptr->text[1] = 'a';
            close(fd);

        }

    > client
        #include <sys/mman.h>
        #include <sys/stat.h>
        #include <sys/types.h>
        #include <fcntl.h>
        typedef struct _ST{
            char text[5];
        }ST;

        int main(){
            int fd = shm_open("file",O_RDWR,FILE_MODE);
            ptr = mmap(NULL,sizeof(ST),PROT_READ|PROT_WRITE,MAP_SHARED,fd,0);
            printf("%c",ptr->text[1]);
            close(fd);
        }
```

### 消息队列（一种有名管道）
```
(1) 概念
    消息队列提供了一种从一个进程向另一个进程发送一个数据块的方法。 
    消息队列可以认为是一个消息链表，某个进程往一个消息队列中写入消息之前，不需要另外某个进程在该队列上等待消息的达到，这一点与管道和FIFO相反。

(2) 消息队列与rpc
    1) RPC系统结构：
    +----------+     +----------+
    | Consumer | <=> | Provider |
    +----------+     +----------+
    Consumer调用的Provider提供的服务。
    //特点
    同步调用，对于要等待返回结果/处理结果的场景，RPC是可以非常自然直觉的使用方式。# RPC也可以是异步调用。
    由于等待结果，Consumer（Client）会有线程消耗。
    如果以异步RPC的方式使用，Consumer（Client）线程消耗可以去掉。但不能做到像消息一样暂存消息/请求，压力会直接传导到服务Provider。

    2) Message Queue系统结构：
    +--------+     +-------+     +----------+
    | Sender | <=> | Queue | <=> | Receiver |
    +--------+     +-------+     +----------+
    Sender发送消息给Queue；Receiver从Queue拿到消息来处理。
    //特点
    Message Queue把请求的压力保存一下，逐渐释放出来，让处理者按照自己的节奏来处理。
    Message Queue引入一下新的结点，让系统的可靠性会受Message Queue结点的影响。
    Message Queue是异步单向的消息。发送消息设计成是不需要等待消息处理的完成。
    所以对于有同步返回需求，用Message Queue则变得麻烦了。
```

### 进程间通信总结
```
服务器内进程间通信就用socket，低耦合，性能也还行，不用考虑其他乱七八糟的东西。但是如果考虑高性能，可以FIFO或者共享内存
服务期间用rpc、socket、消息队列
```