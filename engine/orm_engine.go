package engine

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"test/model"
)

var _dbEngine *Orm

type Orm struct {
	*xorm.Engine
}

func GetOrmEngine() *Orm {
	return _dbEngine
}

func NewOrmEngine(appInfo *model.AppInfo) (*xorm.Engine, error) {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", appInfo.UserName, appInfo.Password, appInfo.Host, appInfo.Port, appInfo.DataBase)
	engine, err := xorm.NewEngine(appInfo.DriverName, url)

	fmt.Println(appInfo)

	if err != nil {
		return nil, err
	}

	// 创建表
	// Sync2 synchronize structs to database tables
	err = engine.Sync2(new(model.Genshin))
	if err != nil {
		return nil, err
	}

	orm := new(Orm)
	orm.Engine = engine
	_dbEngine = orm

	return engine, nil
}
