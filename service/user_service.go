package service

import (
	"crypto/md5"
	"errors"
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
	exist, _ := checkAccountExist(params.Account)
	if exist {
		return 0, errors.New("account already exist")
	}
	params.Id = util.GenID()
	params.Password = fmt.Sprintf("%x", md5.Sum([]byte(params.Password)))
	params.CreateTime = time.Now()

	userDao := dao.UserDao{DbEngine: engine.GetOrmEngine()}
	return userDao.Add(model.User(params))
}

func (us *UserService) Login(params model.UserParams) (string, error) {
	userDao := dao.UserDao{DbEngine: engine.GetOrmEngine()}
	status, user := userDao.QueryByUserAccount(params.Account)
	if !status {
		return "", errors.New("account doesn't exist")
	}
	if user.Password != fmt.Sprintf("%x", md5.Sum([]byte(params.Password))) {
		return "", errors.New("password doesn't match")
	} else {
		return "login success", nil
	}
}

func checkAccountExist(account string) (bool, *model.User) {
	userDao := dao.UserDao{DbEngine: engine.GetOrmEngine()}
	return userDao.QueryByUserAccount(account)
}
