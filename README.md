# raptor 猛禽运维平台

<div align=center>
<img src="https://img.shields.io/badge/raptor-0.1-blue"/>
<img src="https://img.shields.io/badge/golang-1.16-blue"/>
<img src="https://img.shields.io/badge/gin-1.7.0-lightBlue"/>
<img src="https://img.shields.io/badge/vue-3.2.25-brightgreen"/>
<img src="https://img.shields.io/badge/element--plus-2.0.1-green"/>
<img src="https://img.shields.io/badge/gorm-1.22.5-red"/>
</div>

## 介绍
### 0.1  项目介绍

```
计划从头开始写一个 新的运维平台，尽量保持每月都更新。
项目架构就是vue+go。具体实现的功能会在下面罗列。 

项目模板用的 gin-vue-admin,特别感谢此项目团队！
```

### 0.2 已完成功能

* 钉钉用户系统对接
    *  钉钉登录 (前端需要修改 appid地址和redirect_uri地址,后端需要更新AppKey)
    *  钉钉部门用户信息 定时更新 (task/ding.go   需要更新config ding的AppKey信息)
* 资产管理
    *  云平台密钥管理 (阿里云)
    *  产品线管理 (前端 后端 )
    *  主机管理(根据密钥，定时更新 阿里云主机信息)
* 服务管理
  *  项目管理 (例如: raptor-web)
  *  构建管理 (正在开发中,计划可以在平台打包构建)

### 0.3 demo (不确保是最新版,实现功能根据上面所列)
> http://81.70.255.33:8080      账号admin 密码123456

### 0.4 交流群

> qq: 620176501

### 0.5 快速部署

```shell 
git clone https://github.com/hequan2017/raptor
cd raptor

修改server/config.yaml 里面得配置信息  数据库信息 redis连接信息
docker mysql 是 177.7.0.13 
docker redis 是 177.7.0.14 

docker-compose up -d

连接容器mysql  把raptor.sql 导入。 

docker-compose   restart 

登录平台账号admin，密码123456
```

## 1 开发说明
### 1.1 版本
```
- node版本 > v12.18.3
- golang版本 >= v1.16  < 1.18
- IDE推荐：Goland
-替换掉项目中的七牛云公钥，私钥，仓名和默认url地址，以免发生测试文件数据错乱
```

### 1.2 server项目

使用 `Goland` 等编辑工具，打开server目录

```bash

# 克隆项目
# 进入server文件夹
cd server
# 使用 go mod 并安装go依赖包
go generate
# 编译 
go build -o server main.go (windows编译命令为go build -o server.exe main.go )
# 运行二进制
./server (windows运行命令为 server.exe)
```

### 1.3 web项目

```bash
# 进入web文件夹
cd web

# 安装依赖
cnpm install || npm install

# 启动web项目
npm run serve
```

```
导入数据库  raptor.sql， 库名字叫raptor，数据库连接文件在  server/config.yaml里面 修改连接信息。

测试是用 config.yaml 的 mysql 线上用MysqlProd,根据主机名判断。 有需求可以搜索修改相关设置

if name, _ := os.Hostname(); name == "raptor" {
    fmt.Println("线上环境")
    m = global.GVA_CONFIG.MysqlProd
} else {
    fmt.Println("测试环境")
}

```

---

## 2 swagger自动化API文档

#### 2.1 安装 swagger

##### （1）可以访问外国网站

````
go get -u github.com/swaggo/swag/cmd/swag
````

##### （2）无法访问外国网站

由于国内没法安装 go.org/x 包下面的东西，推荐使用 [goproxy.cn](https://goproxy.cn) 或者 [goproxy.io](https://goproxy.io/zh/)

```bash
# 如果您使用的 Go 版本是 1.13 - 1.15 需要手动设置GO111MODULE=on, 开启方式如下命令, 如果你的 Go 版本 是 1.16 ~ 最新版 可以忽略以下步骤一
# 步骤一、启用 Go Modules 功能
go env -w GO111MODULE=on 
# 步骤二、配置 GOPROXY 环境变量
go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct

# 如果嫌弃麻烦,可以使用go generate 编译前自动执行代码, 不过这个不能使用 `Goland` 或者 `Vscode` 的 命令行终端
cd server
go generate -run "go env -w .*?"

# 使用如下命令下载swag
go get -u github.com/swaggo/swag/cmd/swag
```

#### 2.2 生成API文档

```` shell
cd server
swag init
````

> 执行上面的命令后，server目录下会出现docs文件夹里的 `docs.go`, `swagger.json`, `swagger.yaml` 三个文件更新，启动go服务之后, 在浏览器输入 [http://localhost:8888/swagger/index.html](http://localhost:8888/swagger/index.html) 即可查看swagger文档


## 3 技术选型

- 前端：用基于 [Vue](https://vuejs.org) 的 [Element](https://github.com/ElemeFE/element) 构建基础页面。
- 后端：用 [Gin](https://gin-gonic.com/) 快速搭建基础restful风格API，[Gin](https://gin-gonic.com/) 是一个go语言编写的Web框架。
- 数据库：采用`MySql`(5.6.44)版本，使用 [gorm](http://gorm.cn) 实现对数据库的基本操作。
- 缓存：使用`Redis`实现记录当前活跃用户的`jwt`令牌并实现多点登录限制。
- API文档：使用`Swagger`构建自动化文档。
- 配置文件：使用 [fsnotify](https://github.com/fsnotify/fsnotify) 和 [viper](https://github.com/spf13/viper) 实现`yaml`格式的配置文件。
- 日志：使用 [zap](https://github.com/uber-go/zap) 实现日志记录。

## 4. 项目架构

### 4.1 系统架构图

![系统架构图](http://qmplusimg.henrongyi.top/gva/gin-vue-admin.png)

### 4.2 前端详细设计图 （提供者:<a href="https://github.com/baobeisuper">baobeisuper</a>）

![前端详细设计图](http://qmplusimg.henrongyi.top/naotu.png)

### 4.3 目录结构

```
    ├── server
        ├── api             (api层)
        │   └── v1          (v1版本接口)
        ├── config          (配置包)
        ├── core            (核心文件)
        ├── docs            (swagger文档目录)
        ├── global          (全局对象)                    
        ├── initialize      (初始化)                        
        │   └── internal    (初始化内部函数)                            
        ├── middleware      (中间件层)                        
        ├── model           (模型层)                    
        │   ├── request     (入参结构体)                        
        │   └── response    (出参结构体)                            
        ├── packfile        (静态文件打包)                        
        ├── resource        (静态资源文件夹)                        
        │   ├── excel       (excel导入导出默认路径)                        
        │   ├── page        (表单生成器)                        
        │   └── template    (模板)                            
        ├── router          (路由层)                    
        ├── service         (service层)                    
        ├── source          (source层)                    
        └── utils           (工具包)                    
            ├── timer       (定时器接口封装)                        
            └── upload      (oss接口封装)                        
    
    └─web            （前端文件）
        ├─public        （发布模板）
        └─src           （源码包）
            ├─api       （向后台发送ajax的封装层）
            ├─assets	（静态文件）
            ├─components（组件）
            ├─router	（前端路由）
            ├─store     （vuex 状态管理仓）
            ├─style     （通用样式文件）
            ├─utils     （前端工具库）
            └─view      （前端页面）

```

## 5. 主要功能

- 权限管理：基于`jwt`和`casbin`实现的权限管理。
- 文件上传下载：实现基于`七牛云`, `阿里云`, `腾讯云` 的文件上传操作(请开发自己去各个平台的申请对应 `token` 或者对应`key`)。
- 分页封装：前端使用 `mixins` 封装分页，分页方法调用 `mixins` 即可。
- 用户管理：系统管理员分配用户角色和角色权限。
- 角色管理：创建权限控制的主要对象，可以给角色分配不同api权限和菜单权限。
- 菜单管理：实现用户动态菜单配置，实现不同角色不同菜单。
- api管理：不同用户可调用的api接口的权限不同。
- 配置管理：配置文件可前台修改(在线体验站点不开放此功能)。
- 条件搜索：增加条件搜索示例。
- restful示例：可以参考用户管理模块中的示例API。
	- 前端文件参考: [web/src/view/superAdmin/api/api.vue](https://github.com/flipped-aurora/gin-vue-admin/blob/master/web/src/view/superAdmin/api/api.vue)
    - 后台文件参考: [server/router/sys_api.go](https://github.com/flipped-aurora/gin-vue-admin/blob/master/server/router/sys_api.go)
- 多点登录限制：需要在`config.yaml`中把`system`中的`use-multipoint`修改为true(需要自行配置Redis和Config中的Redis参数，测试阶段，有bug请及时反馈)。
- 分片长传：提供文件分片上传和大文件分片上传功能示例。
- 表单生成器：表单生成器借助 [@form-generator](https://github.com/JakHuang/form-generator) 。
- 代码生成器：后台基础逻辑以及简单curd的代码生成器。


## 作者

> 何全



