package repository

import "github.com/rishad004/project-gv/user-service/internal/domain"

type UserRepo interface {
	CreateUser(user domain.Users) (string, error)
	ProfileEditing(edits domain.Users) error
	Login(Username, Password string) (string, error)
	EmailVerify(key string) (int, error)
	Profile(id int) (domain.Users, error)
	Subscribing(paymentId string, sid, amount int) error
	SubscriptionCheck(userid int, id int) error
	Subscribed(userid int, PaymentId string) error
	Following(streamerId, userId int) error
	WalletAdd(paymentId string, amount, userId int32) error
	WalletAdded(paymentId string, userid int) error
	SuperChat(Amount int, UserId int) error
	WalletShow(userId int) (int, error)
	CreateWallet(userId int) error
	SuperChatTotal() (int, error)
}
