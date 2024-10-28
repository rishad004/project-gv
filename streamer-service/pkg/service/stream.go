package service

import (
	"context"

	stream_pb "github.com/rishad004/Gv_protofiles/stream"
)

func (s *streamerService) FindByStreamKey(channel string) (int32, string, string, string, error) {
	id, streamKey, err := s.repo.FindByStreamKey(channel)
	if err != nil {
		return 0, "", "", "", err
	}

	res, er := s.stream_pb.StreamDetails(context.Background(), &stream_pb.Stream{StreamerId: id})
	if er != nil {
		return id, "", "", streamKey, er
	}

	return id, res.Title, res.Description, streamKey, nil
}

func (s *streamerService) StreamStart(key string) error {
	id, err := s.repo.StreamStart(key)
	if err != nil {
		return err
	}

	if _, er := s.stream_pb.StartStream(context.Background(), &stream_pb.Stream{StreamerId: id}); er != nil {
		return er
	}

	return nil
}

func (s *streamerService) StreamEnd(key string) error {
	id, err := s.repo.StreamStart(key)
	if err != nil {
		return err
	}

	if _, er := s.stream_pb.EndStream(context.Background(), &stream_pb.Stream{StreamerId: id}); er != nil {
		return er
	}

	return nil
}
