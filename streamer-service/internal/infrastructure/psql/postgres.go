package psql

import (
	"fmt"

	"github.com/rishad004/project-gv/streamer-service/internal/domain"
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

	db.AutoMigrate(&domain.Streamer{},domain.Subscription{})

	return db, nil
}
