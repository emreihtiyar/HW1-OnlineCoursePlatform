package main

import (
	"HW1-OnlineCoursePlatform/cmd/web"
	"HW1-OnlineCoursePlatform/internal/config"
	"HW1-OnlineCoursePlatform/internal/handlers"
)

func main() {
	var conf config.Conf
	conf.GetConf()
	db := conf.InitDatabase()
	handlers.InitPostgresHandler(db)
	web.Routes(&conf)
}
