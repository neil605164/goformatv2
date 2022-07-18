package router

import (
	"goformatv2/app/global/helper"
	"goformatv2/app/middleware"

	"github.com/gin-gonic/gin"
)

// RouteProvider 路由提供者
func RouteProvider(r *gin.Engine) {
	// 組合log基本資訊
	if helper.IsDeveloperEnv() {
		r.Use(middleware.WriteLog)
	}

	// api route
	LoadBackendRouter(r)
}
