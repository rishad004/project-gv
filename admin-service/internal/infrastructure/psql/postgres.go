package psql

import (
	"fmt"

	"github.com/rishad004/project-gv/admin-service/internal/domain"
	"github.com/rishad004/project-gv/admin-service/utils"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func PsqlConn() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(viper.GetString("PSQL_URL")))

	if err != nil {
		return nil, err
	} else {
		fmt.Println("Connect to Psql")
	}

	db.AutoMigrate(&domain.Admin{})

	hash, salt, _ := utils.HashPassword("superadmin")

	db.Create(&domain.Admin{
		Email:  "superadmin@gmail.com",
		Hashed: hash,
		Salted: salt,
		Super:  true,
	})

	return db, nil
}
