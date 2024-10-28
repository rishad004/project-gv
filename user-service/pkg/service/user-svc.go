package service

import (
	"context"

	payment_pb "github.com/rishad004/Gv_protofiles/payment"
	streamer_pb "github.com/rishad004/Gv_protofiles/streamer"
	"github.com/rishad004/project-gv/user-service/internal/domain"
)

type userService struct {
	repo        UserRepo
	Streamer_Pb streamer_pb.StreamerServiceClient
	Payment_Pb  payment_pb.PaymentServiceClient
}

func NewUserService(repo UserRepo, streamer streamer_pb.StreamerServiceClient, payment payment_pb.PaymentServiceClient) *userService {
	return &userService{repo: repo, Streamer_Pb: streamer, Payment_Pb: payment}
}

func (s *userService) SignUp(user domain.Users) (string, error) {
	link, err := s.repo.CreateUser(user)

	if err != nil {
		return "", err
	}

	return link, nil
}

func (s *userService) EmailVerification(key string) error {
	if err := s.repo.EmailVerify(key); err != nil {
		return err
	}
	return nil
}

func (s *userService) Profile(id int) (domain.Users, error) {
	user, err := s.repo.Profile(id)
	if err != nil {
		return domain.Users{}, err
	}

	return user, nil
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

func (s *userService) Following(streamerId, userId int) error {
	if _, err := s.Streamer_Pb.GettingFollowed(context.Background(),
		&streamer_pb.Verification{Userid: int32(userId),
			Id: int32(streamerId)}); err != nil {
		return err
	}

	if err := s.repo.Following(streamerId, userId); err != nil {
		return err
	}

	return nil
}
