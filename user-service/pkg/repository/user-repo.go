package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/rishad004/project-gv/user-service/internal/domain"
	"github.com/rishad004/project-gv/user-service/utils"
	"gorm.io/gorm"
)

type userRepo struct {
	Db  *gorm.DB
	Rdb *redis.Client
}

func NewUserRepo(DB *gorm.DB, RDB *redis.Client) UserRepo {
	return &userRepo{Db: DB, Rdb: RDB}
}

func (r *userRepo) CreateUser(user domain.Users) (string, error) {
	var err error

	user.Hashed, user.Salted, err = utils.HashPassword(user.Hashed)
	if err != nil {
		return "", err
	}

	if err = r.Db.Create(&user).Error; err != nil {
		return "", err
	}

	key := utils.StringKey()
	r.Rdb.Set(key, user.ID, 24*time.Hour)

	return "localhost:8080/verify?code=" + key, nil
}

func (r *userRepo) EmailVerify(key string) error {

	var user domain.Users

	id, err := r.Rdb.Get(key).Result()
	if err != nil {
		return err
	}

	if er := r.Db.First(&user, id).Error; er != nil {
		return er
	}

	user.Verified = true
	if er := r.Db.Save(&user).Error; er != nil {
		return er
	}

	return nil
}

func (r *userRepo) Login(Username, Password string) (string, error) {
	var user domain.Users

	if err := r.Db.First(&user, "username=?", Username).Error; err != nil {
		return "", err
	}

	if err := utils.VerifyPassword(Password, user.Hashed, user.Salted); err != nil {
		return "", err
	}

	if !user.Verified {
		return "", errors.New("Email not verified. Please check your mail!")
	}

	token, err := utils.JwtCreate(int(user.ID), "user")
	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *userRepo) Profile(id int) (domain.Users, error) {
	var user domain.Users

	if err := r.Db.First(&user, id).Error; err != nil {
		return domain.Users{}, err
	}

	return user, nil
}

func (r *userRepo) ProfileEditing(edits domain.Users) error {
	var user domain.Users

	if err := r.Db.First(&user, edits.ID).Error; err != nil {
		return err
	}

	user.Username = edits.Username
	user.Phone = edits.Phone
	user.Gender = edits.Gender

	if err := r.Db.Save(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepo) Following(streamerId, userId int) error {
	var user domain.Users
	if err := r.Db.First(&user, "id=?", userId).Error; err != nil {
		return err
	}

	user.Following += fmt.Sprint(streamerId, ",")

	if err := r.Db.Save(&user).Error; err != nil {
		return err
	}

	return nil
}
