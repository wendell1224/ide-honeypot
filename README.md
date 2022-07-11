# 一款针对于IDE的反制蜜罐

> 无意中看到的一位师傅的思路：[利用项目配置文件进行-RCE-IDE-Trust-Project-功能探究](https://rmb122.com/2021/10/02/%E5%88%A9%E7%94%A8%E9%A1%B9%E7%9B%AE%E9%85%8D%E7%BD%AE%E6%96%87%E4%BB%B6%E8%BF%9B%E8%A1%8C-RCE-IDE-Trust-Project-%E5%8A%9F%E8%83%BD%E6%8E%A2%E7%A9%B6/)

> 复现之后觉得很有意思，并且结合自己平时的习惯，拿到泄露的源码也肯定是第一时间用ide审，所以自己改了一些点，写了这个项目。

> 项目使用gin框架开发

### 使用方法

```
./ide-honeypot -h [address] -p [port] -f [zipfilename] -c [command]
```

※ zipfilename 无需跟后缀，默认是zip包

example：
```
./ide-honeypot -h 0.0.0.0 -p 6789 -f source -c "command"
```

演示视频：





