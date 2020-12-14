# meteor
An implementation of distributed ID.

[中文](https://github.com/TheFutureIsOurs/meteor/blob/master/README-ZH.md)

Based on Snowflake,but resolve time back, and have higher performance(10 times faster), I call it Meteor.

Here is the benchmark Meteor compare to [Snowflake](https://github.com/TheFutureIsOurs/learncode/blob/master/snow/snow.go)

Machine：Lenovo xiaoxin pro13 Ryzen 5 3550H

go version：1.15

	goos: windows
	goarch: amd64
	BenchmarkSnowflake-8
	4917643	       244 ns/op	       0 B/op	       0 allocs/op
	BenchmarkMeteor-8
	52173231	   22.6 ns/op	       0 B/op	       0 allocs/op



### Meteor


int64



![](http://www.imflybird.cn/static/img/2020/meteor-en.png)

Explain：

0:the first bit do nothing(positive number)

Data section is the time difference when the current NodeID is created. So Meteor just rely on the system time when one NodeID is created. This can resolve time back.

NodeID section，for the result ID unique, we should keep the NodeID unique. For example, if we add new machine or one machine is reload, we should increase the NodeID.

Serial section is 11 bits, increment when the generator is called. the Data section should +1 when the serial reach 2048,
then the serial began from zero.

Rand section is 3 bits. Why we need this? Just because we want discrete IDs. For example, if the result id is 1,2,3,4 etc,which is easy to get next. Why Snowflake do not have rand num? Because Snowflake is based on millisecond, the qps should reach 4096000, the result ID can continuous.

Because we need rand section,what random algorithm we should choose? As we know, most random algorithm is based on the system time as seed. If we use this, we can't reach our goal:high performace than Snowflake and can't remove system time when generator. We need high performance random algorithm. I choose the [Xorshift算法](https://en.wikipedia.org/wiki/Xorshift)。

Let's generator 10 IDs:

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

We got it.

