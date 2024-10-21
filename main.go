package main

import (
	"HW1-OnlineCoursePlatform/cmd/web"
	"HW1-OnlineCoursePlatform/internal/config"
	"HW1-OnlineCoursePlatform/internal/handlers"
)

func main() {
	var conf config.Conf
	conf.GetConf()
	if conf.DB.Active {
		db := conf.InitDatabase()
		handlers.InitPostgresHandler(db)
	} else {
		handlers.InitNeo4jHandler(&conf)
	}
	web.Routes(&conf)
}
