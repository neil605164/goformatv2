package router

import (
	"goformatv2/app/handler/pprof"
	"goformatv2/app/handler/userh"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	// gin-swagger middleware
	// swagger embed files
)

// LoadBackendRouter 路由控制
func LoadBackendRouter(r *gin.Engine) {

	pprof.Register(r, "/api/v1/internal/debug/pprof/") // 性能

	api := r.Group("/api/v1")

	// K8S Health Check
	api.GET("/healthz", func(c *gin.Context) {
		data := map[string]string{
			"service": os.Getenv("PROJECT_NAME"),
			"time":    time.Now().Format("2006-01-02 15:04:05 -07:00"),
		}

		c.JSON(http.StatusOK, data)
	})

	users := api.Group("/users")
	{
		users.GET("/", userh.UserList)
		users.POST("/", userh.CreateUser)
	}

}
