package service

import (
	"test/dao"
	"test/engine"
	"test/model"
)

type UserService struct {
}

func (us *UserService) AddUser(cookie string, qq string) (int64, error) {
	userInfo := model.Genshin{
		Qq:     qq,
		Cookie: cookie,
	}

	genshinDao := dao.GenshinDao{DbEngine: engine.GetOrmEngine()}
	return genshinDao.InsertGenshin(userInfo)
}
