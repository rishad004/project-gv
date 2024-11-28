package di

import (
	"log"

	"github.com/gorilla/mux"
	user_pb "github.com/rishad004/Gv_protofiles/user"
	"github.com/rishad004/project-gv/apiGateway/inertnal/handler"
	"github.com/rishad004/project-gv/apiGateway/inertnal/routers"
	"github.com/rishad004/project-gv/apiGateway/pkg/middleware"
	"google.golang.org/grpc"
)

func InitUser() (*grpc.ClientConn, user_pb.UserServiceClient) {

	connUser, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Failed to connect to user service:", err)
	}

	userSvc := user_pb.NewUserServiceClient(connUser)

	return connUser, userSvc
}

func UserRouting(r *mux.Router, handle *handler.ApiHanlder) {

	userMiddle := r.PathPrefix("/").Subrouter()
	userMiddle.Use(middleware.MiddlewareU)
	routers.UserMiddle(handle, userMiddle)

	userRouter := r.PathPrefix("/").Subrouter()
	routers.UserRouter(handle, userRouter)

}
