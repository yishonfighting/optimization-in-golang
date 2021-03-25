### 处理函数包说明

本包旨在针对一个文本读取的脚本进行pprof的分析，并根据分析结果进行针对性的调优

---
func1 : 原生的处理逻辑,使用简单的时间处理分析
```
time go run main.go moby.txt
```
结果如下：
```
"moby.txt": 428544 words
go run main.go moby.txt  1.02s user 1.17s system 106% cpu 2.060 total
```
针对一个42W数量的文本进行字数的统计，花了1S，直观感觉貌似效率一般般？会是什么原因呢？
针对代码做了pprof 的CPU分析:
![func1](https://github.com/yishonfighting/optimization-in-golang/blob/master/dealFunc/pic/1616637428774.jpg)