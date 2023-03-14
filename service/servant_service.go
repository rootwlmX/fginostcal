package service

import (
	"fginostcal/dao"
	"fginostcal/engine"
	"fginostcal/model"
)

type ServantService struct {
}

func (ss *ServantService) List(page int, size int) *[]model.Servant {
	servantDao := dao.ServantDao{DbEngine: engine.GetOrmEngine()}
	return servantDao.List(page, size)
}
