package userb

import (
	"goformatv2/app/global/errorcode"
	"goformatv2/app/global/structer"
	"goformatv2/app/repository/userr"
	"sync"
)

var singleton *business
var once sync.Once

type IUser interface {
	UserList() ([]structer.UserList, errorcode.Error)
	CreateUser(raw *structer.CreateReq) (apiErr errorcode.Error)
}

type business struct {
	user userr.Interface
}

func Instance() IUser {
	once.Do(func() {
		singleton = &business{
			user: userr.Instance(),
		}
	})
	return singleton
}
