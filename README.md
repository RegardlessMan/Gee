# Gee
## 1、Gee-Cache

一个简易的分布式缓存实现。[源作者以及实现见](https://github.com/geektutu/7days-golang/blob/master/gee-cache/day2-single-node/geecache/byteview.go)，本仓库为个人学习记录

[groupcache](https://github.com/golang/groupcache) 是 Go 语言版的 `memcached`，目的是在某些特定场合替代 `memcached`。`groupcache `的作者也是 `memcached`的作者。无论是了解单机缓存还是分布式缓存，深入学习这个库的实现都是非常有意义的。

`GeeCache` 基本上模仿了 [groupcache](https://github.com/golang/groupcache) 的实现，为了将代码量限制在 500 行左右（`groupcache `约 3000 行），裁剪了部分功能。但总体实现上，还是与 `groupcache `非常接近的。支持特性有：

- 单机缓存和基于 HTTP 的分布式缓存
- 最近最少访问(Least Recently Used, LRU) 缓存策略
- 使用 Go 锁机制防止缓存击穿
- 使用一致性哈希选择节点，实现负载均衡
- 使用 protobuf 优化节点间二进制通信
- …

**代码核心结构**

```go
gee-cache
    │  byteview.go  //缓存值的抽象和封装 
    │  cache.go   	//并发控制
    │  geecache.go	// 负责与外部的实际交互，控制缓存存储和获取的流程
    └─lru
            lru.go   //lru缓存淘汰策略
            lru_test.go 
```

**核心流程**

```
                           是
接收 key --> 检查是否被缓存 -----> 返回缓存值 ⑴
                |  否                         是
                |-----> 是否应当从远程节点获取 -----> 与远程节点交互 --> 返回缓存值 ⑵
                            |  否
                            |-----> 调用`回调函数`，获取值并添加到缓存 --> 返回缓存值 ⑶
```

