package di

import (
	"log"

	stream_pb "github.com/rishad004/Gv_protofiles/stream"
	"google.golang.org/grpc"
)

func InitStream() (*grpc.ClientConn, stream_pb.StreamServiceClient) {

	connStream, err := grpc.Dial("localhost:8084", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Failed to connect to user service:", err)
	}

	streamSvc := stream_pb.NewStreamServiceClient(connStream)

	return connStream, streamSvc
}
