package di

import (
	"log"
	"net"

	pb "github.com/rishad004/Gv_protofiles/payment"
	"github.com/rishad004/project-gv/payment-service/internal/delivery"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func InitPayment(handler *delivery.PaymentHandler) error {

	g := grpc.NewServer()
	pb.RegisterPaymentServiceServer(g, handler)

	listen, err := net.Listen("tcp", viper.GetString("PORT"))
	if err != nil {
		return err
	}

	log.Println("payment-service server listening on port :8083")
	if err := g.Serve(listen); err != nil {
		return err
	}

	return nil
}
