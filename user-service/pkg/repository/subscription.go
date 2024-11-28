package repository

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/rishad004/project-gv/user-service/internal/domain"
)

func (r *userRepo) Subscribing(paymentId string, sid, amount int) error {
	Sid := strconv.Itoa(sid)
	Amount := strconv.Itoa(amount)

	r.Rdb.Set(paymentId, Sid+","+Amount, 24*time.Hour)
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
	return errors.New("no subscription exist")
}

func (r *userRepo) Subscribed(userid int, PaymentId string) error {
	var user domain.Users

	id, err := r.Rdb.Get(PaymentId).Result()
	if err != nil {
		return err
	}

	s := strings.Split(id, ",")

	Id, er := strconv.Atoi(s[0])
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

	if err = r.Db.First(&user, userid).Error; err != nil {
		return err
	}

	r.PushCommentToQueue("Subscription", domain.Kafka{Message: user.Email + "," + s[1]})

	return nil
}
