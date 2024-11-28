package di

import (
	"log"

	"github.com/gorilla/mux"
	stream_pb "github.com/rishad004/Gv_protofiles/stream"
	"github.com/rishad004/project-gv/apiGateway/inertnal/handler"
	"github.com/rishad004/project-gv/apiGateway/inertnal/routers"
	"github.com/rishad004/project-gv/apiGateway/pkg/middleware"
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

func StreamRouting(r *mux.Router, handle *handler.ApiHanlder) {

	streamMiddle := r.PathPrefix("/").Subrouter()
	streamMiddle.Use(middleware.MiddlewareU)
	routers.StreamMiddle(handle, streamMiddle)

	streamRouter := r.PathPrefix("/").Subrouter()
	routers.StreamRouter(handle, streamRouter)
}
