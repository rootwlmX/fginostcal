package util

import (
	"bufio"
	"encoding/json"
	"os"
	"test/model"
)

var _cfg *model.AppInfo = nil

func GetConfig() *model.AppInfo {
	return _cfg
}
func ParseConfig(path string) (*model.AppInfo, error) {

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)

	if err = decoder.Decode(&_cfg); err != nil {
		return nil, err
	}
	return _cfg, nil
}
