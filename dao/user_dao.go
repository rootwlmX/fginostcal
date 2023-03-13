package dao

import (
	"fginostcal/engine"
	"fginostcal/model"
)

type UserDao struct {
	DbEngine *engine.Orm
}

func (gd *UserDao) Add(user model.User) (int64, error) {
	result, err := gd.DbEngine.InsertOne(user)
	if err != nil {
		return 0, err
	}

	return result, err
}
