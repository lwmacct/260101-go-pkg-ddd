// Package ginutil 提供 Gin Context 相关的辅助函数。
//
// 本包为 HTTP Handler 提供统一的 Context 操作，包括：
//   - [GetUserID]: 从 Context 获取当前用户 ID
//   - 未来可扩展: GetOrgID, GetTeamID 等
//
// 使用示例：
//
//	userID, ok := ginutil.GetUserID(c)
//	if !ok {
//	    return // 已自动返回 401 响应
//	}
package ginutil
