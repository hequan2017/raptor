[简体中文](README.md) | [English](README.en.md)

# raptor Ops Platform

An ops platform (CMDB) built with Vue 3 + Gin. It integrates with your DingTalk organization, syncs cloud host assets automatically, and manages product lines and services in one place.

<div align=center>
<img src="https://img.shields.io/badge/raptor-0.1-blue"/>
<img src="https://img.shields.io/badge/golang-1.16-blue"/>
<img src="https://img.shields.io/badge/gin-1.7.0-lightBlue"/>
<img src="https://img.shields.io/badge/vue-3.2.25-brightgreen"/>
<img src="https://img.shields.io/badge/element--plus-2.0.1-green"/>
<img src="https://img.shields.io/badge/gorm-1.22.5-red"/>
</div>

> **⚠️ This project is finished and no longer maintained.** The code and documentation are for reference only. Dependencies such as Go 1.17, MySQL 8.0.21 and Vue 3.2 are outdated; if you intend to reuse this project, please assess dependency security risks and upgrade on your own.

## ✨ Features

raptor is built on top of the [gin-vue-admin](https://github.com/flipped-aurora/gin-vue-admin) architecture, adding ops-oriented CMDB capabilities on top of the general admin features.

### DingTalk Integration

- **DingTalk QR-code login**: after QR authorization, the backend fetches user info by DingTalk code, automatically creates/updates the account and issues a JWT (endpoint `/base/dingLogin`)
- **Scheduled department & user sync**: by default at 16:10 every day, departments, members and avatars are fetched from DingTalk and synced into the platform user table (`server/task/ding.go`)
- **DingTalk robot alerts**: a wrapper for pushing text messages via a DingTalk robot, usable for alerting (`server/task/dingding.go`)

### Asset Management (CMDB)

- **Cloud key management**: store cloud platform AccessKeys (type, region, KeyID, Secret) as credentials for asset sync
- **Host management**: periodically syncs Alibaba Cloud ECS instances by key (hostname, SN, intranet/public IP, CPU/memory, OS, status, start time, etc.); manual sync from the web UI is also supported
- **Product line association**: hosts can be linked to product lines in many-to-many relations, with fields such as primary owner and status, plus conditional search
- **Product line management**: CRUD for product lines

### Service Management

- **Project management**: register services with repository URL, branch, start/stop commands, health-check URL, environment variables, build commands, etc.
- **Build management**: build record management (the original plan was to build packages inside the platform, but the project ended at the record-management stage)

### Admin Capabilities (inherited from gin-vue-admin)

- Permission management based on **JWT + Casbin**; user, role, menu and API management with dynamic menus
- Operation records (audit), dictionary management, multi-point login restriction (requires Redis)
- **Code generator**: auto-generation of backend base logic and simple CRUD code
- **Form generator**: powered by [@form-generator](https://github.com/JakHuang/form-generator)
- **File upload**: local / Qiniu / Alibaba Cloud OSS / Tencent Cloud COS / Huawei Cloud OBS / AWS S3
- **Swagger automated API docs**, zap logging, scheduled cleanup of stale table data

## 🛠 Tech Stack

**Backend** (`server/`)

- Go 1.17 (development requires >= 1.16 and < 1.18), Gin 1.7.0 web framework
- GORM 1.22.5 + MySQL 8.0.21 (docker-compose default), with PostgreSQL configuration reserved
- go-redis v8.11.0, golang-jwt/v4, Casbin v2.11.0
- viper + fsnotify (yaml config), zap logging, robfig/cron scheduled tasks
- swaggo/swag 1.8.0 (API docs); SDKs for Alibaba Cloud ECS/OSS, Qiniu, Tencent COS, Huawei OBS, AWS S3

**Frontend** (`web/`)

- Vue 3.2.25 + Element Plus 2.0.1
- Pinia 2.0.9, Vue Router 4, Axios, ECharts 4.9.0
- Built with Vite 2.8

**Deployment**

- docker-compose (web: nginx / server: golang:alpine / mysql:8.0.21 / redis:6.0.6)
- Kubernetes manifests (`deployment/`)

## 🚀 Quick Start

### One-click Deployment with Docker Compose

```shell
git clone https://github.com/hequan2017/raptor
cd raptor

# Edit the configuration in server/config.yaml: database, Redis, DingTalk AppKey, etc.
# The docker mysql address is 177.7.0.13
# The docker redis address is 177.7.0.14

docker-compose up -d

# Connect to the mysql container (host port 13306) and import raptor.sql from the repo root

docker-compose restart
```

Once started, visit `http://localhost:8080` and log in with account `admin` and password `123456`; the backend API listens on port `8888`.

### Local Development

Requirements: Go >= 1.16 and < 1.18, Node > v12.18.3, MySQL (import `raptor.sql` from the repo root into a database named `raptor`), Redis (needed for multi-point login).

```bash
# Backend
cd server
go generate                    # install go dependencies
go build -o server main.go     # Windows: go build -o server.exe main.go
./server                       # Windows: server.exe

# Frontend
cd web
npm install
npm run serve
```

> The test environment reads the `mysql` config from `config.yaml`; when the hostname is `raptor`, it automatically switches to `mysqlProd` (the production database). See `server/initialize/gorm_mysql.go`.

### Key Configuration (server/config.yaml)

| Section | Purpose |
| --- | --- |
| `mysql` / `mysqlProd` | Test/production database connections, switched automatically by hostname |
| `redis` | Redis connection (multi-point login restriction, etc.) |
| `ding.appkey` / `ding.appsecret` | DingTalk app credentials, used for QR login and organization sync |
| `system.use-multipoint` | Whether to enable multi-point login restriction |
| `jwt.signing-key` | JWT signing key |
| `aliyun-oss` / `qiniu` / `tencent-cos` / `hua-wei-obs` / `aws-s3` / `local` | File upload configuration |
| `timer` | Scheduled cleanup of stale table data (operation records, JWT blacklist) |

> DingTalk login setup: change the appid and redirect_uri in `web/src/view/login/index.vue`, and update `ding` AppKey/AppSecret in `config.yaml`.

### Swagger API Docs

```bash
cd server
swag init
```

After generating, start the server and open [http://localhost:8888/swagger/index.html](http://localhost:8888/swagger/index.html) in your browser. Users in mainland China can configure `goproxy.cn` to speed up dependency downloads.

## 📸 Architecture

![System Architecture](http://qmplusimg.henrongyi.top/gva/gin-vue-admin.png)

Frontend detailed design diagram (provided by [baobeisuper](https://github.com/baobeisuper))

![Frontend Detailed Design](http://qmplusimg.henrongyi.top/naotu.png)

## 📁 Directory Structure

```
raptor
├── server                # Backend (Gin)
│   ├── api/v1            # API layer (system modules / autocode business modules)
│   ├── core              # Entry point and scheduled task registration
│   ├── docs              # Swagger docs
│   ├── initialize        # Initialization (router, DB, Redis, logging)
│   ├── middleware        # Middleware (JWT, Casbin, operation records, etc.)
│   ├── model             # Model layer
│   ├── router            # Router layer
│   ├── service           # Service layer
│   ├── task              # Scheduled tasks (DingTalk sync, Alibaba Cloud asset sync)
│   └── utils             # Utilities
├── web                   # Frontend (Vue 3)
│   └── src/view          # Pages (cmdb / serve / superAdmin / systemTools, etc.)
├── deployment            # Kubernetes manifests
├── docker-compose.yaml   # One-click deployment
└── raptor.sql            # Database init script
```

## 📄 License

Licensed under the [Apache License 2.0](LICENSE).

## 👤 Author & Community

> Author: He Quan ([hequan2017](https://github.com/hequan2017))

### Discussion Group

> qq: 620176501
