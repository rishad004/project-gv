package delivery

import (
	"context"

	pb "github.com/rishad004/Gv_protofiles/user"
	"github.com/rishad004/project-gv/user-service/internal/domain"
	"github.com/rishad004/project-gv/user-service/pkg/service"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	svc service.UserService
}

func NewUserHandler(svc service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

func (h *UserHandler) SignUp(c context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	user := domain.Users{
		Username: req.Username,
		Email:    req.Email,
		Phone:    req.Phone,
		Hashed:   req.Password,
		Gender:   req.Gender,
		Verified: false,
	}

	id, err := h.svc.SignUp(user)

	if err != nil {
		return nil, err
	}

	return &pb.SignUpResponse{Id: int32(id), Message: "User signed up successfully!"}, nil
}

func (h *UserHandler) Login(c context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	token, err := h.svc.Login(req.Username, req.Password)

	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{Token: token, Message: "Logged in successfully!"}, nil
}

func (h *UserHandler) ProfileEditing(c context.Context, req *pb.EditRequest) (*pb.EditResponse, error) {
	edits := domain.Users{
		Username: req.Username,
		Phone:    req.Phone,
		Gender:   req.Gender,
	}
	edits.ID = uint(req.Id)

	if err := h.svc.ProfileEditing(edits); err != nil {
		return nil, err
	}
	return &pb.EditResponse{Message: "Profile edited successfully!"}, nil
}
