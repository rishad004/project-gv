package service

import (
	payment_pb "github.com/rishad004/Gv_protofiles/payment"
	streamer_pb "github.com/rishad004/Gv_protofiles/streamer"
)

type adminService struct {
	repo        AdminRepo
	Streamer_Pb streamer_pb.StreamerServiceClient
	Payment_Pb  payment_pb.PaymentServiceClient
}

func NewAdminService(repo AdminRepo, streamer streamer_pb.StreamerServiceClient, payment payment_pb.PaymentServiceClient) *adminService {
	return &adminService{repo: repo, Streamer_Pb: streamer, Payment_Pb: payment}
}

func (s *adminService) Login(Email, Password string) (string, error) {
	token, err := s.repo.Login(Email, Password)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *adminService) AddAdmin(email, password string) error {
	if err := s.repo.AddAdmin(email, password); err != nil {
		return err
	}

	return nil
}
