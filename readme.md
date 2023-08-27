# 第六届字节跳动青训营

## 一、介绍

基于RPC框架**Kitex**、HTTP框架**Hertz**、ORM框架**GORM**的极简版抖音服务端项目

使用**ETCD**进行服务注册、服务发现，**Jarger**进行链路追踪

使用**JWT**鉴权，**MD5**密码加密

服务治理：cpu过载保护、重试

| 服务名      | 说明（提供相关api的服务）                                                                                                                                                            | 
|----------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| api      | 处理外部http请求                                                                                                                                                                |
| feed     | /douyin/feed/ - 视频流                                                                                                                                                       |
| user     | /douyin/user/register/ - 用户注册 <br/> /douyin/user/login/ - 用户登录 <br/> /douyin/user/ - 用户信息                                                                                 |
| publish  | /douyin/publish/action/ - 视频投稿 <br/> /douyin/publish/list/ - 发布列表                                                                                                         |
| favorite | /douyin/favorite/action/ - 点赞操作 <br/>   /douyin/favorite/list/ - 喜欢列表                                                                                                     | 
| comment  | /douyin/comment/action/ - 评论操作 <br/> /douyin/comment/list/ - 视频评论列表                                                                                                       |
| relation | /douyin/relation/action/ - 关系操作 <br/> /douyin/relatioin/follow/list/ - 用户关注列表 <br/> /douyin/relation/follower/list/ - 用户粉丝列表 <br/> /douyin/relation/friend/list/ - 用户好友列表 |
| message  | /douyin/message/chat/ - 聊天记录 <br/> /douyin/message/action/ - 消息操作                                                                                                         |

## 二、遇到的问题

proto3 由于字段为默认值（比如0值、空串、false等），导致输出json对应字段被隐藏

request参数中有token，总感觉哪里怪怪的

使用Find查询时，查询不到数据不会返回错误

int64(int(claims[constants.IdentityKey].(float64)))  这种写法我蚌埠住了

你肯定关注了你的偶像，但是你的偶像不一定关注你呀

使用空结构体配合map，构成set数据结构：map[int64]struct{}

好友是消息接受方，所以我是消息发送方。这个参数放在userlist中，艹，好混乱啊

请求如果是json数据的话，好像只能使用GetRawData

查询某一个用户的粉丝，这个粉丝的是否被当前登录用户关注，逻辑有点乱呀

默认赋值
int 0
float 0.000000
string ""
指针 nil
数组 []

结构体内部默认赋值就是基础类型赋值

## 三、代码运行

### 1. 更改配置

更改 constants/constant.go 中的地址配置

### 2. 建立基础依赖

```shell
docker-compose up
```

### 3. 运行feed微服务

```shell
cd cmd/feed
sh build.sh
sh output/bootstrap.sh
```

### 4. 运行publish微服务

```shell
cd cmd/publish
sh build.sh
sh output/bootstrap.sh
```

### 5. 运行user微服务

```shell
cd cmd/user
sh build.sh
sh output/bootstrap.sh
```

### 6. 运行favorite微服务

```shell
cd cmd/favorite
sh build.sh
sh output/bootstrap.sh
```

### 7. 运行comment微服务

```shell
cd cmd/comment
sh build.sh
sh output/bootstrap.sh
```

### 8. 运行relation微服务

```shell
cd cmd/relation
sh build.sh
sh output/bootstrap.sh
```

### 9. 运行message微服务

```shell
cd cmd/message
sh build.sh
sh output/bootstrap.sh
```

### 10. 运行api微服务

```shell
cd cmd/api
chmod +x run.sh
./run.sh
```

### 11. Jaeger链路追踪

在浏览器上查看`http://127.0.0.1:16686/`

