package repository

import (
	"github.com/rishad004/project-gv/user-service/internal/domain"
	"github.com/rishad004/project-gv/user-service/utils"
	"gorm.io/gorm"
)

type userRepo struct {
	Db *gorm.DB
}

func NewUserRepo(DB *gorm.DB) UserRepo {
	return &userRepo{Db: DB}
}

func (r *userRepo) CreateUser(user domain.Users) (int, error) {
	var err error

	user.Hashed, user.Salted, err = utils.HashPassword(user.Hashed)
	if err != nil {
		return 0, err
	}

	if err = r.Db.Create(&user).Error; err != nil {
		return 0, err
	}
	return int(user.ID), nil
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

func (r *userRepo) Login(Username, Password string) (string, error) {
	var user domain.Users

	if err := r.Db.First(&user, "username=?", Username).Error; err != nil {
		return "", err
	}

	if err := utils.VerifyPassword(Password, user.Hashed, user.Salted); err != nil {
		return "", err
	}

	token, err := utils.JwtCreate(int(user.ID), "user")
	if err != nil {
		return "", err
	}

	return token, nil
}
