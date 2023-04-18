package main

import (
	"fginostcal/router"
	"fginostcal/util"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	/*	config, err := util.ParseConfig("./config/test.json")
		if err != nil {
			log.Fatal(err.Error())
			return
		}*/

	/*	_, err = engine.NewOrmEngine(config)
		if err != nil {
			log.Fatal(err.Error())
			return
		}*/

	if err := util.Init("2021-12-03", 1); err != nil {
		fmt.Println("Init() failed, err = ", err)
		return
	}

	r := gin.Default()
	router.InitRouter(r)

	err := r.Run()
	if err != nil {
		return
	}
}
