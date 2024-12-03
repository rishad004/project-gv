package handler

import (
	"context"
	"encoding/json"
	"net/http"

	streamer_service "github.com/rishad004/Gv_protofiles/streamer"
	"github.com/rishad004/project-gv/apiGateway/inertnal/domain"
	"github.com/rishad004/project-gv/apiGateway/utils"
)

func (h *ApiHanlder) Registration(w http.ResponseWriter, r *http.Request) {
	var streamer domain.Streamer

	if err := json.NewDecoder(r.Body).Decode(&streamer); err != nil {
		utils.SendJSONResponse(w, "Invalid payload!", http.StatusBadRequest, r)
		return
	}

	res, err := h.StreamerPb.Registration(context.Background(), &streamer_service.RegistrationRequest{
		Name:        streamer.Name,
		Description: streamer.Description,
		Userid:      int32(r.Context().Value("Id").(uint)),
	})

	if err != nil {
		utils.SendJSONResponse(w, err.Error(), http.StatusBadRequest, r)
		return
	}

	utils.SendJSONResponse(w, res, http.StatusOK, r)
}

func (h *ApiHanlder) ChannelView(w http.ResponseWriter, r *http.Request) {

	res, err := h.StreamerPb.ChannelView(context.Background(), &streamer_service.Verification{
		Id:     int32(r.Context().Value("Id").(uint)),
		Userid: int32(r.Context().Value("Id").(uint)),
	})

	if err != nil {
		utils.SendJSONResponse(w, err.Error(), http.StatusBadRequest, r)
		return
	}

	resp := map[string]any{
		"Channel":     res.Name,
		"Description": res.Description,
		"RTMP-url":    "rtmp://34.72.47.88:1935/live",
		"StreamKey":   res.Streamkey,
	}

	utils.SendJSONResponse(w, resp, http.StatusOK, r)
}

func (h *ApiHanlder) ChannelEdit(w http.ResponseWriter, r *http.Request) {
	var streamer domain.Streamer

	if err := json.NewDecoder(r.Body).Decode(&streamer); err != nil {
		utils.SendJSONResponse(w, "Invalid payload!", http.StatusBadRequest, r)
		return
	}

	res, err := h.StreamerPb.EditChannel(context.Background(), &streamer_service.EditRequest{
		Name:        streamer.Name,
		Description: streamer.Description,
		Userid:      int32(r.Context().Value("Id").(uint)),
	})

	if err != nil {
		utils.SendJSONResponse(w, err.Error(), http.StatusBadRequest, r)
		return
	}

	utils.SendJSONResponse(w, res, http.StatusOK, r)

}
