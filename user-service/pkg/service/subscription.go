package service

import (
	"context"
	"errors"

	payment_pb "github.com/rishad004/Gv_protofiles/payment"
	streamer_pb "github.com/rishad004/Gv_protofiles/streamer"
)

func (s *userService) Subscribing(userid int, id int) (string, error) {
	ress, errs := s.Streamer_Pb.SubscriptionCheck(context.Background(), &streamer_pb.Verification{Id: int32(id), Userid: int32(userid)})

	if errs != nil {
		return "", errs
	}

	if err := s.repo.SubscriptionCheck(userid, int(ress.Sid)); err == nil {
		return "", errors.New("subscription already exists")
	}

	resp, errp := s.Payment_Pb.PaymentInitialize(context.Background(), &payment_pb.PaymentInitRequest{Amount: ress.Amount, Type: "Subscription"})

	if errp != nil {
		return "", errp
	}

	if err := s.repo.Subscribing(resp.Paymentid, int(ress.Sid), int(ress.Amount)); err != nil {
		return "", err
	}

	return resp.Paymentid, nil
}

func (r *userService) Subscribed(userid int, PaymentId string) error {
	if err := r.repo.Subscribed(userid, PaymentId); err != nil {
		return err
	}

	return nil
}
