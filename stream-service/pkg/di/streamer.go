package di

import (
	"log"

	streamer_pb "github.com/rishad004/Gv_protofiles/streamer"
	"google.golang.org/grpc"
)

func InitStreamer() (*grpc.ClientConn, streamer_pb.StreamerServiceClient) {

	connStreamer, err := grpc.Dial("streamer-service:8082", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Failed to connect to user service:", err)
	}

	streamerSvc := streamer_pb.NewStreamerServiceClient(connStreamer)

	return connStreamer, streamerSvc
}
