package userh

import (
	"goformatv2/app/business/userb"
	"goformatv2/app/global/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserList 用戶清單
func UserList(c *gin.Context) {

	bus := userb.Instance()
	resp, apiErr := bus.UserList()
	if apiErr != nil {
		c.JSON(http.StatusOK, helper.Fail(apiErr))
		return
	}

	c.JSON(http.StatusOK, helper.Success(resp))
}
