package model

import "time"

type Event struct {
	Id        int       `json:"id" xorm:"id"`
	Name      string    `json:"name" xorm:"name"`
	Spar      int       `json:"spar" xorm:"spar"`
	Coupon    int       `json:"coupon" xorm:"coupon"`
	StartTime time.Time `json:"startTime" xorm:"start_time"`
	EndTime   time.Time `json:"endTime" xorm:"end_time"`
}
