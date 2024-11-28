package delivery

import (
	"context"

	pb "github.com/rishad004/Gv_protofiles/admin"
)

func (h *AdminHandler) Login(c context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {

	token, err := h.svc.Login(req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{Token: token}, nil
}

func (h *AdminHandler) AddAdmin(c context.Context, req *pb.LoginRequest) (*pb.Empty, error) {
	if err := h.svc.AddAdmin(req.Email, req.Password); err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}
