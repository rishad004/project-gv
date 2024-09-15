package service

import "github.com/rishad004/project-gv/user-service/internal/domain"

type UserRepo interface {
	CreateUser(user domain.Users) (int, error)
	ProfileEditing(edits domain.Users) error
	Login(Username, Password string) (string, error)
}

type UserService interface {
	SignUp(user domain.Users) (int, error)
	ProfileEditing(edits domain.Users) error
	Login(Username, Password string) (string, error)
}
