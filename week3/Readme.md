#说明
基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。



思路：
    启动一个http server。
    用errgroup 起两个Go func 
    1. 监听 退出信号，和ctx.Done通道，向一个chan中通知shutdown
    2. 接受chan 中的shutdown数据，保证server退出。
    
    