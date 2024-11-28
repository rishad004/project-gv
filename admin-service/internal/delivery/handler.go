package delivery

import (
	pb "github.com/rishad004/Gv_protofiles/admin"
	"github.com/rishad004/project-gv/admin-service/pkg/service"
)

type AdminHandler struct {
	pb.UnimplementedAdminServiceServer
	svc service.AdminService
}

func NewAdminHandler(svc service.AdminService) *AdminHandler {
	return &AdminHandler{svc: svc}
}
