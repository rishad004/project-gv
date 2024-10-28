package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	streamer_pb "github.com/rishad004/Gv_protofiles/streamer"
	user_pb "github.com/rishad004/Gv_protofiles/user"
	"github.com/rishad004/project-gv/apiGateway/inertnal/domain"
	"github.com/rishad004/project-gv/apiGateway/utils"
)

func (h *ApiHanlder) Subscribing(w http.ResponseWriter, r *http.Request) {

	utils.SetCors(w)

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		utils.SendJSONResponse(w, err.Error(), http.StatusBadRequest, r)
		return
	}

	res, er := h.UserPb.Subscribing(context.Background(), &user_pb.SubscribeRequest{Id: int32(id), Userid: int32(r.Context().Value("Id").(uint))})
	if er != nil {
		utils.SendJSONResponse(w, er.Error(), http.StatusBadRequest, r)
		return
	}
	utils.SendJSONResponse(w, res, http.StatusOK, r)
}

func (h *ApiHanlder) SubscriptionSet(w http.ResponseWriter, r *http.Request) {
	var subscription domain.Subscription

	if err := json.NewDecoder(r.Body).Decode(&subscription); err != nil {
		utils.SendJSONResponse(w, "Invalid payload!", http.StatusBadRequest, r)
		return
	}

	res, err := h.StreamerPb.SubscriptionSet(context.Background(), &streamer_pb.SubscriptionRequest{
		Userid: int32(r.Context().Value("Id").(uint)),
		Amount: int32(subscription.Amount),
	})
	if err != nil {
		utils.SendJSONResponse(w, err.Error(), http.StatusBadRequest, r)
		return
	}

	utils.SendJSONResponse(w, res, http.StatusOK, r)
}

func (h *ApiHanlder) SubscriptionList(w http.ResponseWriter, r *http.Request) {
	res, err := h.StreamerPb.SubscriptionList(context.Background(), &streamer_pb.Empty{})
	if err != nil {
		utils.SendJSONResponse(w, err.Error(), http.StatusBadRequest, r)
		return
	}

	utils.SendJSONResponse(w, res, http.StatusOK, r)
}
