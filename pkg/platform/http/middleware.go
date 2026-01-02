package http

// MiddlewareType 中间件类型
type MiddlewareType string

// 中间件名称常量
const (
	MiddlewareRequestID   MiddlewareType = "request_id"
	MiddlewareOperationID MiddlewareType = "operation_id"
	MiddlewareAuth        MiddlewareType = "auth"
	MiddlewareOrgContext  MiddlewareType = "org_context"
	MiddlewareTeamContext MiddlewareType = "team_context"
	MiddlewareRBAC        MiddlewareType = "rbac"
	MiddlewareAudit       MiddlewareType = "audit"
	MiddlewareCORS        MiddlewareType = "cors"
	MiddlewareLogger      MiddlewareType = "logger"
)
