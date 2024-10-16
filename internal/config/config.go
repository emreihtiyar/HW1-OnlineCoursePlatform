package config

import (
	"HW1-OnlineCoursePlatform/internal/models"
	"fmt"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

const CONFIG_PATH = "internal/config/config.yaml"

var db *gorm.DB

type DbConnector struct {
	User     string
	Password string
	Ip       string
	Port     int
	DbName   string
}

func (dbConnector DbConnector) Connector() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		dbConnector.Ip, dbConnector.User, dbConnector.Password, dbConnector.DbName, dbConnector.Port)
}

type Conf struct {
	App struct {
		Port string `yaml:"port"`
	} `yaml:"app"`
	DB struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Ip       string `yaml:"ip"`
		Port     int    `yaml:"port"`
		DbName   string `yaml:"name"`
		Active   bool   `yaml:"active"`
	} `yaml:"db"`
}

// GetDbConnector function to extract DbConnector from Conf
func GetDbConnector(conf Conf) DbConnector {
	return DbConnector{
		User:     conf.DB.User,
		Password: conf.DB.Password,
		Ip:       conf.DB.Ip,
		Port:     conf.DB.Port,
		DbName:   conf.DB.DbName,
	}
}

func (c *Conf) GetConf() *Conf {
	yamlFile, err := os.ReadFile(CONFIG_PATH)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func (c *Conf) InitDatabase() *gorm.DB {
	var err error
	if c.DB.Active {
		dbConnector := GetDbConnector(*c)

		db, err = gorm.Open(postgres.Open(dbConnector.Connector()), &gorm.Config{})
		if err != nil {
			log.Fatal("Failed to connect to database:", err)
		}

		db.AutoMigrate(&models.Student{}, &models.Course{}, &models.Enrollment{}, &models.Instructor{})
		fmt.Println("Database connected!")
		return db
	}
	return nil
}
