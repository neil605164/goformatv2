package userh

import (
	"goformatv2/app/business/userb"
	"goformatv2/app/global/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserList 用戶清單
// @Description  用戶清單
// @Tags         users
// @Produce      json
// @Success      200  {object} structer.UserList
// @Failure      400  {object} structer.APIResult "異常錯誤"
// @Failure      404  {object} structer.APIResult "異常錯誤"
// @Failure      500  {object} structer.APIResult "異常錯誤"
// @Router       /users/list [get]
func UserList(c *gin.Context) {

	bus := userb.Instance()
	resp, apiErr := bus.UserList()
	if apiErr != nil {
		c.JSON(http.StatusOK, helper.Fail(apiErr))
		return
	}

	c.JSON(http.StatusOK, helper.Success(resp))
}
