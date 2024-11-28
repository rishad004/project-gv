package handler

import (
	"context"
	"encoding/json"
	"net/http"

	userpb "github.com/rishad004/Gv_protofiles/user"
	"github.com/rishad004/project-gv/apiGateway/inertnal/domain"
	"github.com/rishad004/project-gv/apiGateway/utils"
)

func (h *ApiHanlder) WalletAdd(w http.ResponseWriter, r *http.Request) {
	var wallet domain.Message

	if err := json.NewDecoder(r.Body).Decode(&wallet); err != nil {
		utils.SendJSONResponse(w, "Invalid payload!", http.StatusBadRequest, r)
		return
	}

	res, err := h.UserPb.WalletAdd(context.Background(), &userpb.Amount{Userid: int32(r.Context().Value("Id").(uint)), Amount: int32(wallet.Amount)})

	if err != nil {
		utils.SendJSONResponse(w, err.Error(), http.StatusNotFound, r)
		return
	}

	utils.SendJSONResponse(w, res, http.StatusOK, r)

}

func (h *ApiHanlder) Superchat(amount int, userid int) error {

	if _, err := h.UserPb.SuperChat(context.Background(), &userpb.Amount{Userid: int32(userid), Amount: int32(amount)}); err != nil {
		return err
	}

	return nil
}

func (h *ApiHanlder) WalletShow(w http.ResponseWriter, r *http.Request) {

	res, err := h.UserPb.WalletShow(context.Background(), &userpb.Verification{Id: int32(r.Context().Value("Id").(uint))})
	if err != nil {
		utils.SendJSONResponse(w, err.Error(), http.StatusNotFound, r)
		return
	}

	utils.SendJSONResponse(w, map[string]int32{
		"amount": res.Id,
	}, http.StatusOK, r)
}
