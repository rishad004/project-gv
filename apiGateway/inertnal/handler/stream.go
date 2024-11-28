package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	stream_service "github.com/rishad004/Gv_protofiles/stream"
	streamer_pb "github.com/rishad004/Gv_protofiles/streamer"
	"github.com/rishad004/project-gv/apiGateway/utils"
)

func (h *ApiHanlder) LiveStream(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Channel := vars["id"]

	res, err := h.StreamerPb.FindByStreamKey(context.Background(), &streamer_pb.StreamKeyRequest{Channel: Channel})
	if err != nil {
		utils.SendJSONResponse(w, err.Error(), http.StatusBadRequest, r)
		return
	}

	username, er := utils.UserFromCookie(r, "uqsweerr")

	utils.RenderTemplate(w, "live.html", map[string]any{
		"Title":       res.Title,
		"Description": res.Description,
		"StreamerId":  res.Id,
		"StreamKey":   res.Streamkey,
		"Channel":     Channel,
		"Username":    username,
		"check":       er == nil,
	})

}

func (h *ApiHanlder) StreamStart(w http.ResponseWriter, r *http.Request) {
	log.Println("stream is starting...........")

	streamKey := r.FormValue("name")

	if _, err := h.StreamerPb.StreamStart(context.Background(), &streamer_pb.StreamKeyResponse{Streamkey: streamKey}); err != nil {
		utils.SendJSONResponse(w, err.Error(), http.StatusUnauthorized, r)
		return
	}

	utils.SendJSONResponse(w, nil, http.StatusOK, r)
}

func (h *ApiHanlder) StreamEnd(w http.ResponseWriter, r *http.Request) {
	log.Println("stream is ending...........")

	streamKey := r.FormValue("name")

	if _, err := h.StreamerPb.StreamEnd(context.Background(), &streamer_pb.StreamKeyResponse{Streamkey: streamKey}); err != nil {
		utils.SendJSONResponse(w, err.Error(), http.StatusUnauthorized, r)
		return
	}

	utils.SendJSONResponse(w, nil, http.StatusOK, r)
}

func (h *ApiHanlder) StreamDetailing(w http.ResponseWriter, r *http.Request) {

	req := make(map[string]any)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendJSONResponse(w, "Invalid payload!", http.StatusBadRequest, r)
		return
	}

	if _, err := h.StreamPb.StreamDetailing(context.Background(), &stream_service.Data{Id: int32(r.Context().Value("Id").(uint)),
		Title:       req["title"].(string),
		Description: req["description"].(string)}); err != nil {
		utils.SendJSONResponse(w, err.Error(), http.StatusBadRequest, r)
		return
	}

	utils.SendJSONResponse(w, "title and description updated successfully!", http.StatusOK, r)
}
