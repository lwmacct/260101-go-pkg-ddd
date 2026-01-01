// Package main 提供示例服务器入口。
//
// 这是一个完整的服务器示例，展示了如何使用 starter 模块快速搭建应用。
// 使用方可以参考此代码来创建自己的服务器实现。
package main

import (
	"github.com/lwmacct/260101-go-pkg-ddd/starter/config"
	starterfx "github.com/lwmacct/260101-go-pkg-ddd/starter/fx"
	"go.uber.org/fx"
)

func main() {
	cfg := config.DefaultConfig()

	// 从环境变量覆盖配置（可选）
	// 例如：os.Getenv("APP_PGSQL_URL")

	app := fx.New(
		// 提供配置
		fx.Supply(cfg),

		// Starter 模块
		fx.StartTimeout(30*1000000000), // 30s
		fx.StopTimeout(10*1000000000),  // 10s

		// 基础设施模块
		starterfx.InfraModule,
		starterfx.CacheModule,
		starterfx.RepositoryModule,
		starterfx.ServiceModule,
		starterfx.UseCaseModule,
		starterfx.HTTPModule,
	)

	// 运行应用
	app.Run()
}
