package main

import (
	"HW1-OnlineCoursePlatform/cmd/web"
	"HW1-OnlineCoursePlatform/internal/config"
	"HW1-OnlineCoursePlatform/internal/handlers"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"gorm.io/gorm"
)

func initializeDatabase(conf *config.Conf) {
	var db *gorm.DB = nil
	var neo4jDriver neo4j.DriverWithContext = nil
	if conf.DB.Active {
		db = conf.InitDatabase()
	} else {
		neo4jDriver = conf.InitNeo4jDatabase()
	}
	handlers.InitDBHandler(db, neo4jDriver)
}

func main() {
	var conf config.Conf
	conf.GetConf()
	initializeDatabase(&conf)
	web.Routes(&conf)
	defer handlers.CloseDriver()
}
