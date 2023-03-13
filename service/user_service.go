package service

import (
	"fginostcal/dao"
	"fginostcal/engine"
	"fginostcal/model"
	"fginostcal/util"
	"time"
)

type UserService struct {
}

func (us *UserService) AddUser() (int64, error) {

	userInfo := model.User{
		Id:         util.GenID(),
		CreateTime: time.Now(),
	}

	userDao := dao.UserDao{DbEngine: engine.GetOrmEngine()}
	return userDao.Add(userInfo)
}
