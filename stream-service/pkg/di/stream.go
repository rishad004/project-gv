package di

import (
	"log"
	"net"

	pb "github.com/rishad004/Gv_protofiles/stream"
	"github.com/rishad004/project-gv/stream-service/internal/delivery"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func InitStream(handler *delivery.StreamHandler) error {

	g := grpc.NewServer()
	pb.RegisterStreamServiceServer(g, handler)

	listen, err := net.Listen("tcp", viper.GetString("PORT"))
	if err != nil {
		return err
	}

	log.Println("stream-service server listening on port :8084")
	if err := g.Serve(listen); err != nil {
		return err
	}

	return nil
}
