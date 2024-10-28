package delivery

import (
	"context"

	pb "github.com/rishad004/Gv_protofiles/streamer"
)

func (h *StreamerHandler) FindByStreamKey(c context.Context, req *pb.StreamKeyRequest) (*pb.StreamKeyResponse, error) {
	id, title, description, streamKey, err := h.svc.FindByStreamKey(req.Channel)
	if err != nil {
		return nil, err
	}

	return &pb.StreamKeyResponse{Id: id, Streamkey: streamKey, Title: title, Description: description}, nil
}

func (h *StreamerHandler) StreamStart(c context.Context, req *pb.StreamKeyResponse) (*pb.Empty, error) {
	if err := h.svc.StreamStart(req.Streamkey); err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (h *StreamerHandler) StreamEnd(c context.Context, req *pb.StreamKeyResponse) (*pb.Empty, error) {
	if err := h.svc.StreamEnd(req.Streamkey); err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}
