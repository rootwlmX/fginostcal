package model

import (
	"time"
)

type User struct {
	Id         int64     `json:"id" xorm:"id"`
	Account    string    `json:"account" xorm:"account"`
	Password   string    `json:"password" xorm:"password"`
	CreateTime time.Time `json:"createTime" xorm:"create_time"`
}

type UserParams struct {
	Id         int64     `json:"id" xorm:"id" binding:"-"`
	Account    string    `json:"account" xorm:"account"`
	Password   string    `json:"password" xorm:"password"`
	CreateTime time.Time `json:"createTime" xorm:"create_time" binding:"-"`
}
