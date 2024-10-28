package repository

import (
	"errors"
	"fmt"

	"github.com/rishad004/project-gv/streamer-service/internal/domain"
)

func (r *streamerRepo) SubscriptionSetting(userId int, amount int) error {
	var streamer domain.Streamer
	var subscription domain.Subscription

	if err := r.Db.First(&streamer, "user_id=?", userId).Error; err != nil {
		return err
	}

	if err := r.Db.First(&subscription, "streamer_id=?", streamer.ID).Error; err != nil {

		subscription.Amount = amount
		subscription.StreamerId = int(streamer.ID)

		if er := r.Db.Create(&subscription).Error; er != nil {
			return er
		}
	}

	subscription.Amount = amount
	if err := r.Db.Save(&subscription).Error; err != nil {
		return err
	}

	return nil
}

func (r *streamerRepo) SubscriptionCheck(userid int, id int) (int, int, error) {
	var Subscription domain.Subscription

	if err := r.Db.First(&Subscription, "streamer_id=?", id).Error; err != nil {
		return 0, 0, err
	}

	streamer, _ := r.FindStreamerUserId(userid)

	if streamer.ID == uint(Subscription.StreamerId) {
		return 0, 0, errors.New("can't subscribe yourself!")
	}

	return Subscription.Amount, int(Subscription.ID), nil
}

func (r *streamerRepo) SubscriptionList(streamers []domain.Streamer) (map[int]map[string]any, error) {

	var subscriptions []domain.Subscription
	res := make(map[int]map[string]any)

	if err := r.Db.Find(&subscriptions).Error; err != nil {
		return nil, err
	}

	for _, v := range streamers {
		res[int(v.ID)] = map[string]any{
			"streamer_Id":         v.ID,
			"channel_Name":        v.Name,
			"subscription_Id":     0,
			"subscription_Amount": "",
		}
	}

	for _, v := range subscriptions {
		data := res[v.StreamerId]
		data["subscription_Id"] = int(v.ID)
		data["subscription_Amount"] = fmt.Sprint(v.Amount, "rs")
		res[v.StreamerId] = data
	}

	return res, nil
}
