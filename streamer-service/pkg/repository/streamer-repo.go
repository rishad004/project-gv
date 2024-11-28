package repository

import (
	"errors"
	"fmt"
	"strings"

	"github.com/rishad004/project-gv/streamer-service/internal/domain"
	"github.com/rishad004/project-gv/streamer-service/utils"
	"gorm.io/gorm"
)

type streamerRepo struct {
	Db *gorm.DB
}

func NewStreamerRepo(DB *gorm.DB) StreamerRepo {
	return &streamerRepo{Db: DB}
}

func (r *streamerRepo) RegisteringStreamer(streamer domain.Streamer) (string, error) {

	key := utils.StreamKeyCreate()
	streamer.StreamKey = key

	if err := r.Db.Create(&streamer).Error; err != nil {
		return "", err
	}

	return key, nil
}

func (r *streamerRepo) ChannelView(userId int) (domain.Streamer, error) {

	var streamer domain.Streamer

	if err := r.Db.First(&streamer, "user_id=?", userId).Error; err != nil {
		return domain.Streamer{}, err
	}

	return streamer, nil
}

func (r *streamerRepo) EditChannel(streamer domain.Streamer) error {
	var Streamer domain.Streamer

	if err := r.Db.First(&Streamer, "user_id=?", streamer.UserId).Error; err != nil {
		return err
	}

	Streamer.Name = streamer.Name
	Streamer.Description = streamer.Description

	if err := r.Db.Save(&Streamer).Error; err != nil {
		return err
	}

	return nil
}

func (r *streamerRepo) StreamerList() ([]domain.Streamer, error) {

	var streamers []domain.Streamer

	if err := r.Db.Find(&streamers).Error; err != nil {
		return nil, err
	}

	return streamers, nil
}

func (r *streamerRepo) FindStreamerUserId(id int) (domain.Streamer, error) {

	var streamer domain.Streamer

	if err := r.Db.First(&streamer, "user_id=?", id).Error; err != nil {
		return domain.Streamer{}, err
	}

	return streamer, nil
}

func (r *streamerRepo) GettingFollowed(userId, streamerId int) error {
	var streamer domain.Streamer

	if err := r.Db.First(&streamer, "id=?", streamerId).Error; err != nil {
		return err
	}

	if int(streamer.UserId) == userId {
		return errors.New("can't follow yourself!")
	}

	check := strings.Split(streamer.Followers, ",")

	for i := 0; i < len(check)-1; i++ {
		fmt.Println(check[i])
		fmt.Println(fmt.Sprint(userId))
		if check[i] == fmt.Sprint(userId) {
			return errors.New("already followed!")
		}
	}

	streamer.Followers += fmt.Sprint(userId, ",")

	if err := r.Db.Save(&streamer).Error; err != nil {
		return err
	}

	return nil
}
