package service

import (
	"context"

	paymentpb "github.com/rishad004/Gv_protofiles/payment"
)

func (s *userService) WalletAdd(amount, userid int32) (string, error) {

	res, err := s.Payment_Pb.PaymentInitialize(context.Background(), &paymentpb.PaymentInitRequest{
		Amount: amount,
		Type:   "Wallet",
	})

	if err != nil {
		return "", err
	}

	if err = s.repo.WalletAdd(res.Paymentid, amount, userid); err != nil {
		return "", err
	}
	return res.Paymentid, nil
}

func (s *userService) WalletAdded(paymentId string, userId int) error {
	if err := s.repo.WalletAdded(paymentId, userId); err != nil {
		return err
	}
	return nil
}

func (s *userService) SuperChat(Amount int, UserId int) error {
	if err := s.repo.SuperChat(Amount, UserId); err != nil {
		return err
	}
	return nil
}

func (s *userService) WalletShow(userId int) (int, error) {
	amount, err := s.repo.WalletShow(userId)

	if err != nil {
		return 0, err
	}

	return amount, nil
}

func (s *userService) SuperChatTotal() (int, error) {
	sum, err := s.repo.SuperChatTotal()

	if err != nil {
		return 0, err
	}

	return sum, nil
}
