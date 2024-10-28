package di

import (
	"log"
	"net"

	pb "github.com/rishad004/Gv_protofiles/streamer"
	"github.com/rishad004/project-gv/streamer-service/internal/delivery"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func InitStreamer(handler *delivery.StreamerHandler) error {

	g := grpc.NewServer()
	pb.RegisterStreamerServiceServer(g, handler)

	listen, err := net.Listen("tcp", viper.GetString("PORT"))
	if err != nil {
		return err
	}

	log.Println("streamer-service server listening on port :8082")
	if err := g.Serve(listen); err != nil {
		return err
	}

	return nil
}
