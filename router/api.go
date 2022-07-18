package router

import (
	"goformatv2/app/global/helper"
	"goformatv2/app/handler/pprof"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	// swagger embed files
	swaggerfiles "github.com/swaggo/files"
)

// LoadBackendRouter 路由控制
func LoadBackendRouter(r *gin.Engine) {

	pprof.Register(r, "/internal/debug/pprof/") // 性能

	api := r.Group("/api")

	// K8S Health Check
	api.GET("/healthz", func(c *gin.Context) {
		data := map[string]string{
			"service": os.Getenv("PROJECT_NAME"),
			"time":    time.Now().Format("2006-01-02 15:04:05 -07:00"),
		}

		c.JSON(http.StatusOK, data)
	})

	// 載入測試用API
	if helper.IsDeveloperEnv() {
		// Swagger
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}

}
