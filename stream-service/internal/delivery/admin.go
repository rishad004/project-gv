package delivery

import (
	"context"

	pb "github.com/rishad004/Gv_protofiles/stream"
)

func (h *StreamHandler) StreamerCount(c context.Context, req *pb.Empty) (*pb.Stream, error) {
	count, err := h.svc.StreamCount()

	if err != nil {
		return nil, err
	}

	return &pb.Stream{StreamerId: int32(count)}, nil
}
