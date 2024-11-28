package service

import (
	streamer_pb "github.com/rishad004/Gv_protofiles/streamer"
	"github.com/rishad004/project-gv/stream-service/internal/domain"
)

type StreamRepo interface {
	EndStream(id int32) error
	StartStream(id int32) error
	StreamDetailing(streamerId int32, title, description string) error
	StreamDetails(streamerId int32) (domain.StreamData, error)
	StreamCount() (int, error)
}

type StreamService interface {
	EndStream(id int32) error
	StartStream(id int32) error
	StreamDetailing(userId int32, title, description string) error
	StreamDetails(streamerId int32) (domain.StreamData, error)
	StreamCount() (int, error)
}

type streamSvc struct {
	repo        StreamRepo
	streamerSvc streamer_pb.StreamerServiceClient
}

func NewStreamService(svc StreamRepo, streamerSvc streamer_pb.StreamerServiceClient) StreamService {
	return &streamSvc{repo: svc, streamerSvc: streamerSvc}
}
