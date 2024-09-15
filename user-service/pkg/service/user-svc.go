package service

import "github.com/rishad004/project-gv/user-service/internal/domain"

type userService struct {
	repo UserRepo
}

func NewUserService(repo UserRepo) *userService {
	return &userService{repo: repo}
}

func (s *userService) SignUp(user domain.Users) (int, error) {
	id, err := s.repo.CreateUser(user)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *userService) ProfileEditing(edits domain.Users) error {
	if err := s.repo.ProfileEditing(edits); err != nil {
		return err
	}
	return nil
}

func (s *userService) Login(Username, Password string) (string, error) {
	token, err := s.repo.Login(Username, Password)
	if err != nil {
		return "", err
	}
	return token, nil
}
