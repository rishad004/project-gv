package delivery

import (
	"context"

	pb "github.com/rishad004/Gv_protofiles/payment"
	"github.com/rishad004/project-gv/payment-service/pkg/service"
)

type PaymentHandler struct {
	pb.UnimplementedPaymentServiceServer
	svc service.PaymentService
}

func NewPaymentHandler(svc service.PaymentService) *PaymentHandler {
	return &PaymentHandler{svc: svc}
}

func (h *PaymentHandler) PaymentInitialize(c context.Context, req *pb.PaymentInitRequest) (*pb.PaymentInitResponse, error) {
	id, paymentId, err := h.svc.PaymentInitialize(int(req.Amount), req.Type)
	if err != nil {
		return nil, err
	}

	return &pb.PaymentInitResponse{Id: int32(id), Paymentid: paymentId}, nil
}

func (h *PaymentHandler) PaymentVerify(c context.Context, req *pb.PaymentVerifyRequest) (*pb.PaymentVerifyResponse, error) {

	if err := h.svc.PaymentVerifying(int(req.Id), req.Signature, req.Orderid, req.Paymentid); err != nil {
		return &pb.PaymentVerifyResponse{Status: false}, err
	}

	return &pb.PaymentVerifyResponse{Status: true}, nil
}
