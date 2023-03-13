package model

type AppInfo struct {
	AppName    string `json:"app_name"`
	AppMode    string `json:"app_mode"`
	DriverName string `json:"driver_name"`
	UserName   string `json:"username"`
	Password   string `json:"password"`
	Host       string `json:"host"`
	Port       string `json:"port"`
	DataBase   string `json:"database"`
}

type Genshin struct {
	Id        int64  `json:"id" xorm:"id"`
	Qq        string `json:"qq" xorm:"qq"`
	Uid       int64  `json:"uid" xorm:"uid"`
	Cookie    string `json:"cookie" xorm:"cookie"`
	Automatic int64  `json:"automatic" xorm:"automatic"`
	Time      int64  `json:"time" xorm:"time"`
}
