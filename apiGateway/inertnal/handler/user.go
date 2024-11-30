package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	user_pb "github.com/rishad004/Gv_protofiles/user"
	"github.com/rishad004/project-gv/apiGateway/inertnal/domain"
	"github.com/rishad004/project-gv/apiGateway/utils"
)

func (h *ApiHanlder) SignupUser(w http.ResponseWriter, r *http.Request) {

	var user domain.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.SendJSONResponse(w, "Invalid payload!", http.StatusBadRequest, r)
		return
	}

	res, err := h.UserPb.SignUp(context.Background(),
		&user_pb.SignUpRequest{Username: user.Username,
			Password: user.Password,
			Email:    user.Email,
			Phone:    user.Phone,
			Gender:   user.Gender,
		})

	if err != nil {
		utils.SendJSONResponse(w, err.Error(), http.StatusBadRequest, r)
		return
	}

	utils.SendJSONResponse(w, res, http.StatusAccepted, r)
}

func (h *ApiHanlder) EmailVerify(w http.ResponseWriter, r *http.Request) {

	code := r.URL.Query().Get("code")

	res, err := h.UserPb.EmailVerification(context.Background(), &user_pb.VerificationRequest{Key: code})
	if err != nil {
		utils.SendJSONResponse(w, "User not found!", http.StatusNotFound, r)
		return
	}

	utils.SendJSONResponse(w, res, http.StatusOK, r)
}

func (h *ApiHanlder) LoginUser(w http.ResponseWriter, r *http.Request) {

	var user domain.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.SendJSONResponse(w, "Invalid payload!", http.StatusBadRequest, r)
		return
	}

	res, err := h.UserPb.Login(context.Background(),
		&user_pb.LoginRequest{
			Username: user.Username,
			Password: user.Password,
		})
	if err != nil {
		utils.SendJSONResponse(w, err.Error(), http.StatusNotFound, r)
		return
	}

	cookie := http.Cookie{
		Name:     "uqsweerr",
		Value:    res.Token,
		Expires:  time.Now().Add(48 * time.Hour),
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, &cookie)

	utils.SendJSONResponse(w, res, http.StatusOK, r)
}

func (h *ApiHanlder) ProfileU(w http.ResponseWriter, r *http.Request) {

	res, err := h.UserPb.Profile(context.Background(), &user_pb.Verification{
		Id: int32(r.Context().Value("Id").(uint)),
	})
	if err != nil {
		utils.SendJSONResponse(w, err, http.StatusNotFound, r)
		return
	}
	utils.SendJSONResponse(w, res, http.StatusOK, r)
}

func (h *ApiHanlder) ProfileUEdit(w http.ResponseWriter, r *http.Request) {

	var user domain.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.SendJSONResponse(w, "Invalid payload!", http.StatusBadRequest, r)
		return
	}

	res, err := h.UserPb.ProfileEditing(context.Background(),
		&user_pb.EditRequest{
			Id:       int32(r.Context().Value("Id").(uint)),
			Username: user.Username,
			Phone:    user.Phone,
			Gender:   user.Gender,
		})

	if err != nil {
		utils.SendJSONResponse(w, err.Error(), http.StatusBadRequest, r)
		return
	}

	utils.SendJSONResponse(w, res, http.StatusAccepted, r)
}

func (h *ApiHanlder) LogoutU(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "uqsweerr",
		Value:    "",
		Expires:  time.Now().Add(48 * time.Hour),
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
	}
	http.SetCookie(w, &cookie)

	utils.SendJSONResponse(w, "Logged out successfully!", http.StatusOK, r)
}

func (h *ApiHanlder) Following(w http.ResponseWriter, r *http.Request) {
	log.Println("following")

	utils.SetCors(w)

	orderId := r.URL.Query().Get("id")
	streamerId, _ := strconv.Atoi(orderId)

	if _, err := h.UserPb.Following(context.Background(), &user_pb.SubscribeRequest{Id: int32(streamerId),
		Userid: int32(r.Context().Value("Id").(uint))}); err != nil {
		log.Println("kkkk")
		utils.SendJSONResponse(w, err.Error(), http.StatusBadRequest, r)
		return
	}

	utils.SendJSONResponse(w, "Followed successfully!", http.StatusOK, r)
}
