# 项目开发记录

## 1.环境的配置

主要是将一个代理配置到当前的环境变量当中

GO111MODULE 	on

GOPROXY				https://goproxy.io,direct

把这两个加在系统变量里，理论上来说这样就行了

如果还不行看配置的代理号
git config --global http.proxy http://127.0.0.1:port

git config --global https.proxy https://127.0.0.1:port

将ip地址和端口改成自己本机上的端口号和IP

**然后配置go.mod，一定要保证**

module helloWorld

go 1.23.2

**这里的module helloWorld指向的是自己项目的地址，保证把当前的项目作为主包**

此后就可以使用go get添加新的包

简单提一下go.mod和go.sum的区别，go.sum可以指向特定包，即某个包的某个特定版本

### go.work

go 1.18 引入了功能泛型 (Generics), 同时还引入的多模块工作区 (Workspaces) 和模糊测试 (Fuzzing)。

Go 多模块工作区能够使开发者能够更容易地同时处理多个模块的工作, 如:

方便进行依赖的代码调试(打断点、修改代码)、排查依赖代码 bug
方便同时进行多个仓库/模块并行开发调试

- `go work init` 初始化工作区文件, 用于生成 `go.work` 工作区文件
- go work use 添加新的模块到工作区

## 2.脚手架

脚手架"（scaffold）通常指的是一种快速生成基本项目结构的工具或框架。脚手架通常用于初始化新项目，提供一些基本的目录结构、文件模板和配置文件，以帮助开发人员快速开始项目开发，而不必从头开始手动创建这些基础结构

脚手架一般和DIL一起用

IDL（Interface Definition Language）是一种用于描述软件组件接口的语言。IDL通常用于分布式系统、远程过程调用（RPC）和接口定义方面。通过使用IDL，开发人员可以定义接口的方法、参数和数据类型，而无需关注底层通信细节。

IDL写好，脚手架快速生成一个框架

脚手架CWGO

安装：go install github.com/cloudwego/cwgo@latest 

DIL：thrift

安装go install github.com/cloudwego/thriftgo@latest

转好之后将提前写好的.thrift文件转化

cwgo server --type RPC --module demo --service demo_thrift --idl ../idl/echo.thrift

另外一种工具是porto

这里详细记录一下怎么写

syntax="proto3";//指定版本

package auth;//指定当前Porto文件的路径

还有一个option，指定生成语言的地址，但是这里用的是cwgo，不单一生成一个go

**message** A{

​	int32  user_id= 0;//字段定义和编号

},A自己取名字

解释一下，每一种类型都应该有一个自己的编号，这些字段编号用于标识消息二进制格式中的字段，且在您的消息类型被使用后不应改变。

即通过编号作为标识，后续会把我的字段变成二进制编码

范围1到 15 中的字段编号使用一个字节进行编码，包括字段编号和字段类型。16至 2047范围内的字段编号占用两个字节。因此，您应为出现频率较高的消息元素预留 1到 15 之间的数字。请务必为将来可能添加的常见元素留出一些空间。

尽量编码1到15间

每一个message都会生成一个结构体，用于存储字段信息

service，要处理接口的逻辑

service AuthService {    

rpc DeliverTokenByRPC(DeliverTokenReq) returns (DeliveryResp) {}   

 rpc VerifyTokenByRPC(VerifyTokenReq) returns (VerifyReq) {}

 }

每一个rpc就是一个服务，这里Rpc所接受返回的参数，就是一个个message

## 3.服务注册与发现

使用的技术栈是consul，Consul是一个开源的、分布式的服务发现和配置管理工具，它能够帮助开发人员构建和管理现代化的分布式系统。

服务注册和发现，就是找到用户需要一个服务，找到服务的地址并管理它，你需要发现服务，管理服务，不使用是注销服务

consul的启动指令主要有三种-dv ,-client ,-server

使用consul之前要先拉取镜像，打开容器，同时注意配置好端口地址

## 4.orm

ORM 是 Object Relational Mapping 的缩写，译为“对象关系映射"，它解决了对象和关系型数据库之间的数据交互问题。
简单说就是 使用一个类表示一张表，类中的属性表示表的字段，类的实例化对象表示一条记录

### 使用对象的方法操作数据库

和自动生成 SQL语句相比，手动编写 SQL语句的缺点是非常明显的，主要体现在以下两个方面:
对象的属性名和数据表的字段名往往不一致，我们在编写 SQL语句时需要非常小心，要逐一核对属性名和字段名，确保它们不会出错，而且彼此之间要一一对应。
此外，当 SQL语句出错时，数据库的提示信息往往也不精准，这给排错带来了不小的困难
不同的数据库，对应的 sql语句也不太一样**
sql注入问题

golang里就是grom

![image-20241117230730855](C:/Users/18291/AppData/Roaming/Typora/typora-user-images/image-20241117230730855.png)



这个报错的主要原因是我初始化了两次mysql导致报错

# 项目开发代码

## 1.前端

### 1.1air

air 热加载工具，实时监听文件的变化重启服务

安装：go install github.com/air-verse/air@latest

这样当你的tmpl文件发生改变时，就不用再次启动服务，而是会自动更改，如下图所示

![image-20241107084744252](https://gitee.com/hu-pingyan/picture_bed/raw/master/img/image-20241107084744252.png)

### 1.2 bootstrap

这是一个免费的开源Web开发框架，由HTML、CSS和基于JS的脚本组成

首先保证你的文件夹下面有js和css文件，前端的三个内容，html是骨架，css是装饰，js是电器，增加互动性

**关于为什么JS文件或者CSS文件写在body之后仍然影响到body的文本？**

当浏览器加载并解析HTML文档时，它会按顺序读取标签并构建DOM树。`<h1>Hello, world!</h1>`会被解析成一个DOM节点，并添加到DOM树中。

加载完之后，才会使用JS文件或者CSS文件进行操作

## 2.后端

关于后端的代码，我一直都理解错了，现在开始整理

![image-20241112204936355](https://gitee.com/hu-pingyan/picture_bed/raw/master/img/image-20241112204936355.png)

比如这个get他看上去是一个请求，实际上是一个针对get请求的响应，他设置了你使用get方法/sign_in这个路由时，我的响应是什么，我的响应就是调用这个请求方法中的html打开一个给html文件，响应这个请求，状态码设为200，那我是怎么到达的，在前端里写有超链接

![image-20241112205520258](https://gitee.com/hu-pingyan/picture_bed/raw/master/img/image-20241112205520258.png)

你按下sign in会在URL上加上新的路径，而get请求实际上是在web网址上按下回车是发出的      2024年11月12日21:06:58

### 2.1会话

HTTP是一种无状态的协议，处理客户端的请求时，服务器不会保留任何关于客户端状态的信息。这个特点是把双刃剑，其优势是服务器无需额外增加内存开销去记住客户状态;其问题就是每次过来的请求，服务器都不知道这个请求源自于谁，无法更好的提供服务。
会话即浏览器和服务器的一次通话(session)，它的出现让服务器可实现客户状态的记录。如果要形象去比喻会话，可以理解为一次电话沟通，电话线一端是服务器，另外一端则是浏览器，只要浏览器把请求成功发到服务器上了，电话就接通会话即刻开始。后面的请求-响应过程都在这个会话里面进行，直到浏览器关闭或者服务器清空会话信息(还比双方任意一方挂断了电话)。

使用session进行会话管理

### 2.2rpc

本小节用于记录如何编写一个完整的rpc通讯

首先明确我们需要什么，依据rpc的基本通讯流程，我首先需要client和service的代码，利用cwgo（kitex）的代码生成能力分别生成双端的代码，需要注意的点有两个，1.服务端和客户端共用一个idl作为依赖2.服务端的代码不再生成客户端代码，将客户端代码统一管理（从代码角度就是要pass -use）

完成这些就有了一个基础的rpc框架

在开发前，有一个很操蛋的问题，在 `go mod` 开发模式下，相对路径的导包方式不支持。会报错误:
`local import "./package1" in non-local package`

解决方案

在 `go.mod` 文件中使用 `replace` 语法，参考见视频用户（中）的最后几分钟

  github.com/hcdog/pycode/gomall/rpc_gen => ../../rpc_gen



#### 2.2.1 定义用户功能的数据模型

数据模型：数据库系统的核心组成部分，它提供了一种抽象的方式来描述和组织数据，以及这些数据之间的关系。数据模型定义了数据的结构、类型、约束条件以及数据之间的操作。它是数据库设计和实现的基础，为数据库系统的构建、使用和维护提供了指导。

建立一个结构体，为后续的代码提供可传递的变量
