package dao

import (
	"fginostcal/engine"
	"fginostcal/model"
	"time"
)

type EventDao struct {
	DbEngine *engine.Orm
}

func (ed *EventDao) List(page int, size int) *[]model.Event {
	events := make([]model.Event, 0)
	err := ed.DbEngine.Limit(size, (page-1)*size).Find(&events)
	if err != nil {
		return nil
	}
	return &events
}

func (ed *EventDao) GetEventsInDayRange(startTime, endTime time.Time) *[]model.Event {
	events := make([]model.Event, 0)
	//err := ed.DbEngine.Where("startTime >= ? and endTime <= ? ", startTime, endTime).Find(&events)
	err := ed.DbEngine.Find(&events)
	if err != nil {
		return nil
	}
	return &events
}
