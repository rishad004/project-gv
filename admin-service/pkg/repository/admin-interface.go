package repository

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type AdminRepo interface {
	Login(Email, Password string) (string, error)
	AddAdmin(email, password string) error
}

type adminRepo struct {
	Db  *gorm.DB
	Rdb *redis.Client
}

func NewAdminRepo(DB *gorm.DB, RDB *redis.Client) AdminRepo {
	return &adminRepo{Db: DB, Rdb: RDB}
}
