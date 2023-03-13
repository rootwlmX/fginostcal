package model

import "time"

type User struct {
	Id         int64     `json:"id" xorm:"id"`
	CreateTime time.Time `json:"createTime" xorm:"create_time"`
}
