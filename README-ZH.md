# meteor
An implementation of distributed ID

在雪花算法的设计基础上，解决了时间回拨问题，同时性能更高。

其和[雪花算法](https://github.com/TheFutureIsOurs/learncode/blob/master/snow/snow.go)的benchmark对比。

运行机器：联想小新pro13 Ryzen 5 3550H

go版本：1.15

	goos: windows
	goarch: amd64
	BenchmarkSnowflake-8
	4917643	       244 ns/op	       0 B/op	       0 allocs/op
	BenchmarkMeteor-8
	52173231	   22.6 ns/op	       0 B/op	       0 allocs/op

可以看出流星算法比雪花算法快10倍

### 如何使用

>go get -u github.com/TheFutureIsOurs/meteor

```go

import "github.com/TheFutureIsOurs/meteor"

node, _ := meteor.NewNode(0)
id, _ := node.Generate()

```

强烈建议在线上使用时，保证NodeID的唯一性，包括新增机器及当前机器重启时。你可以借助mysql的自增id来确保NodeID唯一.


### 流星算法


64位，

算法组成：

![](http://www.imflybird.cn/static/img/2020/meteor.png)


[设计初衷见博文](https://www.imflybird.cn/2020/12/09/%E4%BB%8EMongoID%E8%AE%A8%E8%AE%BA%E5%88%86%E5%B8%83%E5%BC%8F%E5%94%AF%E4%B8%80ID%E7%94%9F%E6%88%90%E6%96%B9%E6%A1%88/)

解释：

第一位同样保留（正数）

Data位为当前NodeID被创建时的秒级差，这样只在NodeID创建时需要依赖系统时间，后续生成ID时就无需系统时间，就可以防止时间回拨。

NodeID位为节点ID，为了确保生成ID唯一，如果发生了新增机器或服务重启，则NodeID需要每次增加。这样即使发生了时间回拨，由于NodeID唯一，则可以保证最终生成ID唯一性。

自增序列号，11位，最大2048。每次生成，自增序列号+1，当加满后，Data位+1。

随机数位，3位。为什么不把自增序列号和随机数位合成为自增序列号位呢？主要是为了特性4：非连续性。

雪花算法在生成时依赖了毫秒，时间位很细，只有都在这一毫秒内连续生成的ID才会连续，这种条件非常苛刻（qps达到4096000生成才会连续）。

所以加了这个随机数位来保证生成的ID非连续。那随机数如何生成呢？大多数随机数以系统时间作为种子，但是这样就达不到去系统时间的高性能了，我希望一种不依赖系统时间的高效随机数生成算法。最终选用了[Xorshift算法](https://en.wikipedia.org/wiki/Xorshift)。

生成10个ID如下：

	5016762319896585
	5016762319896596
	5016762319896600
	5016762319896614
	5016762319896616
	5016762319896626
	5016762319896635
	5016762319896640
	5016762319896651
	5016762319896656

总结下来就是，
>1.NodeID创建时依赖当前系统时间，但是生成时不需系统时间来达到去系统时间化，这样就解决了时间回拨。

>2.NodeID每次需要跟以前不重复。这样就保证了全局唯一。

>3.随机数位保证了生成ID非连续。

在使用时建议使用Mysql自增ID来注册NodeID。

