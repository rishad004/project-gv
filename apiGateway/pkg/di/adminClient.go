package di

import (
	"log"

	"github.com/gorilla/mux"
	admin_pb "github.com/rishad004/Gv_protofiles/admin"
	"github.com/rishad004/project-gv/apiGateway/inertnal/handler"
	"github.com/rishad004/project-gv/apiGateway/inertnal/routers"
	"github.com/rishad004/project-gv/apiGateway/pkg/middleware"
	"google.golang.org/grpc"
)

func InitAdmin() (*grpc.ClientConn, admin_pb.AdminServiceClient) {

	connAdmin, err := grpc.Dial("admin-service:8085", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Failed to connect to admin service:", err)
	}

	adminSvc := admin_pb.NewAdminServiceClient(connAdmin)

	return connAdmin, adminSvc
}

func AdminRouting(r *mux.Router, handle *handler.ApiHanlder) {

	adminMiddle := r.PathPrefix("/admin").Subrouter()
	adminMiddle.Use(middleware.MiddlewareA)
	routers.AdminMiddle(handle, adminMiddle)

	adminRouter := r.PathPrefix("/admin").Subrouter()
	routers.AdminRouter(handle, adminRouter)

}
