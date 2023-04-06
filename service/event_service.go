package service

import (
	"fginostcal/dao"
	"fginostcal/engine"
)

type EventService struct {
}

func (es *EventService) List(page int, size int) interface{} {
	eventDao := dao.EventDao{DbEngine: engine.GetOrmEngine()}
	return eventDao.List(page, size)
}
