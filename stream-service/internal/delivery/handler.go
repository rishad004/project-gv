package delivery

import (
	pb "github.com/rishad004/Gv_protofiles/stream"
	"github.com/rishad004/project-gv/stream-service/pkg/service"
)

type StreamHandler struct {
	pb.UnimplementedStreamServiceServer
	svc service.StreamService
}

func NewStreamHandler(svc service.StreamService) *StreamHandler {
	return &StreamHandler{svc: svc}
}
