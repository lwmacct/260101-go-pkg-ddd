// Package permission 定义权限系统的核心抽象。
//
// 本包提供 URN 格式的权限标识符和匹配算法，不包含任何业务特定的常量。
//
// # 核心类型
//
//   - [Operation]: 操作标识符 {scope}:{type}:{action}
//   - [Resource]: 资源标识符 {scope}:{type}:{id}
//   - [Resolver]: 运行时变量解析器（@me, @org, @team）
//
// # 使用方式
//
// 定义业务操作常量（在适配层或使用方）：
//
//	// 在 HTTP 适配层定义
//	const UserCreate permission.Operation = "admin:users:create"
//	const ProfileUpdate permission.Operation = "self:profile:update"
//
// 检查权限匹配：
//
//	if permission.MatchOperation("admin:users:*", string(UserCreate)) {
//	    // 有权限
//	}
//
// 运行时变量替换（用于资源权限）：
//
//	r := permission.NewResolver(map[string]string{"@me": "123", "@org": "acme"})
//	resource := r.ResolveResource("self:user:@me")     // "self:user:123"
//	resource = r.ResolveResource("org.@org:team:*")    // "org.acme:team:*"
//
// # URN 格式
//
// 三段式结构：{scope}:{type}:{identifier}
//
//	Scope 层级：
//	  - public: 公开（无需认证）
//	  - sys:    系统管理
//	  - self:   当前用户
//	  - org.{id}: 组织级（如 org.acme）
//	  - org.{id}.team.{id}: 团队级（如 org.acme.team.dev）
//
// 通配符：
//   - *:*:* 匹配所有
//   - sys:*:* 匹配 sys 域所有操作
//   - sys.*:*:* 匹配 sys 及其子域（如 sys.admin）
//
// # 设计原则
//
// 本包是通用库，可复用于 HTTP、CLI、gRPC 等任何场景。
// 业务特定的操作常量应在各自的适配层定义。
package permission
