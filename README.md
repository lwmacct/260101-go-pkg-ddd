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

## 架构概览

```
pkg/                          # 可复用的 DDD 原子能力
├── domain/                   # 领域层（User、Role、Product...）
├── application/              # 应用层（UseCase Handlers）
├── infrastructure/           # 基础设施层（Repository 实现）
└── adapters/                 # 适配器层（HTTP Handler）

internal/                     # 示例项目（如何组装）
├── container/                # ★ Fx 依赖注入组装点
├── domain/order/             # 示例：自定义领域模块
└── manualtest/               # 集成测试

your-project/                 # 你的项目
├── internal/
│   ├── container/            # 复制自本库，按需裁剪
│   ├── domain/invoice/       # 你的业务领域
│   └── application/...
└── cmd/server/main.go
```

**依赖方向**: `Adapters → Application → Domain ← Infrastructure`

## 快速开始

### 运行示例服务器

```bash
# 确保依赖服务运行（PostgreSQL + Redis）
docker run -d -p 5432:5432 -e POSTGRES_PASSWORD=postgres postgres:16
docker run -d -p 6379:6379 redis:alpine

# 初始化数据库
go run cmd/server/main.go db reset --force

# 启动服务
go run cmd/server/main.go
# 或使用热重载
air
```

**预置账号**: `admin / admin123`

### 在你的项目中使用

**步骤 1：复制 Container**

```bash
cp -r 260101-go-pkg-ddd/internal/container your-project/internal/
```

Container 文件结构：

| 文件            | 职责              | 修改方式         |
| --------------- | ----------------- | ---------------- |
| `types.go`      | Model 列表、配置  | 添加你的 Model   |
| `infra.go`      | DB/Redis/EventBus | 通常无需修改     |
| `repository.go` | 仓储注册          | 添加你的仓储     |
| `usecase.go`    | UseCase 聚合      | 添加你的 UseCase |
| `http.go`       | Handler + 路由    | 添加你的路由     |
| `hooks.go`      | 生命周期钩子      | 通常无需修改     |

**步骤 2：添加自定义模块**

以 `Invoice` 模块为例：

```go
// 1. 创建领域层 internal/domain/invoice/
type Invoice struct {
    ID      uint
    OrderID uint
    Amount  float64
    Status  string
}

// 2. 创建基础设施层 internal/infrastructure/persistence/
type InvoiceModel struct {
    ID      uint    `gorm:"primaryKey"`
    OrderID uint    `gorm:"index;not null"`
    Amount  float64 `gorm:"type:decimal(10,2)"`
    Status  string  `gorm:"size:20"`
}

// 3. 创建应用层 internal/application/invoice/
type CreateHandler struct {
    cmdRepo invoice.CommandRepository
}

// 4. 创建适配器层 internal/adapters/http/handler/
type InvoiceHandler struct {
    createHandler *appInvoice.CreateHandler
}
```

**步骤 3：注册到 Container**

```go
// repository.go
var RepositoryModule = fx.Module("repository",
    fx.Provide(
        persistence.NewUserRepositories,      // 上游仓储
        internalPersistence.NewInvoiceRepositories, // 你的仓储
    ),
)

// usecase.go
var UseCaseModule = fx.Module("usecase",
    fx.Provide(
        newAuthUseCases,     // 上游 UseCase
        newInvoiceUseCases,  // 你的 UseCase
    ),
)
```

### 裁剪策略

不需要某些模块时，直接从 Container 中移除：

```go
var RepositoryModule = fx.Module("repository",
    fx.Provide(
        persistence.NewUserRepositories,
        // persistence.NewRoleRepositories,    // 不需要角色管理
        // persistence.NewSettingRepositories, // 不需要系统设置
    ),
)
```

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

# 手动集成测试
MANUAL=1 go test -v ./internal/manualtest/...
```

## API 文档

运行服务后访问 `/swagger/index.html`

## 参考示例

本库的 `internal/` 目录就是一个完整的使用示例：

- `internal/domain/order/` - 自定义订单领域
- `internal/application/order/` - 订单用例
- `internal/infrastructure/persistence/order_*.go` - 订单持久化
- `internal/adapters/http/handler/order.go` - 订单 HTTP Handler
- `internal/container/` - 组装本库 + 自定义模块

## License

MIT
