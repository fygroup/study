### muduo 架构
```rust
EventLoopThread
    EventLoop*
    Thread
    MutexLock
    Condition

EventLoopThreadPool
    EventLoop*
    vector<EventLoopThread*>        threads_

EventLoop
    thread                          threadId_           CurrentThread::tid()
    eventfd                         wakeupFd_           createEventfd()
    Poller*
    TimerQueue*
    Channel*                        wakeupChannel_
    Channel*                        currentActiveChannel_
    vector<Channel*>                activeChannels_
    vector<Functor>                 pendingFunctors_

    wakeupChannel_->setReadCallback(EventLoop::handleRead)
    wakeupChannel_->enableReading()

TimerQueue
    EventLoop*
    timefd                          timerfd_            createTimerfd()
    Channel                         timerfdChannel_     Channel(loop, timefd)
    TimerList                       timers_
    TimerSet                        activeTimers_
    TimerSet                        cancelingTimers_

TcpServer
    EventLoop*                      loop_
    EventLoopThreadPool*            threadPool_
    Acceptor*                       acceptor_
    map<string, TcpConnection*>     connections_
    
    acceptor_.newConnectionCallback_ = TcpServer::newConnection

Acceptor
    EventLoop*
    Socket                          acceptSocket_       ::socket(listenAddr)
    Channel                         acceptChannel_      Channel(loop, acceptSocket_.fd())
    NewConnectionCallback           newConnectionCallback_;
    bool                            listening_;
    int                             idleFd_;

    acceptChannel_.setReadCallback(Acceptor::handleRead)

TcpConnection
    EventLoop*
    Socket*
    Channel*
    InetAddress                     localAddr_
    InetAddress                     peerAddr_
    Buffer                          inputBuffer_
    Buffer                          outputBuffer_

Channel
    EventLoop*
    const int                       fd_
    int                             events_(0)
    int                             revents_(0)
    int                             index_(-1)

TcpClient
    EventLoop*
    Connector*
    TcpConnection*

Connector
    EventLoop*
    Channel*

Poller
    EventLoop*
    int                             epollfd_
    map<int, Channel*>              channels_
    vector<epoll_event>             events_

```

### muduo 函数调用
```c++

// 1
TcpServer::start()
    threadPool_->start()
    loop_->runInLoop(std::bind(&Acceptor::listen, get_pointer(acceptor_)))

// 2
EventLoop::loop()
    activeChannels_ = poll()
    for (curChannel : activeChannels_)
        curChannel.handleEvent()
            curChannel::handleEventWithGuard(receiveTime)
                EventLoop::handleRead(receiveTime) 




```