package userr

import (
	"goformatv2/app/global"
	"goformatv2/app/global/errorcode"
	"goformatv2/app/global/helper"
	"goformatv2/app/models"
)

// UserList 會員清單
func (r *repo) UserList() (data []models.User, apiErr errorcode.Error) {

	db, apiErr := r.DB.DBConn()
	if apiErr != nil {
		return
	}

	if err := db.Preload("Review").Find(&data).Error; err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, errorcode.Code.GetUserListError, err)
		return
	}
	return
}
