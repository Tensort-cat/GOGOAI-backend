# GOGOAI Backend

GOGOAI 是一个基于 Go 实现的 AI 应用服务平台，这里是后端部分。

## 项目描述

- 基于 Go 构建 AI 应用服务平台，使用 Gin 框架提供高性能 Web 服务
- 集成 AI 助手聊天、图像识别等功能，并开发 Vue 前端应用
- 支持用户注册登录、会话管理、多功能 AI 交互
- 聊天模式分为：
    - 普通模式
    - RAG 模式
    - MCP 模式
- 用户可根据需求自由切换聊天模式

前端项目：

- GitHub：[GOGOAI Frontend](https://github.com/Tensort-cat/GOGOAI-frontend)

---

## 技术栈

- Go
- Gin
- GORM
- Eino
- RAG
- MCP
- Redis
- MySQL
- RabbitMQ
- WebSocket
- Vue

---

## 项目启动

### 1. 修改配置文件

部署前请自行修改：

```bash
config/config.toml
```

---

### 2. 导入数据库表

数据库 DDL 文件：

```bash
model/tables.sql
```

---

### 3. 启动后端服务

```bash
go run main.go
```

---

### 4. 启动 MCP 服务

MCP 服务需要单独启动：

```bash
cd common/mcp
go run main.go
```

---

## 项目结构

```text
GOGOAI-/
├── utils/                     # 通用工具包
│   ├── utils.go
│   └── myjwt/                 # JWT工具
│       └── jwt.go
│
├── service/                   # 业务逻辑层
│   ├── user/
│   ├── session/
│   └── image/
│
├── router/                    # 路由定义
│   ├── user.go
│   ├── router.go
│   ├── Image.go
│   └── AI.go
│
├── model/                     # 数据模型
│   ├── user.go
│   ├── session.go
│   └── message.go
│
├── middleware/                # 中间件
│   └── jwt/
│       └── jwt.go
│
├── dao/                       # 数据访问层
│   ├── user/
│   ├── session/
│   └── message/
│
├── controller/                # 控制器层
│   ├── user/
│   ├── session/
│   ├── image/
│   └── common.go
│
├── config/                    # 配置管理
│   ├── config.toml
│   └── config.go
│
├── common/                    # 公共组件
│   ├── redis/
│   ├── rabbitmq/
│   ├── mysql/
│   ├── image/
│   ├── email/
│   ├── code/
│   └── aihelper/
│
├── main.go
├── go.mod
└── go.sum
```

---

## 功能特性

- 用户注册 / 登录
- JWT 身份认证
- AI 多模式聊天
- RAG 检索增强生成
- MCP 工具调用
- WebSocket 实时通信
- 图像识别
- Redis 缓存
- RabbitMQ 异步消息处理

---