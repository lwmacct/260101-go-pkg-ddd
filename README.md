# Go DDD Package Library

基于领域驱动设计（DDD）和 CQRS 模式的可复用 Go 模块库。

## 特性

- **四层架构**：Domain → Application → Infrastructure → Adapter
- **CQRS 分离**：Command/Query Repository 独立
- **依赖注入**：基于 Uber Fx
- **认证授权**：JWT + PAT 双重认证，URN 风格 RBAC
- **审计日志**：完整操作追踪
- **2FA 支持**：TOTP 双因素认证

## 技术栈

| 组件     | 技术           |
| -------- | -------------- |
| Web 框架 | Gin            |
| ORM      | GORM           |
| 数据库   | PostgreSQL     |
| 缓存     | Redis          |
| 依赖注入 | Uber Fx        |
| 配置管理 | Koanf          |
| API 文档 | Swagger (swag) |

## 快速开始

### 作为库使用

```go
import (
    "github.com/lwmacct/260101-go-pkg-ddd/starter/fx"
    "github.com/lwmacct/260101-go-pkg-ddd/starter/gin"
)

func main() {
    fx.New(
        fx.Supply(cfg),
        starterfx.InfraModule,
        starterfx.RepositoryModule,
        starterfx.UseCaseModule,
        starterfx.HTTPModule,
    ).Run()
}
```

### 运行示例服务器

```bash
# 确保依赖服务运行（PostgreSQL + Redis）
# 方式 1: 使用 Docker
docker run -d -p 5432:5432 -e POSTGRES_PASSWORD=postgres postgres:16
docker run -d -p 6379:6379 redis:alpine

# 方式 2: 使用本地服务

# 初始化数据库
go run cmd/server/main.go db reset --force

# 启动服务
go run cmd/server/main.go
# 或使用热重载
air
```

**预置账号**: `admin / admin123`

## 项目结构

```
pkg/
├── domain/          # 领域层 - 实体、Repository 接口、领域错误
├── application/     # 应用层 - Use Cases (Command/Query Handler)、DTO
├── infrastructure/  # 基础设施层 - Repository 实现、数据库、缓存、事件
└── adapter/         # 适配器层 - HTTP Handler、路由、中间件

starter/
├── fx/              # Fx 依赖注入容器模块
├── gin/             # Gin HTTP 服务器、路由、Handler 基类
└── config/          # 配置管理

cmd/server/          # 示例服务器入口
```

**依赖方向**: `Adapters → Application → Domain ← Infrastructure`

## 开发命令

```bash
# 单元测试
go test ./...

# 编译检查
go build -o /dev/null ./...

# Lint 检查
golangci-lint run --new

# 数据库迁移
go run cmd/server/main.go db migrate

# 重置数据库
go run cmd/server/main.go db reset --force
```

## API 文档

运行服务后访问 `/swagger/index.html`

## License

MIT
