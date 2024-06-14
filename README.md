## 简介

> 本项目旨在编辑一个gin框架的通识模版，下面是项目的目录结构

```
/ginTemplate
├── cmd
│   └── api
│       └── main.go
├── pkg
│   ├── api
│   │   ├── controllers
│   │   │   ├── user_controller.go
│   │   │   └── product_controller.go
│   │   ├── middlewares
│   │   │   ├── auth.go
│   │   │   └── logging.go
│   │   ├── models
│   │   │   ├── user.go
│   │   │   └── product.go
│   │   ├── repositories
│   │   │   ├── user_repository.go
│   │   │   └── product_repository.go
│   │   ├── services
│   │   │   ├── user_service.go
│   │   │   └── product_service.go
│   │   └── validators
│   │       ├── user_validator.go
│   │       └── product_validator.go
│   ├── config
│   │   └── config.go
│   ├── docs
│   │   └── swagger.yaml
│   ├── migrations
│   │   ├── 20230601120000_create_users_table.up.sql
│   │   ├── 20230601120000_create_users_table.down.sql
│   │   └── ...
│   ├── utils
│   │   ├── hash.go
│   │   └── jwt.go
│   ├── logger
│   │   ├── logger.go
│   │   └── logger_test.go
│   └── routes
│       └── routes.go
├── test
│   ├── integration
│   │   ├── user_integration_test.go
│   │   └── product_integration_test.go
│   ├── unit
│   │   ├── user_service_test.go
│   │   └── product_service_test.go
│   └── test_helpers.go
├── .env
├── .gitignore
├── go.mod
└── go.sum
```

## 目录结构详解

- cmd/: 包含应用的入口点 
  - api/: API 入口点 
    - main.go: 应用的主文件
- pkg/: 包含所有的业务逻辑和代码
  - api/: 包含应用的核心业务逻辑
    - controllers/: 处理 HTTP 请求，调用服务层，并返回响应 
    - middlewares/: 定义 Gin 中间件，如认证和日志记录 
    - models/: 定义数据模型和数据库表结构 
    - repositories/: 数据访问层，处理数据库的增删改查操作 
    - services/: 服务层，包含业务逻辑 
    - validators/: 定义请求数据的验证器
  - config/: 包含应用的配置文件，如数据库配置、环境变量等 
  - docs/: 包含 API 文档，如 Swagger 文件 
  - migrations/: 包含数据库迁移文件 
  - utils/: 实用工具库 
  - logger/: 自定义日志库 
  - routes/: 定义应用的路由
- test/: 包含测试文件
  - integration/: 集成测试 
  - unit/: 单元测试
  - test_helpers.go: 测试辅助函数
- .env: 环境变量文件 
- .gitignore: Git 忽略文件 
- go.mod 和 go.sum: Go 依赖管理文件