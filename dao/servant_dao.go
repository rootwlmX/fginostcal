package dao

import (
	"fginostcal/engine"
	"fginostcal/model"
)

type ServantDao struct {
	DbEngine *engine.Orm
}

func (sd *ServantDao) List(page int, size int) *[]model.Servant {
	servants := make([]model.Servant, 0)
	err := sd.DbEngine.Limit(size, (page-1)*size).Find(&servants)
	if err != nil {
		return nil
	}
	return &servants
}
