package delivery

import (
	"context"

	pb "github.com/rishad004/Gv_protofiles/user"
)

func (h *UserHandler) WalletAdd(c context.Context, req *pb.Amount) (*pb.SubscribeResponse, error) {

	paymentId, err := h.svc.WalletAdd(req.Amount, req.Userid)

	if err != nil {
		return nil, err
	}

	return &pb.SubscribeResponse{Paymentid: paymentId, Message: "Wallet adding initialized"}, nil
}

func (h *UserHandler) WalletAdded(c context.Context, req *pb.SubscribedRequest) (*pb.Empty, error) {

	if err := h.svc.WalletAdded(req.Paymentid, int(req.Id)); err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (h *UserHandler) SuperChat(c context.Context, req *pb.Amount) (*pb.Empty, error) {
	if err := h.svc.SuperChat(int(req.Amount), int(req.Userid)); err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (h *UserHandler) WalletShow(c context.Context, req *pb.Verification) (*pb.Verification, error) {
	amount, err := h.svc.WalletShow(int(req.Id))
	if err != nil {
		return nil, err
	}

	return &pb.Verification{Id: int32(amount)}, nil
}

func (h *UserHandler) SuperChatTotal(c context.Context, req *pb.Empty) (*pb.Verification, error) {
	sum, err := h.svc.SuperChatTotal()

	if err != nil {
		return nil, err
	}

	return &pb.Verification{Id: int32(sum)}, nil
}
