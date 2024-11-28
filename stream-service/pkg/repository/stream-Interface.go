package repository

import (
	"github.com/rishad004/project-gv/stream-service/internal/domain"
	"gorm.io/gorm"
)

type StreamRepo interface {
	EndStream(id int32) error
	StartStream(id int32) error
	StreamDetailing(streamerId int32, title, description string) error
	StreamDetails(streamerId int32) (domain.StreamData, error)
	StreamCount() (int, error)
}

type streamRepo struct {
	Db *gorm.DB
}

func NewStreamRepo(Db *gorm.DB) StreamRepo {
	return &streamRepo{Db: Db}
}
