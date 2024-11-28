package repository

import (
	"github.com/rishad004/project-gv/streamer-service/internal/domain"
)

type StreamerRepo interface {
	RegisteringStreamer(streamer domain.Streamer) (string, error)
	ChannelView(userId int) (domain.Streamer, error)
	EditChannel(streamer domain.Streamer) error
	SubscriptionSetting(userId int, amount int) error
	SubscriptionCheck(userid int, id int) (int, int, error)
	SubscriptionList(streamers []domain.Streamer) (map[int]map[string]any, error)
	StreamerList() ([]domain.Streamer, error)
	FindStreamerUserId(id int) (domain.Streamer, error)
	FindByStreamKey(channel string) (int32, string, error)
	StreamStart(key string) (int32, error)
	StreamEnd(key string) (int32, error)
	GettingFollowed(userId, streamerId int) error
}
