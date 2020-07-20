# [Nginx简介](https://www.cnblogs.com/lililixuefei/p/11823948.html)

### **1. 什么是Nginx**

　　*Nginx* (engine x) 是一个高性能的[HTTP](https://baike.baidu.com/item/HTTP)和[反向代理](https://baike.baidu.com/item/反向代理/7793488)web服务器，Nginx是一款[轻量级](https://baike.baidu.com/item/轻量级/10002835)的[Web](https://baike.baidu.com/item/Web/150564) 服务器/[反向代理](https://baike.baidu.com/item/反向代理/7793488)服 务器及[电子邮件](https://baike.baidu.com/item/电子邮件/111106)（IMAP/POP3）代理服务器，在BSD-like 协议下发行。其特点是占有内存少，[并发](https://baike.baidu.com/item/并发/11024806)能力强，事实上nginx的并发能力在同类型的网页服务器中表现较好，中国大陆使用nginx网站用户有：百度、[京东](https://baike.baidu.com/item/京东/210931)、[新浪](https://baike.baidu.com/item/新浪/125692)、[网易](https://baike.baidu.com/item/网易/185754)、[腾讯](https://baike.baidu.com/item/腾讯/112204)、[淘宝](https://baike.baidu.com/item/淘宝/145661)等。、

Nginx就是反向代理服务器。

首先我们先来看看什么是代理服务器，代理服务器一般是指局域网内部的机器通过代理服务发送请求到互联网上的服务器，代理服务器一般作用于客户端。比如GoAgent，FQ神器。

![img](https://images0.cnblogs.com/i/345474/201408/081121520378740.png)

一个完整的代理请求过程为：客户端首先与代理服务器创建连接，然后根据代理服务器所使用的代理协议，请求对目标服务器创建连接、或则获得目标服务器的指定资源。Web代理服务器是网络的中间实体。代理位于Web客户端和Web服务器之间，扮演“中间人”的角色。 
HTTP的代理服务器既是Web服务器又是Web客户端。

代理服务器是介于客户端和Web服务器之间的另一台服务器，有了它之后，浏览器不是直接到Web服务器去取回网页，而是通过向代理服务器发送请求，信号会先送到代理服务器，由代理服务器来取回浏览器所需要的信息并传送给你的浏览器。

正向代理是一个位于客户端和原始服务器之间的服务器，为了从原始服务器取的内容，**客户端向代理发送一个请求并指定目标（原始服务器）**，然后代理向原始服务器转交请求并将获得的内容返回给客户端，客户端必须要进行一些**特别的设置**才能使用正向代理。

反向代理服务器：在服务器端接收客户端的请求，然后把**请求分发给具体的服务器进行处理**，然后再将服务器的响应结果反馈给客户端。Nginx就是其中的一种反向代理服务器软件。
Nginx：Nginx（“engine x”），Nginx是俄罗斯人Igor Sysoev(塞索耶夫)编写的一款高性能的 HTTP 和反向代理服务器。也是一个IMAP/POP3/SMTP代理服务器，也就是说，Nginx本身就可以托管网站，进行HTTP服务处理，也可以作为反向代理服务器使用。

正向代理客户端必须设置正向代理服务器，当然前提是要知道正向代理服务器的IP地址，还有代理程序的端口。
反向代理正好与正向代理相反，对于客户端而言代理服务器就像是原始服务器，并且客户端不需要进行任何特别的设置。**客户端向反向代理的命名空间中的内容发送普通请求，接着反向代理将判断向哪个原始服务器转交请求，并将获得的内容返回给客户端。**

![img](https://images0.cnblogs.com/i/345474/201408/081123484126526.png)

用户A始终认为它访问的是原始服务器B而不是代理服务器Z，但实际上反向代理服务器接受用户A的应答，
从原始资源服务器B中取得用户A的需求资源，然后发送给用户A。**由于防火墙的作用，只允许代理服务器Z访问原始资源服务器B。尽管在这个虚拟的环境下，防火墙和反向代理的共同作用保护了原始资源服务器B，但用户A并不知情。**

简单的说：
正向代理：客户端知道服务器端，通过代理端连接服务器端。代理端代理的是服务器端。
反向代理：所谓反向，是对正向而言的。服务器端知道客户端，客户端不知道服务器端，通过代理端连接服务器端。代理端代理的是客户端。代理对象刚好相反，所以叫反向代理。

**2.Nginx的应用现状**
Nginx 已经在俄罗斯最大的门户网站── Rambler Media（www.rambler.ru）上运行了3年时间，同时俄罗斯超过20%的虚拟主机平台采用Nginx作为反向代理服务器。
在国内，已经有 淘宝、新浪博客、新浪播客、网易新闻、六间房、56.com、Discuz!、水木社区、豆瓣、YUPOO、海内、迅雷在线 等多家网站使用 Nginx 作为Web服务器或反向代理服务器。

**先介绍一下几个概念**

**3. 反向代理**

　　反向代理服务器位于用户与目标服务器之间，但是对于用户而言，反向代理服务器就相当于目标服务器，即用户直接访问反向代理服务器就可以获得目标服务器的资源。同时，用户不需要知道目标服务器的地址，也无须在用户端作任何设定。反向代理服务器通常可用来作为Web加速，即使用反向代理作为Web服务器的前置机来降低网络和服务器的负载，提高访问效率。

**4. 负载均衡**

　　负载均衡*（Load Balance）*其意思就是分摊到多个操作单元上进行执行，例如Web[服务器](https://baike.baidu.com/item/服务器/100571)、[FTP服务器](https://baike.baidu.com/item/FTP服务器)、[企业](https://baike.baidu.com/item/企业/707680)关键应用服务器和其它关键任务服务器等，从而共同完成工作任务。

单个服务器解决不了，我们增加服务器的数量，然后将请求分发到各个服务器上面，将原先请求到单个服务器上面的情况改为将请求分发到多个服务器上，将负载分发到不同的服务器，这就是所说的负载均衡。

**5. 动静分离**

　　为了加快网站的解析速度，可以把动态页面和静态页面由不同的服务器来解析，加快解析速度，降低单个服务器的压力。

**6. Linux下Nginx的下载安装（请自行百度教程！！！）**

**7. Nginx常用命令**

1. 使用nginx操作命令之前必须要**进入nginx目录**

1. 查看nginx的版本： ./nginx -v

1. 启动nginx： ./nginx

1. 关闭nginx： ./nginx -s stop
2. 重新加载nginx： ./nginx -s reload

**8. nginx的配置文件**

**1. nginx配置文件的位置： /usr/local/nginx/conf/nginx.conf**

**2. nginx配置文件的组成：**

　　（1） nginx配置文件有三部分组成：

　　第一部分：全局块

从配置文件开始到events块之间的内容，主要设置一些影响nginx服务器整体运行的配置指令。例如：worker_processes 1; worker_processes的值越大，可以支持的并发处理量也会越多。

　　第二部分：events块

events块涉及的指令主要影响Nginx服务器与用户的网络连接，例如：worker_connect 1024; 表示最大连接数。

　　第三部分：http块

 Nginx 服务器配置中最频繁的部分，http块中包含了http全局块和server块。

**8.Nginx的特点**
（1）跨平台：Nginx 可以在大多数 Unix like OS编译运行，而且也有Windows的移植版本。
（2）配置异常简单，非常容易上手。配置风格跟程序开发一样，神一般的配置
（3）非阻塞、高并发连接：数据复制时，磁盘I/O的第一阶段是非阻塞的。官方测试能够支撑5万并发连接，在实际生产环境中跑到2～3万并发连接数.(这得益于Nginx使用了最新的epoll模型)
（4）事件驱动：通信机制采用epoll模型，支持更大的并发连接。
**（5）master/worker结构：一个master进程，生成一个或多个worker进程**
（6）内存消耗小：处理大并发的请求内存消耗非常小。在3万并发连接下，开启的10个Nginx 进程才消耗150M内存（15M*10=150M） 
（7）成本低廉：Nginx为开源软件，可以免费使用。而购买F5 BIG-IP、NetScaler等硬件负载均衡交换机则需要十多万至几十万人民币
（8）内置的健康检查功能：如果 Nginx Proxy 后端的某台 Web 服务器宕机了，不会影响前端访问。
（9）节省带宽：支持 GZIP 压缩，可以添加浏览器本地缓存的 Header 头。
（10）稳定性高：用于反向代理，宕机的概率微乎其微



**9.Nginx的不为人知的特点**
（1）nginx代理和后端web服务器间无需长连接；
（2）接收用户请求是异步的，即先将用户请求全部接收下来，再一次性发送到后端web服务器，极大的减轻后端web服务器的压力
（3）发送响应报文时，是边接收来自后端web服务器的数据，边发送给客户端的
（4）网络依赖型低。NGINX对网络的依赖程度非常低，理论上讲，只要能够ping通就可以实施负载均衡，而且可以有效区分内网和外网流量
（5）支持服务器检测。NGINX能够根据应用服务器处理页面返回的状态码、超时信息等检测服务器是否出现故障，并及时返回错误的请求重新提交到其它节点上



## apache简介

### 1.1.1 当前互联网主流web服务说明

**静态服务**

1. apache --->中小型静态web服务的主流，web服务器中的老大哥
2. nginx --->大型新兴网站静态web服务主流，web服务器中的出生牛犊
3. lighttpd --->静态web服务不温不火，逐渐被淘汰的意味，社区不活跃，静态效率很高

**动态服务**

1. IIS --->微软的web服务器(asp,aspx)
2. tomcat --->中小型企业动态web服务主流，互联网java容器主流(jsp,do)
3. resin --->大型动态web服务器主流，互联网java容器主流(jsp,do)
4. php(fcgi) --->大中小网站，php程序的解析容器
   配合apache，php不是守护进程，而是mod_php5.so(module)
   配合nginx，lighttpd，php守护进程模式，FCGI模式

### 1.1.2 apache介绍

- Apache HTTP Server（简称Apache）是Apache软件基金会的一个开放源码的网页服务器，是目前世界上使用最广泛的一种web server，它以跨平台，高效和稳定而闻名，可以运行在几乎所有广泛使用的计算机平台上。Apache的特点是简单、速度快、性能稳定，并可做代理服务器来使用。
- Apache是用C语言开发的基于模块化设计的web应用，总体上看起来代码的可读性高于php代码，它的核心代码并不多，大多数的功能都被分割到各种模块中，各个模块在系统启动时按需载入。
- 支持SSL技术，支持多个虚拟主机。Apache是以进程的Prefork模式（还有基于线程的Worker模式）为基础的结构，进程要比线程消耗更多的系统开支，不太适合于多处理器环境，因此，在一个Apache Web站点扩容时，通常是增加服务器或扩充群集节点而不是增加处理器

### 1.1.3 apahce的特点及应用场合

#### 1.1.3.1 apahce的特点

功能强大，配置简单，速度快，应用广泛，性能稳定可靠，并可做代理服务器或负载均衡来使用

1、几乎可以运行在所有的计算机平台上.
2、支持最新的http/1.1协议
3、简单而且强有力的基于文件的配置(httpd.conf).
4、支持通用网关接口(cgi)
5、支持虚拟主机.
6、支持http认证.
7、集成perl.
8、集成的代理服务器
9、可以通过web浏览器监视服务器的状态, 可以自定义日志.
10、支持服务器端包含命令(ssi).
11、支持安全socket层(ssl).
12、具有用户会话过程的跟踪能力.
13、支持fastcgi

#### 1.1.3.2 apache的应用场合

- 使用apache运行静态html网页，图片(处理静态小文件能力不及nginx)
- 使用apache结合php引擎运行php，perl等程序，LAMP被称为经典组合
- 使用apache结合tomcat/redis运行jsp，java等程序，成为中小企业的首选
- 使用apache做代理，负载均衡，rewrite规则过滤等待

## 1.2 安装apache

### 1.2.1 系统环境

```bash
[root@apache ~]# cat /etc/redhat-release 
CentOS Linux release 7.4.1708 (Core) 
[root@apache ~]# uname -a
Linux apache 3.10.0-693.2.2.el7.x86_64 #1 SMP Tue Sep 12 22:26:13 UTC 2017 x86_64 x86_64 x86_64 GNU/Linux
[root@apache ~]# systemctl status firewalld
● firewalld.service - firewalld - dynamic firewall daemon
   Loaded: loaded (/usr/lib/systemd/system/firewalld.service; disabled; vendor preset: enabled)
   Active: inactive (dead)
     Docs: man:firewalld(1)
[root@apache ~]# getenforce 
Disabled

[root@apache ~]# rpm -qa|grep httpd #检查是否安装apache
```

### 1.2.2 安装apache

#### 1.2.2.1 yum安装

```bash
yum install httpd
systemctl enable httpd
systemctl start httpd
```

#### 1.2.2.2 编译安装

```bash
安装依赖
[root@apache ~]# yum -y install gcc gcc-c++ apr-devel apr-util-devel pcre pcre-devel openssl openssl-devel zlib-devel
[root@apache ~]# mkdir /server/tools -p
[root@apache ~]# cd /server/tools
[root@apache tools]# wget http://archive.apache.org/dist/httpd/httpd-2.4.6.tar.gz
[root@apache tools]# tar xf httpd-2.4.6.tar.gz 
[root@apache tools]# cd httpd-2.4.6/
[root@apache httpd-2.4.6]#
./configure \
--prefix=/application/apache2.4.6 \   #安装目录
--enable-deflate \                    #压缩
--enable-expires \                    #浏览器缓存
--enable-headers \                    #http头部
--enable-modules=most \               #激活大多数模块
--enable-so \        
--with-mpm=worker \                   #进程模式，并发大一点
--enable-rewrite                      #伪静态
[root@apache httpd-2.4.6]# make
[root@apache httpd-2.4.6]# make install
[root@apache httpd-2.4.6]# echo $?
0   返回值为0说明成功
[root@apache httpd-2.4.6]# ln -s /application/apache2.4.6/ /application/apache   #创建软连接
[root@apache httpd-2.4.6]# /application/apache/bin/apachectl start   #启动apache
[root@apache httpd-2.4.6]# netstat -lntup|grep 80  #查看端口是否启动
tcp        0      0 0.0.0.0:80              0.0.0.0:*               LISTEN      8930/httpd 

```



**关闭防火墙,service**

```
`[root@http ~]# service iptables stop`
`iptables: Setting chains to policy ACCEPT: filter          [  OK  ]`
`iptables: Flushing firewall rules:                         [  OK  ]`
`iptables: Unloading modules:                               [  OK  ]`
`[root@http ~]# setenforce 0`
```

浏览器访问ip地址就行：

![img](https://img-blog.csdn.net/20180803163302256?watermark/2/text/aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2xpYW5nX29wZXJhdGlvbnM=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70)能够看到It works!说明安装完成！恭喜你，赶紧使用apache吧

###### 修改首页内容

```
[root@http ~]# cat /application/apache/htdocs/index.html   
<html><body><h1>hello httpd!</h1></body></html>		
```

再进行网页查看

![img](https://img-blog.csdn.net/2018080316332370?watermark/2/text/aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2xpYW5nX29wZXJhdGlvbnM=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70)

## 1.3apache目录结构

| 命令    |                          作用                          |
| ------- | :----------------------------------------------------: |
| apache  |           服务安装成功后，主要的目录结构如下           |
| bin     | 程序命令目录[apache执行文件的目录如apachectl，htpassed |
| build   |                                                        |
| cgi-bin |            预设给一些CGI网页程序存放的目录             |
| conf    |                      配置文件目录                      |
| error   |                    默认错误应答目录                    |
| htdocs  |             编译安装时站点目录，web根目录              |
| icons   |                提供apache预设给予的图标                |
| include |                                                        |
| lib     |                                                        |
| logs    |                      默认日志文件                      |
| man     |                    帮助手册所在目录                    |
| manua   |                                                        |
| modules | 动态加载模块目录。例如phpmemcache编译后的模块在这里面  |

| 左对齐        |      居中       | 右对齐 |
| :------------ | :-------------: | -----: |
| col 3 is      | some wordy text |  $1600 |
| col 2 is      |    centered     |    $12 |
| zebra stripes |    are neat     |     $1 |

