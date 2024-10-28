package service

import (
	"context"
	"fmt"
	"log"

	user_pb "github.com/rishad004/Gv_protofiles/user"
)

type paymentSvc struct {
	repo    PaymentRepo
	UserSvc user_pb.UserServiceClient
}

func NewPaymentService(svc PaymentRepo, userSvc user_pb.UserServiceClient) PaymentService {
	return &paymentSvc{repo: svc, UserSvc: userSvc}
}

func (s *paymentSvc) PaymentInitialize(Amount int, Type string) (int, string, error) {
	id, paymentId, err := s.repo.PaymentInitialize(Amount, Type)
	if err != nil {
		return 0, "", err
	}

	return id, paymentId, nil
}

func (s *paymentSvc) PaymentVerifying(id int, Sig, Ord, Pay string) error {
	err, Type := s.repo.PaymentVerifying(Sig, Ord, Pay)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	if Type == "Subscription" {
		log.Println(Type)
		if _, err := s.UserSvc.Subscribed(context.Background(), &user_pb.SubscribedRequest{Id: int32(id)}); err != nil {
			return err
		}
	}

	return nil
}
