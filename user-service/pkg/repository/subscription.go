package repository

import (
	"errors"
	"strconv"
	"time"

	"github.com/rishad004/project-gv/user-service/internal/domain"
)

func (r *userRepo) Subscribing(paymentId string, sid int) error {

	r.Rdb.Set(paymentId, sid, 24*time.Hour)
	return nil
}

func (r *userRepo) SubscriptionCheck(userid int, id int) error {
	var Subscribed []domain.Subscribed

	if err := r.Db.Where("user_id=? AND subscription_id=?", userid, id).Find(&Subscribed).Error; err != nil {
		return err
	}

	for _, v := range Subscribed {
		if time.Now().Before(v.Expiry) {
			return nil
		}
	}
	return errors.New("no subscription exist!")
}

func (r *userRepo) Subscribed(userid int, PaymentId string) error {
	id, err := r.Rdb.Get(PaymentId).Result()
	if err != nil {
		return err
	}

	Id, er := strconv.Atoi(id)
	if er != nil {
		return er
	}

	if errr := r.Db.Create(&domain.Subscribed{
		UserId:         userid,
		SubscriptionId: Id,
		PaymentId:      PaymentId,
		Expiry:         time.Now().AddDate(0, 0, 30),
	}).Error; errr != nil {
		return errr
	}

	return nil
}
