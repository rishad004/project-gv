package di

import (
	"log"
	"net"

	pb "github.com/rishad004/Gv_protofiles/admin"
	"github.com/rishad004/project-gv/admin-service/internal/delivery"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func InitAdmin(handler *delivery.AdminHandler) error {

	g := grpc.NewServer()
	pb.RegisterAdminServiceServer(g, handler)

	listen, err := net.Listen("tcp", viper.GetString("PORT"))
	if err != nil {
		return err
	}

	log.Println("admin-service server listening on port :8085")
	if err := g.Serve(listen); err != nil {
		return err
	}

	return nil
}
