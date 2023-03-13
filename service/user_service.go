package service

import (
	"crypto/md5"
	"fginostcal/dao"
	"fginostcal/engine"
	"fginostcal/model"
	"fginostcal/util"
	"fmt"
	"time"
)

type UserService struct {
}

func (us *UserService) Register(params model.UserParams) (int64, error) {
	params.Id = util.GenID()
	params.Password = fmt.Sprintf("%x", md5.Sum([]byte(params.Password)))
	params.CreateTime = time.Now()

	userDao := dao.UserDao{DbEngine: engine.GetOrmEngine()}
	return userDao.Add(model.User(params))
}
