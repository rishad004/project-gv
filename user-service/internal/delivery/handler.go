package delivery

import (
	pb "github.com/rishad004/Gv_protofiles/user"
	"github.com/rishad004/project-gv/user-service/pkg/service"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	svc service.UserService
}

func NewUserHandler(svc service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}
