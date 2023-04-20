package config

import (
	"fmt"
	"mygram/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Gorm struct {

	// db config
	Username string
	Password string
	Port     string
	Address  string
	Database string

	// db connection
	DB *gorm.DB
}

type GormDb struct {
	*Gorm
}

var (
	GORM *GormDb
)

func InitGorm() (*gorm.DB, error) {
	GORM = new(GormDb)

	GORM.Gorm = &Gorm{
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Address:  os.Getenv("POSTGRES_ADDRESS"),
		Database: os.Getenv("POSTGRES_DB"),
	}

	db, err := GORM.Gorm.OpenConnection()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (p *Gorm) OpenConnection() (*gorm.DB, error) {

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", p.Address, p.Port, p.Username, p.Password, p.Database)

	dbConnection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "my_gram.",
			SingularTable: false,
		},
	})

	if err != nil {
		panic("Failed to connect to Database")
	}

	// ping connection db
	p.DB = dbConnection

	err = p.DB.Debug().AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.SocialMedia{})
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to database using GORM")

	return dbConnection, nil
}
