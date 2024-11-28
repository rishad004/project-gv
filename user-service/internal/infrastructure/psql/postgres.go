package psql

import (
	"fmt"

	"github.com/rishad004/project-gv/user-service/internal/domain"
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

	db.AutoMigrate(&domain.Users{}, domain.Subscribed{}, domain.Wallet{}, &domain.WalletTransactions{})

	return db, nil
}
