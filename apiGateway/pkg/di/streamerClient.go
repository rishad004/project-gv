package di

import (
	"log"

	"github.com/gorilla/mux"
	streamer_pb "github.com/rishad004/Gv_protofiles/streamer"
	"github.com/rishad004/project-gv/apiGateway/inertnal/handler"
	"github.com/rishad004/project-gv/apiGateway/inertnal/routers"
	"github.com/rishad004/project-gv/apiGateway/pkg/middleware"
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

func StreamerRouting(r *mux.Router, handle *handler.ApiHanlder) {

	streamerRouter := r.PathPrefix("/").Subrouter()
	streamerRouter.Use(middleware.MiddlewareU)
	routers.StreamerRouter(handle, streamerRouter)

}
