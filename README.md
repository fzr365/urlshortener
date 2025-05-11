# 🚀 短链接生成器

![Go](https://img.shields.io/badge/Go-1.23.4-blue?logo=go)
![MySQL](https://img.shields.io/badge/MySQL-8.0-blue?logo=mysql)
![Redis](https://img.shields.io/badge/Redis-6.2-red?logo=redis)
![Docker](https://img.shields.io/badge/Docker-20.10-blue?logo=docker)
![License](https://img.shields.io/badge/License-MIT-green)

## 📖 项目简介
短链接生成器是一个将长链接转换为短链接的服务，支持自定义短链接、过期时间设置等功能。  
它是一个 **读多写少** 的高性能项目，适用于高并发场景。

---

## 🛠 技术架构
- **后端语言**：Go
- **数据库**：MySQL
- **缓存**：Redis
- **容器化**：Docker

---

## 🌟 项目亮点
### ⚡ 高性能
- **缓存加速**：使用 Redis 缓存热点数据，减少数据库压力，提升响应速度。
- **异步重试机制**：短链接生成时支持重试，确保唯一性。

### 🔒 稳定可靠
- **唯一性校验**：短链接生成时校验唯一性，避免冲突。
- **过期管理**：支持短链接过期时间设置，自动清理过期数据。

### 📈 可扩展性
- **模块化设计**：清晰的分层架构，方便扩展和维护。
- **支持自定义短链接**：用户可自定义短链接，提高灵活性。

---

## 📋 功能列表
- **长链接转短链接**  
  - 支持自定义短链接  
  - 支持设置过期时间  
- **短链接重定向**  
  - 快速跳转到原始链接  
  - 支持缓存加速  

---

## 📦 快速开始

### 1️⃣ 环境准备
- 安装 [Go](https://golang.org/)
- 安装 [Docker](https://www.docker.com/)

### 2️⃣ 启动服务
```bash
# 启动 MySQL
make lanch_mysql

# 启动 Redis
make lanch_redis

# 数据库迁移
make migrate_up
```

### 3️⃣ 运行项目
```bash
go run main.go
```

## 📂 项目结构
```
├── config/          # 配置文件
├── database/        # 数据库迁移和查询
├── internal/        # 内部核心逻辑
│   ├── api/         # API 层
│   ├── cache/       # 缓存实现
│   ├── model/       # 数据模型
│   ├── repo/        # 数据库操作
│   └── service/     # 业务逻辑
├── pkg/             # 公共工具包
└── [README.md](http://_vscodecontentref_/1)        # 项目说明文档
```

## 📊 架构图
```
用户请求 --> API 层 --> Service 层 --> Repo 层 --> 数据库
                      ↘ Cache 层 ↙
```

## 🧩 贡献指南

欢迎贡献代码！请提交 Pull Request 或 Issue。

## 📜 许可证

本项目基于 MIT License 开源。