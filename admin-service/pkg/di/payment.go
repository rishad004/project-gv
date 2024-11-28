package di

import (
	"log"

	payment_pb "github.com/rishad004/Gv_protofiles/payment"
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
