package bootstrap

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lwmacct/260101-go-pkg-ddd/pkg/config"
)

// Server 管理 HTTP Server 生命周期
type Server struct {
	engine *gin.Engine
	addr   string
}

// NewServer 创建 HTTP Server
func NewServer(engine *gin.Engine, cfg *config.Config) *Server {
	return &Server{
		engine: engine,
		addr:   cfg.Server.Addr,
	}
}

// Start 启动 HTTP Server
func (s *Server) Start() error {
	srv := &http.Server{
		Addr:    s.addr,
		Handler: s.engine,
	}
	return srv.ListenAndServe()
}

// Stop 优雅关闭 HTTP Server
func (s *Server) Stop(ctx context.Context) error {
	// TODO: 实现优雅关闭
	return nil
}
