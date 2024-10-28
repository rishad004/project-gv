package di

import (
	"log"
	"net"

	pb "github.com/rishad004/Gv_protofiles/user"
	"github.com/rishad004/project-gv/user-service/internal/delivery"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func InitUser(handler *delivery.UserHandler) error {

	g := grpc.NewServer()
	pb.RegisterUserServiceServer(g, handler)

	listen, err := net.Listen("tcp", viper.GetString("PORT"))
	if err != nil {
		return err
	}

	log.Println("user-service server listening on port :8081")
	if err := g.Serve(listen); err != nil {
		return err
	}

	return nil
}
