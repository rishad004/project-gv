package delivery

import (
	"context"

	pb "github.com/rishad004/Gv_protofiles/streamer"
	"github.com/rishad004/project-gv/streamer-service/internal/domain"
	"github.com/rishad004/project-gv/streamer-service/pkg/service"
)

type StreamerHandler struct {
	pb.UnimplementedStreamerServiceServer
	svc service.StreamerService
}

func NewStreamerHandler(svc service.StreamerService) *StreamerHandler {
	return &StreamerHandler{svc: svc}
}

func (h *StreamerHandler) Registration(c context.Context, req *pb.RegistrationRequest) (*pb.RegistrationResponse, error) {
	key, err := h.svc.RegisteringStreamer(domain.Streamer{
		Name:        req.Name,
		Description: req.Description,
		UserId:      uint(req.Userid),
	})
	if err != nil {
		return nil, err
	}
	return &pb.RegistrationResponse{Message: "Channel created successfully!", Streamkey: key}, nil
}

func (h *StreamerHandler) ChannelView(c context.Context, req *pb.Verification) (*pb.ChannelResponse, error) {

	streamer, err := h.svc.ChannelView(int(req.Id))

	if err != nil {
		return nil, err
	}

	return &pb.ChannelResponse{Id: int32(streamer.ID), Name: streamer.Name, Description: streamer.Description, Streamkey: streamer.StreamKey}, nil
}

func (h *StreamerHandler) EditChannel(c context.Context, req *pb.EditRequest) (*pb.EditResponse, error) {
	if err := h.svc.EditChannel(domain.Streamer{
		Name:        req.Name,
		Description: req.Description,
		UserId:      uint(req.Userid),
	}); err != nil {
		return nil, err
	}

	return &pb.EditResponse{Message: "Channel edited successfully!"}, nil
}

func (h *StreamerHandler) GettingFollowed(c context.Context, req *pb.Verification) (*pb.Empty, error) {

	if err := h.svc.GettingFollowed(int(req.Userid), int(req.Id)); err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}
