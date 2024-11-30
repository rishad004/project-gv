package di

import (
	"log"

	user_pb "github.com/rishad004/Gv_protofiles/user"
	"google.golang.org/grpc"
)

func InitUser() (*grpc.ClientConn, user_pb.UserServiceClient) {
	connUser, err := grpc.Dial("user-service:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Failed to connect to user service:", err)
	}

	userSvc := user_pb.NewUserServiceClient(connUser)

	return connUser, userSvc
}
