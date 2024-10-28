package service

import "github.com/rishad004/project-gv/stream-service/internal/domain"

type StreamRepo interface {
	// EndStream(id int32) error
	// StartStream(id int32) error
	StreamDetailing(streamerId int32, title, description string) error
	StreamDetails(streamerId int32) (domain.Stream, error)
}

type StreamService interface {
	// EndStream(id int32) error
	// StartStream(id int32) error
	StreamDetailing(userId int32, title, description string) error
	StreamDetails(streamerId int32) (domain.Stream, error)
}
