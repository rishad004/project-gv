package repository

import (
	"errors"
	"strconv"
	"time"

	"github.com/rishad004/project-gv/user-service/internal/domain"
)

func (r *userRepo) WalletAdd(paymentId string, amount, userId int32) error {
	r.Rdb.Set(paymentId, amount, 24*time.Hour)
	return nil
}

func (r *userRepo) CreateWallet(userId int) error {

	if err := r.Db.Create(&domain.Wallet{
		UserId: userId,
		Amount: 0,
	}).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepo) WalletAdded(paymentId string, userid int) error {
	var wallet domain.Wallet
	var user domain.Users

	amount, err := r.Rdb.Get(paymentId).Result()
	if err != nil {
		return err
	}
	Amount, er := strconv.Atoi(amount)
	if er != nil {
		return er
	}

	if err = r.Db.First(&wallet, "user_id=?", userid).Error; err != nil {
		return err
	}

	wallet.Amount += Amount
	if err = r.Db.Save(&wallet).Error; err != nil {
		return err
	}

	if err = r.Db.Create(&domain.WalletTransactions{
		WalletId: int(wallet.ID),
		Amount:   Amount,
		Type:     "Deposit",
	}).Error; err != nil {
		return err
	}

	if err := r.Db.First(&user, userid).Error; err != nil {
		return err
	}

	r.PushCommentToQueue("WalletAdd", domain.Kafka{Message: user.Email + "," + strconv.Itoa(Amount)})
	return nil
}

func (r *userRepo) SuperChat(Amount int, UserId int) error {
	var wallet domain.Wallet
	var user domain.Users

	if err := r.Db.First(&wallet, "user_id=?", UserId).Error; err != nil {
		return err
	}

	if Amount > wallet.Amount {
		return errors.New("insufficient balance")
	}

	wallet.Amount -= Amount
	if err := r.Db.Save(&wallet).Error; err != nil {
		return err
	}

	if err := r.Db.Create(&domain.WalletTransactions{
		WalletId: int(wallet.ID),
		Amount:   -Amount,
		Type:     "Super Chat",
	}).Error; err != nil {
		return err
	}

	if err := r.Db.First(&user).Error; err != nil {
		return err
	}

	r.PushCommentToQueue("Superchat", domain.Kafka{Message: user.Email + "," + strconv.Itoa(Amount)})

	return nil
}

func (r *userRepo) WalletShow(userId int) (int, error) {
	var wallet domain.Wallet

	if err := r.Db.First(&wallet, "user_id=?", userId).Error; err != nil {
		return 0, err
	}

	return wallet.Amount, nil
}

func (r *userRepo) SuperChatTotal() (int, error) {
	var superchat []domain.WalletTransactions
	var sum int

	if err := r.Db.Find(&superchat, "type=?", "Super Chat").Error; err != nil {
		return 0, err
	}

	for _, v := range superchat {
		sum += v.Amount
	}

	return sum, nil
}
