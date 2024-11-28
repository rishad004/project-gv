package delivery

import (
	"context"
	"log"

	pb "github.com/rishad004/Gv_protofiles/user"
)

func (h *UserHandler) Subscribing(c context.Context, req *pb.SubscribeRequest) (*pb.SubscribeResponse, error) {

	log.Println("")
	log.Println("----------------Subscribing-------------")

	paymentId, err := h.svc.Subscribing(int(req.Userid), int(req.Id))

	if err != nil {
		return nil, err
	}

	return &pb.SubscribeResponse{Paymentid: paymentId, Message: "Order created successfully, please complete payment!"}, nil
}

func (h *UserHandler) Subscribed(c context.Context, req *pb.SubscribedRequest) (*pb.VerificationResponse, error) {

	log.Println("")
	log.Println("----------------Subscribed-------------")

	if err := h.svc.Subscribed(int(req.Id), req.Paymentid); err != nil {
		return nil, err
	}

	return &pb.VerificationResponse{Message: "Subscribed successfully!"}, nil
}
