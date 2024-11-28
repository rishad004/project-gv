package delivery

import (
	"context"

	pb "github.com/rishad004/Gv_protofiles/stream"
)

func (h *StreamHandler) StartStream(c context.Context, req *pb.Stream) (*pb.Empty, error) {

	if err := h.svc.StartStream(req.StreamerId); err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (h *StreamHandler) EndStream(c context.Context, req *pb.Stream) (*pb.Empty, error) {

	if err := h.svc.EndStream(req.StreamerId); err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (h *StreamHandler) StreamDetailing(c context.Context, req *pb.Data) (*pb.Empty, error) {

	if err := h.svc.StreamDetailing(req.Id, req.Title, req.Description); err != nil {
		return &pb.Empty{}, err
	}

	return &pb.Empty{}, nil
}

func (h *StreamHandler) StreamDetails(c context.Context, req *pb.Stream) (*pb.Data, error) {

	stream, err := h.svc.StreamDetails(req.StreamerId)
	if err != nil {
		return &pb.Data{Id: req.StreamerId, Title: "Welcome to my stream", Description: "Welcome to my stream"}, nil
	}

	return &pb.Data{Id: int32(stream.StreamerId), Title: stream.Title, Description: stream.Description}, nil
}
