package service

import (
	"context"

	streamer_pb "github.com/rishad004/Gv_protofiles/streamer"
	"github.com/rishad004/project-gv/stream-service/internal/domain"
)


func (s *streamSvc) StartStream(id int32) error {
	if err := s.repo.StartStream(id); err != nil {
		return err
	}

	return nil
}

func (s *streamSvc) EndStream(id int32) error {
	if err := s.repo.EndStream(id); err != nil {
		return err
	}

	return nil
}

func (s *streamSvc) StreamDetailing(userId int32, title, description string) error {
	res, err := s.streamerSvc.ChannelView(context.Background(), &streamer_pb.Verification{
		Id:     userId,
		Userid: userId,
	})

	if err != nil {
		return err
	}

	if er := s.repo.StreamDetailing(res.Id, title, description); er != nil {
		return er
	}

	return nil
}

func (s *streamSvc) StreamDetails(streamerId int32) (domain.StreamData, error) {
	stream, err := s.repo.StreamDetails(streamerId)

	if err != nil {
		return domain.StreamData{}, err
	}

	return stream, nil
}
