package repository

import (
	"github.com/rishad004/project-gv/admin-service/internal/domain"
	"github.com/rishad004/project-gv/admin-service/utils"
)

func (r *adminRepo) Login(Email, Password string) (string, error) {
	var admin domain.Admin

	if err := r.Db.First(&admin, "email=?", Email).Error; err != nil {
		return "", err
	}

	if err := utils.VerifyPassword(Password, admin.Hashed, admin.Salted); err != nil {
		return "", err
	}

	if admin.Super {
		token, err := utils.JwtCreate(int(admin.ID), Email, "super")
		if err != nil {
			return "", err
		}
		return token, nil
	}

	token, err := utils.JwtCreate(int(admin.ID), Email, "admin")
	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *adminRepo) AddAdmin(email, password string) error {
	hash, salt, er := utils.HashPassword(password)

	if er != nil {
		return er
	}

	if err := r.Db.Create(&domain.Admin{
		Email:  email,
		Hashed: hash,
		Salted: salt,
		Super:  false,
	}).Error; err != nil {
		return err
	}

	return nil
}
