package model

import "time"

type Servant struct {
	Id           int       `json:"id" xorm:"id"`
	Name         string    `json:"name" xorm:"name"`
	Star         int       `json:"star" xorm:"star"`
	Ultimate     int       `json:"ultimate" xorm:"ultimate"`
	UltimateType int       `json:"ultimateType" xorm:"ultimate_type"`
	CreateTime   time.Time `json:"createTime" xorm:"create_time"`
}
