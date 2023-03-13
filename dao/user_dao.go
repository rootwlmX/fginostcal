package dao

import (
	"fmt"
	"test/engine"
	"test/model"
)

type UserDao struct {
	DbEngine *engine.Orm
}

func (gd *UserDao) Add(genshin model.User) (int64, error) {
	result, err := gd.DbEngine.InsertOne(genshin)
	fmt.Println(result, err)
	if err != nil {
		return 0, err
	}

	return result, err
}
