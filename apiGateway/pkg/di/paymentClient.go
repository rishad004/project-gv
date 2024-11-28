package di

import (
	"log"

	"github.com/gorilla/mux"
	payment_pb "github.com/rishad004/Gv_protofiles/payment"
	"github.com/rishad004/project-gv/apiGateway/inertnal/handler"
	"github.com/rishad004/project-gv/apiGateway/inertnal/routers"
	"github.com/rishad004/project-gv/apiGateway/pkg/middleware"
	"google.golang.org/grpc"
)

func InitPayment() (*grpc.ClientConn, payment_pb.PaymentServiceClient) {

	connPayment, err := grpc.Dial("localhost:8083", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Failed to connect to user service:", err)
	}

	paymentSvc := payment_pb.NewPaymentServiceClient(connPayment)

	return connPayment, paymentSvc
}

func PaymentRouting(r *mux.Router, handle *handler.ApiHanlder) {

	paymentRouter := r.PathPrefix("/payment").Subrouter()
	paymentRouter.Use(middleware.MiddlewareU)
	routers.PaymentRouter(handle, paymentRouter)

}
