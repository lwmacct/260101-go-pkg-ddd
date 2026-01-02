package http

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/lwmacct/260101-go-pkg-ddd/ddd/core/adapters/http/middleware"
	"github.com/lwmacct/260101-go-pkg-ddd/ddd/core/adapters/http/routes"
)

// MiddlewareFactory 无参数中间件工厂函数。
type MiddlewareFactory func(*RouterDependencies) gin.HandlerFunc

// MiddlewareFactoryWithRoute 需要 Route 的中间件工厂函数。
type MiddlewareFactoryWithRoute func(*RouterDependencies, routes.Route) gin.HandlerFunc

// MiddlewareFactoryWithConfig 需要 Config 的中间件工厂函数（用于支持 Options 参数）。
type MiddlewareFactoryWithConfig func(*RouterDependencies, routes.MiddlewareConfig) gin.HandlerFunc

// Factory 中间件工厂定义。
type Factory struct {
	// Factory 工厂函数，为以下类型之一：
	//   - MiddlewareFactory
	//   - MiddlewareFactoryWithRoute
	//   - MiddlewareFactoryWithConfig
	Factory any

	// NeedsRoute 标记工厂函数是否需要访问 Route。
	NeedsRoute bool

	// NeedsDeps 标记工厂函数是否需要 RouterDependencies。
	NeedsDeps bool

	// UsesConfig 标记工厂函数是否需要访问 MiddlewareConfig（用于 Options 参数）。
	UsesConfig bool
}

// MiddlewareRegistry 中间件注册表，将中间件类型映射到工厂函数。
//
// 外部项目可以覆盖此注册表以自定义中间件实现：
//
//	http.MiddlewareRegistry[routes.MiddlewareAuth] = http.Factory{
//	    Factory: func(deps *http.RouterDependencies) gin.HandlerFunc {
//	        return myCustomAuthMiddleware()
//	    },
//	    NeedsDeps: true,
//	}
var MiddlewareRegistry = map[routes.MiddlewareType]Factory{
	// RequestID - 无依赖
	routes.MiddlewareRequestID: {
		Factory: MiddlewareFactory(func(_ *RouterDependencies) gin.HandlerFunc {
			return middleware.RequestID()
		}),
		NeedsRoute: false,
		NeedsDeps:  false,
	},

	// OperationID - 需要 Route.Op
	routes.MiddlewareOperationID: {
		Factory: MiddlewareFactoryWithRoute(func(_ *RouterDependencies, route routes.Route) gin.HandlerFunc {
			return middleware.SetOperationID(route.Op.Identifier())
		}),
		NeedsRoute: true,
		NeedsDeps:  false,
	},

	// Auth - 需要 JWTManager, PATService, PermissionCacheService
	routes.MiddlewareAuth: {
		Factory: MiddlewareFactory(func(deps *RouterDependencies) gin.HandlerFunc {
			return middleware.Auth(
				deps.JWTManager,
				deps.PATService,
				deps.PermissionCacheService,
			)
		}),
		NeedsRoute: false,
		NeedsDeps:  true,
	},

	// OrgContext - 需要 OrgMemberQuery
	routes.MiddlewareOrgContext: {
		Factory: MiddlewareFactory(func(deps *RouterDependencies) gin.HandlerFunc {
			return middleware.OrgContext(deps.OrgMemberQuery)
		}),
		NeedsRoute: false,
		NeedsDeps:  true,
	},

	// TeamContext - 需要 TeamQuery, TeamMemberQuery，支持 optional 选项
	routes.MiddlewareTeamContext: {
		Factory: MiddlewareFactoryWithConfig(func(deps *RouterDependencies, config routes.MiddlewareConfig) gin.HandlerFunc {
			// 提取 optional 选项
			optional := false
			if val, ok := config.Options["optional"].(bool); ok {
				optional = val
			}

			if optional {
				return middleware.TeamContextOptional(deps.TeamQuery, deps.TeamMemberQuery)
			}
			return middleware.TeamContext(deps.TeamQuery, deps.TeamMemberQuery)
		}),
		NeedsRoute: false,
		NeedsDeps:  true,
		UsesConfig: true,
	},

	// RBAC - 需要 Route.Op
	routes.MiddlewareRBAC: {
		Factory: MiddlewareFactoryWithRoute(func(_ *RouterDependencies, route routes.Route) gin.HandlerFunc {
			return middleware.RequireOperation(route.Op)
		}),
		NeedsRoute: true,
		NeedsDeps:  false,
	},

	// Audit - 需要 CreateLogHandler 和 Route.Op
	routes.MiddlewareAudit: {
		Factory: MiddlewareFactoryWithRoute(func(deps *RouterDependencies, route routes.Route) gin.HandlerFunc {
			return middleware.AuditMiddlewareWithOp(deps.CreateLogHandler, route.Op)
		}),
		NeedsRoute: true,
		NeedsDeps:  true,
	},

	// CORS - 无依赖
	routes.MiddlewareCORS: {
		Factory: MiddlewareFactory(func(_ *RouterDependencies) gin.HandlerFunc {
			return middleware.CORS()
		}),
		NeedsRoute: false,
		NeedsDeps:  false,
	},

	// Logger - 无依赖
	routes.MiddlewareLogger: {
		Factory: MiddlewareFactory(func(_ *RouterDependencies) gin.HandlerFunc {
			return middleware.LoggerSkipPaths("/health")
		}),
		NeedsRoute: false,
		NeedsDeps:  false,
	},
}

// buildMiddlewaresFromConfig 从声明式配置构建中间件链。
// 根据 Route.Middlewares 配置，通过中间件注册表动态构建中间件链。
func buildMiddlewaresFromConfig(deps *RouterDependencies, route routes.Route) []gin.HandlerFunc {
	var mws []gin.HandlerFunc

	for _, config := range route.Middlewares {
		factory, ok := MiddlewareRegistry[config.Name]
		if !ok {
			slog.Warn("unknown middleware", "name", config.Name)
			continue
		}

		var handler gin.HandlerFunc
		switch f := factory.Factory.(type) {
		case MiddlewareFactory:
			handler = f(deps)
		case MiddlewareFactoryWithRoute:
			handler = f(deps, route)
		case MiddlewareFactoryWithConfig:
			handler = f(deps, config)
		default:
			slog.Warn("invalid factory signature", "name", config.Name)
			continue
		}

		mws = append(mws, handler)
	}

	return mws
}
