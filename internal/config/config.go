package config

import (
	"HW1-OnlineCoursePlatform/internal/models"
	"fmt"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
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

		// Retry logic for database connection
		for i := 0; i < 5; i++ { // Retry up to 5 times
			db, err = gorm.Open(postgres.Open(dbConnector.Connector()), &gorm.Config{})
			if err == nil {
				break // Successfully connected
			}
			log.Printf("Failed to connect to database, retrying in 2 seconds...: %v", err)
			time.Sleep(2 * time.Second) // Wait before retrying
		}

		if err != nil {
			log.Fatalf("Failed to connect to database after multiple attempts: %v", err)
		}

		err = db.AutoMigrate(&models.Student{}, &models.Course{}, &models.Enrollment{}, &models.Instructor{})
		if err != nil {
			log.Fatalf("Failed to migrate database: %v", err)
		}
		fmt.Println("Database connected and migrated!")
		return db
	}
	return nil
}
