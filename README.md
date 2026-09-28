[简体中文](README.md) | [English](README.en.md)

# raptor 猛禽运维平台

基于 Vue 3 + Gin 的运维平台（CMDB），打通钉钉组织架构，自动同步云上主机资产，统一管理产品线与服务。

<div align=center>
<img src="https://img.shields.io/badge/raptor-0.1-blue"/>
<img src="https://img.shields.io/badge/golang-1.16-blue"/>
<img src="https://img.shields.io/badge/gin-1.7.0-lightBlue"/>
<img src="https://img.shields.io/badge/vue-3.2.25-brightgreen"/>
<img src="https://img.shields.io/badge/element--plus-2.0.1-green"/>
<img src="https://img.shields.io/badge/gorm-1.22.5-red"/>
</div>

> **⚠️ 本项目已完结，后续不再更新维护。** 代码与文档仅供参考。项目依赖的 Go 1.17、MySQL 8.0.21、Vue 3.2 等版本均已较旧，如需二次使用请自行评估依赖安全风险并升级。

## ✨ 功能特性

raptor 基于 [gin-vue-admin](https://github.com/flipped-aurora/gin-vue-admin) 架构二次开发，在通用后台管理能力之上，叠加了面向运维场景的 CMDB 能力。

### 钉钉组织对接

- **钉钉扫码登录**：前端扫码授权后，后端根据钉钉 code 拉取用户信息，自动创建/更新账号并签发 JWT（接口 `/base/dingLogin`）
- **部门与用户定时同步**：默认每日 16:10 拉取钉钉部门、成员、头像等信息，同步到平台用户表（`server/task/ding.go`）
- **钉钉机器人报警**：封装钉钉机器人 text 消息推送，可供告警使用（`server/task/dingding.go`）

### 资产管理（CMDB）

- **云平台密钥管理**：维护云平台 AccessKey（类型、区域、KeyID、Secret），作为资产同步凭证
- **主机管理**：根据密钥定时同步阿里云 ECS 实例信息（主机名、SN、内外网 IP、CPU/内存、系统、状态、上线时间等），也支持页面上手动同步
- **产品线关联**：主机资产与产品线多对多挂靠，支持一级负责人、资产状态等字段维护与条件搜索
- **产品线管理**：产品线增删改查

### 服务管理

- **项目管理**：登记服务信息，包括代码库地址、分支、启动/停止命令、探活地址、环境变量、打包命令等
- **构建管理**：构建记录管理（原计划在平台内完成打包构建，开发至记录管理阶段项目即完结）

### 后台管理能力（继承自 gin-vue-admin）

- 基于 **JWT + Casbin** 的权限管理；用户、角色、菜单、API 管理与动态菜单
- 操作记录（审计）、字典管理、多点登录限制（需开启 Redis）
- **代码生成器**：后台基础逻辑与简单 CRUD 代码自动生成
- **表单生成器**：借助 [@form-generator](https://github.com/JakHuang/form-generator)
- **文件上传**：本地 / 七牛云 / 阿里云 OSS / 腾讯云 COS / 华为云 OBS / AWS S3
- **Swagger 自动化 API 文档**、zap 日志、定时清理过期数据表

## 🛠 技术栈

**后端**（`server/`）

- Go 1.17（开发要求 >= 1.16 且 < 1.18），Web 框架 Gin 1.7.0
- GORM 1.22.5 + MySQL 8.0.21（docker-compose 默认），预留 PostgreSQL 配置
- go-redis v8.11.0、golang-jwt/v4、Casbin v2.11.0
- viper + fsnotify（yaml 配置）、zap 日志、robfig/cron 定时任务
- swaggo/swag 1.8.0（API 文档）、阿里云 ECS/OSS、七牛云、腾讯云 COS、华为云 OBS、AWS S3 SDK

**前端**（`web/`）

- Vue 3.2.25 + Element Plus 2.0.1
- Pinia 2.0.9、Vue Router 4、Axios、ECharts 4.9.0
- Vite 2.8 构建

**部署**

- docker-compose（web: nginx / server: golang:alpine / mysql:8.0.21 / redis:6.0.6）
- Kubernetes 编排文件（`deployment/`）

## 🚀 快速开始

### Docker Compose 一键部署

```shell
git clone https://github.com/hequan2017/raptor
cd raptor

# 修改 server/config.yaml 里的配置信息：数据库、Redis、钉钉 AppKey 等
# docker mysql 地址为 177.7.0.13
# docker redis 地址为 177.7.0.14

docker-compose up -d

# 连接容器内 mysql（宿主机端口 13306），把根目录下的 raptor.sql 导入

docker-compose restart
```

启动后访问 `http://localhost:8080`，默认账号 `admin`，密码 `123456`；后端 API 服务监听 `8888` 端口。

### 本地开发

环境要求：Go >= 1.16 且 < 1.18、Node > v12.18.3、MySQL（导入根目录 `raptor.sql`，库名 `raptor`）、Redis（开启多点登录时需要）。

```bash
# 后端
cd server
go generate                    # 安装 go 依赖包
go build -o server main.go     # Windows: go build -o server.exe main.go
./server                       # Windows: server.exe

# 前端
cd web
npm install
npm run serve
```

> 测试环境读取 `config.yaml` 的 `mysql` 配置；当主机名为 `raptor` 时自动切换为 `mysqlProd`（线上库），相关逻辑见 `server/initialize/gorm_mysql.go`。

### 关键配置（server/config.yaml）

| 配置段 | 用途 |
| --- | --- |
| `mysql` / `mysqlProd` | 测试/生产数据库连接，按主机名自动切换 |
| `redis` | Redis 连接信息（多点登录限制等） |
| `ding.appkey` / `ding.appsecret` | 钉钉企业应用密钥，用于扫码登录与组织同步 |
| `system.use-multipoint` | 是否开启多点登录限制 |
| `jwt.signing-key` | JWT 签名密钥 |
| `aliyun-oss` / `qiniu` / `tencent-cos` / `hua-wei-obs` / `aws-s3` / `local` | 文件上传配置 |
| `timer` | 定时清理过期数据表（操作记录、jwt 黑名单） |

> 钉钉登录需要：前端修改 `web/src/view/login/index.vue` 中的 appid 与 redirect_uri，后端更新 `config.yaml` 中 `ding` 的 AppKey/AppSecret。

### Swagger API 文档

```bash
cd server
swag init
```

生成后启动服务，浏览器访问 [http://localhost:8888/swagger/index.html](http://localhost:8888/swagger/index.html) 查看 API 文档。国内网络可配置 `goproxy.cn` 加速依赖下载。

## 📸 系统架构

![系统架构图](http://qmplusimg.henrongyi.top/gva/gin-vue-admin.png)

前端详细设计图（提供者：[baobeisuper](https://github.com/baobeisuper)）

![前端详细设计图](http://qmplusimg.henrongyi.top/naotu.png)

## 📁 目录结构

```
raptor
├── server                # 后端（Gin）
│   ├── api/v1            # 接口层（system 系统模块 / autocode 业务模块）
│   ├── core              # 启动入口与定时任务注册
│   ├── docs              # Swagger 文档
│   ├── initialize        # 初始化（路由、DB、Redis、日志）
│   ├── middleware        # 中间件（JWT、Casbin、操作记录等）
│   ├── model             # 模型层
│   ├── router            # 路由层
│   ├── service           # 服务层
│   ├── task              # 定时任务（钉钉同步、阿里云资产同步）
│   └── utils             # 工具包
├── web                   # 前端（Vue 3）
│   └── src/view          # 页面（cmdb / serve / superAdmin / systemTools 等）
├── deployment            # Kubernetes 编排文件
├── docker-compose.yaml   # 一键部署
└── raptor.sql            # 数据库初始化脚本
```

## 📄 License

本项目基于 [Apache License 2.0](LICENSE) 开源。

## 👤 作者 & 交流

> 作者：何全（[hequan2017](https://github.com/hequan2017)）

### 交流群

> qq: 620176501
