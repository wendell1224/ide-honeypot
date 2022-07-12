# 一款针对于IDE的反制蜜罐 IDE-honeypot

> 无意中看到了这位师傅的 [利用项目配置文件进行-RCE-IDE-Trust-Project-功能探究](https://rmb122.com/2021/10/02/%E5%88%A9%E7%94%A8%E9%A1%B9%E7%9B%AE%E9%85%8D%E7%BD%AE%E6%96%87%E4%BB%B6%E8%BF%9B%E8%A1%8C-RCE-IDE-Trust-Project-%E5%8A%9F%E8%83%BD%E6%8E%A2%E7%A9%B6/)

> 复现之后觉得很有意思，结合自身，在网站下到源码也是第一时间开ide审计，所以稍加修改，把这个项目实战化

> 本项目使用gin框架开发

## 运行

[![jcygUA.jpg](https://s1.ax1x.com/2022/07/11/jcygUA.jpg)](https://imgtu.com/i/jcygUA)

## 基本配置

· view 目录下放index.html模板文件

· js/css/fonts/img 等目录则是放静态资源

· favicon.ico放在与main.go同级目录（也可以不需要）

· source目录则是jb小子要打开的目录，src下可以放一下没用的源码增加zip包体积诱惑攻击队

## 反制思路

当jb小子对网站进行目录扫描时，发现一个源码包，恰好又会点代审，便会落入蓝队的陷阱

前端的页面推荐使用登录框，使得jb小子束手无策时把思路转向目录扫描，zip包名称推荐 source.zip [子域名].zip 等

※ 需要注意的是，idea 2020.3.3 之后的版本加入了安全模式，会在控制询问是否执行

## 使用方法

```
./ide-honeypot -h [address] -p [port] -f [zipfilename] -c [command]
```

zipfilename不用加后缀，默认是zip包
port端口参数可以不需要，默认8080

example:

```
./ide-honeypot -h 0.0.0.0 -p 6789 -f source -c "open /System/Applications/Calculator.app/Contents/MacOS/Calculator"
```
## 支持平台

· ✅Linux

· ✅Windows

· ✅MacOS

只需要在静态文件放进对应的文件夹之后运行：

```
go build main.go
```

生成可执行文件即可

## 演示视频：

[演示视频](https://www.bilibili.com/video/BV1JY4y1J7NQ/)

## 项目后续计划

□ 加入vscode反制
