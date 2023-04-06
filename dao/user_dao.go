package dao

import (
	"fginostcal/engine"
	"fginostcal/model"
	"fmt"
	"log"
)

type UserDao struct {
	DbEngine *engine.Orm
}

func (ud *UserDao) Add(user model.User) (int64, error) {
	result, err := ud.DbEngine.InsertOne(user)
	if err != nil {
		return 0, err
	}

	return result, err
}

func (ud *UserDao) QueryByUserAccount(account string) (bool, *model.User) {
	user := model.User{Account: account}
	has, err := ud.DbEngine.Get(&user)
	if err != nil {
		log.Panic(err)
		return false, nil
	}
	fmt.Println(has)
	return has, &user
}
