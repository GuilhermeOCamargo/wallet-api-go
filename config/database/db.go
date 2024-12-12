package database

import (
	"log"

	"github.com/GuilhermeOCamargo/go-wallet-api/repository/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDb() {
	connectionString := "host=localhost user=root password=root dbname=wallet port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	DB.AutoMigrate(&entities.Owner{})

	DB.AutoMigrate(&entities.Wallet{})
}
