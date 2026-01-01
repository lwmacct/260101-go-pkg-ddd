// Package main 提供示例服务器入口。
//
// 这是一个完整的服务器示例，展示了如何使用 starter 模块快速搭建应用。
// 使用方可以参考此代码来创建自己的服务器实现。
package main

import (
	"time"

	"github.com/lwmacct/251207-go-pkg-cfgm/pkg/cfgm"
	"github.com/lwmacct/251219-go-pkg-logm/pkg/logm"
	"github.com/lwmacct/260101-go-pkg-ddd/starter/config"
	starterfx "github.com/lwmacct/260101-go-pkg-ddd/starter/fx"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

// nopLogger 空日志记录器，不输出任何 Fx 框架日志。
type nopLogger struct{}

func (nopLogger) LogEvent(fxevent.Event) {}

func main() {
	cfg := cfgm.MustLoad(
		config.DefaultConfig(),
		cfgm.WithEnvPrefix("APP_"), // 配置环境变量前缀，如 APP_SERVER_FX_LOG_ENABLED
	)

	logm.MustInit(logm.PresetAuto()...)

	// 构建 Fx 选项
	fxOptions := []fx.Option{
		// 提供配置
		fx.Supply(cfg),

		// Starter 模块
		fx.StartTimeout(30 * time.Second),
		fx.StopTimeout(10 * time.Second),

		// 基础设施模块
		starterfx.InfraModule,
		starterfx.CacheModule,
		starterfx.RepositoryModule,
		starterfx.ServiceModule,
		starterfx.UseCaseModule,
		starterfx.HTTPModule,
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
