# 一款针对于IDE的反制蜜罐

> 无意中看到的一位师傅的思路：[利用项目配置文件进行-RCE-IDE-Trust-Project-功能探究](https://rmb122.com/2021/10/02/%E5%88%A9%E7%94%A8%E9%A1%B9%E7%9B%AE%E9%85%8D%E7%BD%AE%E6%96%87%E4%BB%B6%E8%BF%9B%E8%A1%8C-RCE-IDE-Trust-Project-%E5%8A%9F%E8%83%BD%E6%8E%A2%E7%A9%B6/)

> 复现之后觉得很有意思，并且结合自己平时的习惯，拿到泄露的源码也肯定是第一时间用ide审，所以自己改了一些点，写了这个项目。

> 项目使用gin框架开发

### 运行

[![jcygUA.jpg](https://s1.ax1x.com/2022/07/11/jcygUA.jpg)](https://imgtu.com/i/jcygUA)

### 基本配置

本项目无需过多配置，只需要一份前端静态页的源码

· view 目录下放index.html js,css,fonts,img等目录则是放静态资源，favicon.ico也可以直接放在main.go的同级目录

· source目录下则是jb小子要打开的目录，里面的src目录可以放一下没用的源码来增加压缩包的体积

### 反制思路

当jb小子遇到登录框束手无策时都会把思路转向目录扫描，这时如果扫出来一个诱人的目录，如[source.zip] ["子域名".zip] 等，都会打开，然后落入蓝队的陷阱。

※ 需要注意的是 idea 2020.3.3 版本之前都是可以直接rce，这个版本之后加入了安全模式，会在控制台询问

### 使用方法

```
./ide-honeypot -h [address] -p [port] -f [zipfilename] -c [command]
```

※ zipfilename 无需跟后缀，默认是zip包

example：
```
./ide-honeypot -h 0.0.0.0 -p 6789 -f source -c "open /System/Applications/Calculator.app/Contents/MacOS/Calculator"
```

演示视频：

[演示视频](https://www.bilibili.com/video/BV1JY4y1J7NQ/)



