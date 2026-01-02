package precommit_test

import (
	"testing"

	"github.com/lwmacct/260101-go-pkg-ddd/pkg/adapters/http/routes"
	"github.com/stretchr/testify/assert"
)

// TestRoutes_Bindings 检查声明式路由绑定的完整性。
// 规则：operation 中的每个操作都必须有有效的元数据。
//
// 注意：此测试需要 Registry 已被 BuildRegistryFromRoutes() 填充。
// 在 CI 环境或直接运行 go test 时，Registry 为空，测试会被跳过。
// 要运行此测试，需先启动服务器或手动调用 BuildRegistryFromRoutes()。
func TestRoutes_Bindings(t *testing.T) {
	ops := routes.All()
	if len(ops) == 0 {
		t.Skip("Registry is empty - run server or call BuildRegistryFromRoutes() first")
	}

	for _, o := range ops {
		t.Run(o.String(), func(t *testing.T) {
			// 验证每个 operation 都有有效的元数据
			assert.NotEmpty(t, routes.Method(o), "operation %s missing Method", o)
			assert.NotEmpty(t, routes.Path(o), "operation %s missing Path", o)
		})
	}
}

// TestRoutes_PathFormat 检查路径格式的一致性。
// 规则：所有 API 路径必须以 /api/ 开头。
func TestRoutes_PathFormat(t *testing.T) {
	ops := routes.All()
	if len(ops) == 0 {
		t.Skip("Registry is empty - run server or call BuildRegistryFromRoutes() first")
	}

	for _, o := range ops {
		t.Run(o.String(), func(t *testing.T) {
			path := routes.Path(o)
			assert.True(t, len(path) > 0 && path[0] == '/',
				"path should start with /: got %q", path)
			assert.Contains(t, path, "/api/",
				"API path should contain /api/: got %q", path)
		})
	}
}

// TestRoutes_AuditActionsConsistency 检查审计操作的一致性。
// 规则：同一分类的审计操作应该使用一致的命名模式。
func TestRoutes_AuditActionsConsistency(t *testing.T) {
	actions := routes.AllAuditActions()
	if len(actions) == 0 {
		t.Skip("Registry is empty - run server or call BuildRegistryFromRoutes() first")
	}

	for _, a := range actions {
		t.Run(a.Action, func(t *testing.T) {
			assert.NotEmpty(t, a.Category, "audit action %s missing Category", a.Action)
			assert.NotEmpty(t, a.Operation, "audit action %s missing Operation", a.Action)
			assert.NotEmpty(t, a.Label, "audit action %s missing Label", a.Action)
		})
	}
}
