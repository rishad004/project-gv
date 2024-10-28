package service

import (
	stream_pb "github.com/rishad004/Gv_protofiles/stream"
	"github.com/rishad004/project-gv/streamer-service/internal/domain"
)

type streamerService struct {
	repo      StreamerRepo
	stream_pb stream_pb.StreamServiceClient
}

func NewStreamerService(repo StreamerRepo, streamSvc stream_pb.StreamServiceClient) *streamerService {
	return &streamerService{repo: repo, stream_pb: streamSvc}
}

func (s *streamerService) RegisteringStreamer(streamer domain.Streamer) (string, error) {
	key, err := s.repo.RegisteringStreamer(streamer)
	if err != nil {
		return "", err
	}
	return key, nil
}

func (s *streamerService) ChannelView(userId int) (domain.Streamer, error) {
	streamer, err := s.repo.ChannelView(userId)

	if err != nil {
		return domain.Streamer{}, err
	}

	return streamer, nil
}

func (s *streamerService) EditChannel(streamer domain.Streamer) error {
	if err := s.repo.EditChannel(streamer); err != nil {
		return err
	}
	return nil
}

func (s *streamerService) GettingFollowed(userId, streamerId int) error {
	if err := s.repo.GettingFollowed(userId, streamerId); err != nil {
		return err
	}

	return nil
}
