package ginutil

import (
	"github.com/gin-gonic/gin"

	authDomain "github.com/lwmacct/260101-go-pkg-ddd/ddd/iam/domain/auth"
	"github.com/lwmacct/260101-go-pkg-gin/pkg/response"
)

// GetUserID 从 Context 获取当前用户 ID。
// 失败时自动返回 401 响应，调用方应检查 ok 后立即 return。
func GetUserID(c *gin.Context) (uint, bool) {
	userID, exists := c.Get("user_id")
	if !exists {
		response.Unauthorized(c, "")
		return 0, false
	}

	id, ok := userID.(uint)
	if !ok {
		response.Unauthorized(c, authDomain.ErrInvalidUserContext.Error())
		return 0, false
	}

	return id, true
}
