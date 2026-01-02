// Package main 提供简单示例服务器入口。
//
// 这是一个最小化的服务器示例，展示如何使用 starter 模块快速搭建应用。
// 相比 cmd/server，此示例不包含 CLI 命令支持，适合学习和快速原型开发。
//
// 完整的服务器实现（含 db migrate/reset/seed 命令）请参考 cmd/server。
package main

import (
	"time"

	"github.com/gin-gonic/gin"

	// Swagger docs - 空白导入触发 docs.go 的 init() 函数
	// 注意：首次运行前需要执行 swag init 生成文档
	// swag init -g examples/simple-server/main.go -o examples/simple-server/docs --parseDependency
	_ "github.com/lwmacct/260101-go-pkg-ddd/examples/simple-server/docs"

	// Swagger UI
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/lwmacct/251207-go-pkg-cfgm/pkg/cfgm"
	"github.com/lwmacct/251219-go-pkg-logm/pkg/logm"
	"github.com/lwmacct/260101-go-pkg-ddd/pkg/config"
	starterfx "github.com/lwmacct/260101-go-pkg-ddd/starter/fx"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

// Swagger 总体配置 - 示例项目可自定义
//
//	@title           Simple Server Example
//	@version         1.0
//	@description     展示如何扩展 go-pkg-ddd 框架的示例服务器
//	@host            localhost:8080
//	@BasePath        /
//
//	@contact.name    Example Support
//	@contact.url     https://github.com/lwmacct/260101-go-pkg-ddd
//
//	@license.name    MIT
//	@license.url     https://opensource.org/licenses/MIT

// nopLogger 空日志记录器，不输出任何 Fx 框架日志。
type nopLogger struct{}

func (nopLogger) LogEvent(fxevent.Event) {}

func main() {
	logm.MustInit(logm.PresetAuto()...)
	cfg := cfgm.MustLoad(
		config.DefaultConfig(),
		cfgm.WithEnvPrefix("APP_"), // 配置环境变量前缀，如 APP_SERVER_FX_LOG_ENABLED
	)

	// 构建 Fx 选项
	fxOptions := []fx.Option{
		// 提供配置
		fx.Supply(cfg),

		// Starter 模块（按依赖顺序）
		fx.StartTimeout(30 * time.Second),
		fx.StopTimeout(10 * time.Second),

		// 基础设施模块
		starterfx.InfraModule,
		starterfx.CacheModule,
		starterfx.RepositoryModule,
		starterfx.ServiceModule,
		starterfx.UseCaseModule,
		starterfx.HTTPModule,
		starterfx.HooksModule,

		// Swagger 端点注册 - 使用者决定是否启用
		fx.Invoke(func(r *gin.Engine) {
			r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}),
	}

	// 根据 config 控制日志：默认禁用 Fx 框架的依赖注入日志
	if !cfg.Server.FxLogEnabled {
		fxOptions = append(fxOptions, fx.WithLogger(func() fxevent.Logger {
			return nopLogger{}
		}))
	}

	app := fx.New(fxOptions...)

	// 运行应用
	app.Run()
}
