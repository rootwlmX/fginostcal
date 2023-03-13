package dao

import (
	"fmt"
	"test/engine"
	"test/model"
)

type GenshinDao struct {
	DbEngine *engine.Orm
}

func (gd *GenshinDao) InsertGenshin(genshin model.Genshin) (int64, error) {
	result, err := gd.DbEngine.InsertOne(genshin)
	fmt.Println(result, err)
	if err != nil {
		return 0, err
	}

	return result, err
}
