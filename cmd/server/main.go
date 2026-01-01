// Package main 提供完整的服务器入口，支持 CLI 命令。
//
// 命令：
//   - server      启动 HTTP 服务器（默认）
//   - db migrate  执行数据库迁移
//   - db reset    重置数据库（删表+重建+种子数据）
//   - db seed     执行种子数据
package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/lwmacct/251207-go-pkg-cfgm/pkg/cfgm"
	"github.com/lwmacct/251219-go-pkg-logm/pkg/logm"
	"github.com/lwmacct/260101-go-pkg-ddd/pkg/config"
	starterfx "github.com/lwmacct/260101-go-pkg-ddd/starter/fx"
	"github.com/urfave/cli/v3"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

var (
	// 全局 flags（持久化到子命令）
	configFile  string
	envPrefix   string
	fxLogEnable bool
)

func main() {
	// 初始化日志
	logm.MustInit(logm.PresetAuto()...)

	app := &cli.Command{
		Name:    "server",
		Usage:   "Go DDD Package Library Server - 完整的生产级 HTTP 服务器",
		Version: "1.0.0",
		Description: `基于 DDD + CQRS 架构的可复用模块库。

		示例:
		  server                  启动 HTTP 服务器
		  server db migrate       执行数据库迁移
		  server db reset         重置数据库
		  server -c config.yaml   指定配置文件
		  server --fx-log        启用 Fx 日志`,
		Commands: []*cli.Command{
			dbCommand(),
		},
		// Persistent flags 会自动传递给所有子命令
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				Usage:       "配置文件路径",
				Destination: &configFile,
			},
			&cli.StringFlag{
				Name:        "env-prefix",
				Usage:       "环境变量前缀",
				Value:       "APP_",
				Destination: &envPrefix,
			},
			&cli.BoolFlag{
				Name:        "fx-log",
				Usage:       "启用 Fx 依赖注入日志",
				Destination: &fxLogEnable,
			},
		},
		EnableShellCompletion: true,
		Action:                startServer,
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		os.Exit(1)
	}
}

// startServer 启动 HTTP 服务器。
func startServer(ctx context.Context, cmd *cli.Command) error {
	cfg := loadConfig(cmd)

	fxOptions := buildFxOptions(cfg)

	fxApp := fx.New(fxOptions...)
	if err := fxApp.Err(); err != nil {
		return fmt.Errorf("create fx app: %w", err)
	}

	fxApp.Run()
	return nil
}

// dbCommand 数据库操作命令组。
// 使用 Persistent flags 使子命令继承父命令的配置选项。
func dbCommand() *cli.Command {
	cmd := &cli.Command{
		Name:  "db",
		Usage: "数据库操作",
		Description: `数据库迁移和种子数据管理。

		示例:
		  server db migrate       执行迁移
		  server db migrate -c dev.yaml  使用指定配置
		  server db reset         重置数据库
		  server db seed          执行种子数据`,
	}

	// 配置 flags，使用变量捕获以便子命令访问
	cmd.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "config",
			Aliases:     []string{"c"},
			Usage:       "配置文件路径",
			Destination: &configFile,
		},
		&cli.StringFlag{
			Name:        "env-prefix",
			Usage:       "环境变量前缀",
			Value:       "APP_",
			Destination: &envPrefix,
		},
	}

	// 添加子命令
	cmd.Commands = []*cli.Command{
		{
			Name:    "migrate",
			Aliases: []string{"m"},
			Usage:   "执行数据库迁移",
			Action:  migrateDatabase,
		},
		{
			Name:    "reset",
			Aliases: []string{"r"},
			Usage:   "重置数据库（删表+重建+种子数据）",
			Action:  resetDatabase,
		},
		{
			Name:    "seed",
			Aliases: []string{"s"},
			Usage:   "执行种子数据",
			Action:  seedDatabase,
		},
	}

	return cmd
}

// loadConfig 加载配置。
// 使用 cfgm.MustLoadCmd 从 CLI 命令和配置文件加载配置。
func loadConfig(cmd *cli.Command) *config.Config {
	// 从 CLI 获取 env-prefix 值
	prefix := envPrefix
	if prefix == "" {
		prefix = cmd.String("env-prefix")
		if prefix == "" {
			prefix = "APP_"
		}
	}

	// 构建 cfgm 选项
	opts := []cfgm.Option{
		cfgm.WithEnvPrefix(prefix),
	}

	// 如果指定了配置文件，添加该路径
	path := configFile
	if path == "" {
		path = cmd.String("config")
	}
	if path != "" {
		opts = append(opts, cfgm.WithConfigPaths(path))
	}

	// 使用 cfgm.MustLoadCmd 加载配置
	return cfgm.MustLoadCmd(cmd, config.DefaultConfig(), "", opts...)
}

// buildFxOptions 构建 Fx 选项。
func buildFxOptions(cfg *config.Config) []fx.Option {
	fxOptions := []fx.Option{
		fx.Supply(cfg),
		fx.StartTimeout(30 * time.Second),
		fx.StopTimeout(10 * time.Second),
		starterfx.InfraModule,
		starterfx.CacheModule,
		starterfx.RepositoryModule,
		starterfx.ServiceModule,
		starterfx.UseCaseModule,
		starterfx.HTTPModule,
		starterfx.HooksModule,
	}

	// CLI --fx-log 优先级高于配置文件
	if !cfg.Server.FxLogEnabled && !fxLogEnable {
		fxOptions = append(fxOptions, fx.WithLogger(func() fxevent.Logger {
			return nopLogger{}
		}))
	}

	return fxOptions
}

// migrateDatabase 执行数据库迁移。
func migrateDatabase(ctx context.Context, cmd *cli.Command) error {
	cfg := loadConfig(cmd)

	appCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	fxApp := fx.New(
		fx.Supply(cfg),
		fx.StartTimeout(5*time.Minute),
		fx.StopTimeout(10*time.Second),
		starterfx.InfraModule,
		fx.Invoke(starterfx.RunMigration),
		fx.WithLogger(func() fxevent.Logger { return nopLogger{} }),
	)

	if err := fxApp.Err(); err != nil {
		return fmt.Errorf("create fx app: %w", err)
	}

	// CLI 命令执行完成后立即退出，不等待信号
	if err := fxApp.Start(appCtx); err != nil {
		return err
	}
	return fxApp.Stop(appCtx)
}

// resetDatabase 重置数据库。
func resetDatabase(ctx context.Context, cmd *cli.Command) error {
	cfg := loadConfig(cmd)

	appCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	fxApp := fx.New(
		fx.Supply(cfg),
		fx.StartTimeout(5*time.Minute),
		fx.StopTimeout(10*time.Second),
		starterfx.InfraModule,
		fx.Invoke(starterfx.RunReset),
		fx.WithLogger(func() fxevent.Logger { return nopLogger{} }),
	)

	if err := fxApp.Err(); err != nil {
		return fmt.Errorf("create fx app: %w", err)
	}

	// CLI 命令执行完成后立即退出，不等待信号
	if err := fxApp.Start(appCtx); err != nil {
		return err
	}
	return fxApp.Stop(appCtx)
}

// seedDatabase 执行种子数据。
func seedDatabase(ctx context.Context, cmd *cli.Command) error {
	cfg := loadConfig(cmd)

	appCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	fxApp := fx.New(
		fx.Supply(cfg),
		fx.StartTimeout(5*time.Minute),
		fx.StopTimeout(10*time.Second),
		starterfx.InfraModule,
		fx.Invoke(starterfx.RunSeed),
		fx.WithLogger(func() fxevent.Logger { return nopLogger{} }),
	)

	if err := fxApp.Err(); err != nil {
		return fmt.Errorf("create fx app: %w", err)
	}

	// CLI 命令执行完成后立即退出，不等待信号
	if err := fxApp.Start(appCtx); err != nil {
		return err
	}
	return fxApp.Stop(appCtx)
}

// nopLogger 空日志记录器，不输出任何 Fx 框架日志。
type nopLogger struct{}

func (nopLogger) LogEvent(fxevent.Event) {}
