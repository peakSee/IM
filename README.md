# IM(即时通讯)

基于Golang、Websocket、MongoDB 实现即时通讯

# 技术栈

语言: Golang 数据库: MongoDB 框架: GIN 协议: Websocket

# 核心包

https://github.com/gorilla/websocket

# 扩展安装

```bash
go get -u github.com/gin-gonic/gin
go get github.com/gorilla/websocket
go get go.mongodb.org/mongo-driver/mongo
go get -u github.com/golang-jwt/jwt/v4
go get github.com/satori/go.uuid
```

# Docker 安装 mongoDB

```bash
docker run -d --network some-network --name some-mongo \
-e MONGO_INITDB_ROOT_USERNAME=admin \
-e MONGO_INITDB_ROOT_PASSWORD=admin \
-p 27017:27017 \
mongo
```

# 数据字典

## 用户集合

```json
{
  "account": "账号",
  "password": "密码",
  "nickname": "昵称",
  "sex": 1, // 0-未知 1-男 2-女
  "email": "邮箱",
  "avatar": "头像",
  "created_at": 1, //创建时间
  "updated_at": 1, //更新时间
}
```

## 消息集合

```json
{
  "user_identity": "用户唯一标识",
  "room_identity": "房间唯一标识",
  "data": "发送数据",
  "created_at": 1, //创建时间
  "updated_at": 1, //更新时间
}
```

## 房间集合

```json
{
    "number":"房间号",
		"name":"房间名称",
		"info":"房间简介",
		"user_identity":"房间创建者唯一标识",
		"created_at":1, //房间创建时间
		"updated_at":1, //房间更新时间
}
```

## 用户和房间集合

```json
{
    "user_identity":"用户的唯一标识",
		"room_identity":"房间的唯一标识",
		"message_identity":"消息的唯一标识",
		"created_at":1, //创建时间
		"updated_at":1, //更新时间
}
```

