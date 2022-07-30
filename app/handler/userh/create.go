package userh

import (
	"goformatv2/app/business/userb"
	"goformatv2/app/global"
	"goformatv2/app/global/errorcode"
	"goformatv2/app/global/helper"
	"goformatv2/app/global/structer"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUser 建立會員
// @Description  建立會員
// @Tags         users
// @Produce  json
// @Param body body structer.CreateReq true "新增用戶"
// @Success      200  {object} structer.UserList
// @Failure      400  {object} structer.APIResult "異常錯誤"
// @Router       /users [post]
func CreateUser(c *gin.Context) {
	// 取參數
	raw := structer.CreateReq{}
	if err := c.ShouldBind(&raw); err != nil {
		apiErr := helper.ErrorHandle(global.WarnLog, errorcode.Code.BindParamError, err, raw)
		c.JSON(http.StatusOK, helper.Fail(apiErr))
		return
	}

	// 驗證參數
	if err := helper.ValidateStruct(raw); err != nil {
		apiErr := helper.ErrorHandle(global.WarnLog, errorcode.Code.ValidateParamError, err.Error(), raw)
		c.JSON(http.StatusOK, helper.Fail(apiErr))
		return
	}

	// 寫入 DB
	bus := userb.Instance()
	if apiErr := bus.CreateUser(&raw); apiErr != nil {
		c.JSON(http.StatusOK, helper.Fail(apiErr))
		return
	}

	// 回傳參數
}
