package delivery

import (
	"context"

	pb "github.com/rishad004/Gv_protofiles/streamer"
)

func (h *StreamerHandler) SubscriptionSet(c context.Context, req *pb.SubscriptionRequest) (*pb.SubscriptionResponse, error) {
	if err := h.svc.SubscriptionSetting(int(req.Userid), int(req.Amount)); err != nil {
		return nil, err
	}

	return &pb.SubscriptionResponse{Message: "Subscription amount setting done!"}, nil
}

func (h *StreamerHandler) SubscriptionCheck(c context.Context, req *pb.Verification) (*pb.CheckingResponse, error) {
	amount, sid, err := h.svc.SubscriptionCheck(int(req.Userid), int(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb.CheckingResponse{Amount: int32(amount), Sid: int32(sid)}, nil
}

func (h *StreamerHandler) SubscriptionList(c context.Context, req *pb.Empty) (*pb.ListResponse, error) {
	res, err := h.svc.SubscriptionList()
	if err != nil {
		return nil, err
	}

	list := &pb.ListResponse{}

	for _, v := range res {
		data := &pb.SubscriptionList{
			Streamerid:     int32(v["streamer_Id"].(uint)),
			Channel:        v["channel_Name"].(string),
			Subscriptionid: int32(v["subscription_Id"].(int)),
			Amount:         v["subscription_Amount"].(string),
		}

		list.List = append(list.List, data)
	}

	return list, nil
}
