# HTTP

### 1.概述

HTTP即超文本传输协议，是应用层协议，当你上网浏览网页的时候，浏览器（客户端）和web服务器（服务端）之间就会通过HTTP在Internet上进行数据的发送和接受。HTTP是一个基于请求/响应模式的、无状态的协议。即我们通常说的Request/Response.

### 2.特点

#####· 支持客户端/服务端模式

#####·简单快速：客户端向服务端请求服务时，只需传送请求方法和路径。由于HTTP 协议简单，使得 HTTP 服务器的程序规模小，因而通信速度很快

##### ·灵活：HTTP 允许传输任意类型的数据对象。正在传输的类型由 Content-Type 加以标记

#####·无连接：无连接的含义是限制每次链接只处理一个请求。服务器处理完客户的请求，并收到客户的应答后，即断开链接，采用这种方式可以节省传输时间

##### ·无状态：HTTP 协议是无状态协议。无状态是指协议对于事物处理没有记忆能力。缺少状态意味着如果后续处理需要前面的信息，则它必须重传，这样可能会导致每次连接传送的数据量增大。另一方面，在服务器不需要先前信息时它的应答就比较快



# URL

### 1.基本组成

通用的格式：schema://host[:port#]/path/…/[?query-string][#anchor]


schema		：  	访问服务器以获取资源时要使用哪种协议，比如，http，https 和 FTP 等


host      		：   	HTTP 服务器的 IP 地址或域名


port#    		：   	HTTP 服务器的默认端口是 80，这种情况下端口号可以省 略，                                                                            如果使用了  		 别的端口，必须指明，例如[www.cnblogs.com：8080](https://link.juejin.im?target=http%3A%2F%2Fwww.cnblogs.com%EF%BC%9A8080)


path    		：    访问资源的路径


query-string 	： 	发给 http 服务器的数据

anchor            ：     锚

举个例子：

www.hapi666.com/cxx/say/goodbey...

其中：

schema是http

host是www.hapi666.com

path是/cxx/say/goodbey

无Query-string

无anthor



话不多说，上图！

![URL 组成](https://user-gold-cdn.xitu.io/2017/11/16/15fc2525666dc96e?imageView2/0/w/1280/h/960/ignore-error/1)

### HTTP之请求篇

HTTP的请求报文分为三个部分：请求行，请求头，请求体

![请求报文](https://user-gold-cdn.xitu.io/2017/11/16/15fc2525665a3211?imageView2/0/w/1280/h/960/ignore-error/1)

#### 1.请求行

请求行分为三个部分：请求方法、请求地址和协议版本

请求方法：

常见的：GET和POST

总体有：GET、POST、PUT、DELETE、OPTIONS、HEAD、TRACE、CONNECT

如果是RESTful接口的话一般会用到PUT、DELETE、GET、POST（分别对应增删查改）

#### 2.请求头

请求头可以用于传递一些附加信息，格式为：键: 值，注意冒号后面有一个空格。

![请求头](https://user-gold-cdn.xitu.io/2017/11/16/15fc25256b98c98b?imageView2/0/w/1280/h/960/ignore-error/1)

请求和响应常见的Header:

Content-Type       请求体/响应体  的类型，如：text/plain、application/json

Accept                   说明接受的类型，可以多个值用,(英文逗号)分隔开

Content-length     请求体/响应体  的长度，单位字节

Accept-Encoding   请求体/响应体  的长度，单位字节

