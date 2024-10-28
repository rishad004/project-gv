package repository

import "github.com/rishad004/project-gv/user-service/internal/domain"

type UserRepo interface {
	CreateUser(user domain.Users) (string, error)
	ProfileEditing(edits domain.Users) error
	Login(Username, Password string) (string, error)
	EmailVerify(key string) error
	Profile(id int) (domain.Users, error)
	Subscribing(paymentId string, sid int) error
	SubscriptionCheck(userid int, id int) error
	Subscribed(userid int, PaymentId string) error
	Following(streamerId, userId int) error
}
