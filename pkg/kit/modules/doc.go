// Package modules 提供 DDD 模块的稳定对外 API 入口（Facade 模式）。
//
// 本包封装了 IAM、App、CRM 三个业务模块的 UseCase Handler，
// 为外部项目提供长期稳定的访问接口，隐藏内部依赖注入细节。
//
// # 设计原则
//
//   - 稳定性优先：API 保持向后兼容，内部实现可自由重构
//   - 最小暴露：只暴露 UseCase Handler，不暴露仓储、缓存等底层实现
//   - 类型安全：使用明确的结构体类型，避免 interface{} 或 map[string]any
//   - 与 Fx 解耦：不依赖 Fx 容器，由 pkg/kit/app 负责组装
//
// # 使用示例
//
//	import "github.com/lwmacct/260101-go-pkg-ddd/pkg/kit/modules"
//
//	func main() {
//	    // 通过 pkg/kit/app 构建器获取 Facade（后续实现）
//	    facade := app.MustBuild(cfg)
//
//	    // 使用 IAM 模块
//	    ctx := context.Background()
//	    user, err := facade.IAM.User.Get.Handle(ctx, user.GetByIDQuery{ID: 1})
//
//	    // 使用 App 模块
//	    tasks, err := facade.App.Task.List.Handle(ctx, task.ListQuery{})
//
//	    // 使用 CRM 模块
//	    companies, err := facade.CRM.Company.List.Handle(ctx, company.ListQuery{})
//	}
//
// # 模块结构
//
//   - [Facade]: 总入口，聚合 IAM、App、CRM 三个模块
//   - [IAMFacade]: 身份认证与授权模块（Auth、User、Role、Pat、TwoFA）
//   - [AppFacade]: 核心治理模块（Organization、Setting、Task、Audit 等）
//   - [CRMFacade]: 客户关系管理模块（Company、Contact、Lead、Opportunity）
//
// # 架构说明
//
// 本包位于依赖注入架构的最外层，与 internal/container 的关系：
//
//	internal/container (Fx 容器) → pkg/kit/modules (稳定 Facade)
//	         ↓                              ↓
//	   运行时组装                      编译时类型检查
//
// external/container 负责运行时依赖注入和生命周期管理，
// pkg/kit/modules 提供编译时类型安全和稳定的公共 API。
package modules
