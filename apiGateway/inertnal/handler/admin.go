package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	admin_server "github.com/rishad004/Gv_protofiles/admin"
	stream_service "github.com/rishad004/Gv_protofiles/stream"
	user_service "github.com/rishad004/Gv_protofiles/user"
	"github.com/rishad004/project-gv/apiGateway/inertnal/domain"
	"github.com/rishad004/project-gv/apiGateway/utils"
	"github.com/sirupsen/logrus"
)

func (h *ApiHanlder) LoginAdmin(w http.ResponseWriter, r *http.Request) {
	h.Log.Info("AdminLogin: Starting login process.....")

	var admin domain.User

	if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
		h.Log.WithFields(logrus.Fields{
			"function": "LoginAdmin",
			"error":    err,
		}).Error("LoginAdmin: error on decoding json value")

		utils.SendJSONResponse(w, "Invalid payload!", http.StatusBadRequest, r)
		return
	}

	res, err := h.AdminPb.Login(context.Background(), &admin_server.LoginRequest{
		Email:    admin.Email,
		Password: admin.Password,
	})

	if err != nil {
		h.Log.WithFields(logrus.Fields{
			"function": "LoginAdmin",
			"error":    err.Error(),
			"email":    admin.Email,
		}).Error("LoginAdmin: gRPC error on login")

		utils.SendJSONResponse(w, err.Error(), http.StatusBadRequest, r)
		return
	}

	cookie := http.Cookie{
		Name:     "aqdwmeirn",
		Value:    res.Token,
		Expires:  time.Now().Add(48 * time.Hour),
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
	}
	http.SetCookie(w, &cookie)

	h.Log.WithFields(logrus.Fields{
		"function": "LoginAdmin",
		"email":    admin.Email,
	}).Info("LoginAdmin: admin login successfull")

	utils.SendJSONResponse(w, utils.H{
		"token":    res.Token,
		"message ": "Admin loggined successfully",
	}, http.StatusOK, r)
}

func (h *ApiHanlder) AddAdmin(w http.ResponseWriter, r *http.Request) {
	h.Log.Info("AddAdmin: Starting adding admin process.....")

	var admin domain.User

	super, er := utils.UserFromCookie(r, "aqdwmeirn")
	if er != nil {
		utils.SendJSONResponse(w, "Invalid cookie!", http.StatusBadRequest, r)
		return
	}

	if super != "superadmin@gmail.com" {
		h.Log.WithFields(logrus.Fields{
			"function": "AddAdmin",
			"error":    errors.New("not a super admin only super admin can add admin"),
		}).Error("AddAdmin: not a super admin")

		utils.SendJSONResponse(w, "Unauthorized!", http.StatusUnauthorized, r)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
		h.Log.WithFields(logrus.Fields{
			"function": "AddAdmin",
			"error":    err,
		}).Error("AddAdmin: error on decoding json value")

		utils.SendJSONResponse(w, "Invalid payload!", http.StatusBadRequest, r)
		return
	}

	if _, err := h.AdminPb.AddAdmin(context.Background(), &admin_server.LoginRequest{
		Email:    admin.Email,
		Password: admin.Password,
	}); err != nil {
		h.Log.WithFields(logrus.Fields{
			"function": "AddAdmin",
			"error":    err.Error(),
			"email":    r.Context().Value("Email").(string),
		}).Error("AddAdmin: gRPC error on adding admin")

		utils.SendJSONResponse(w, err.Error(), http.StatusBadRequest, r)
		return
	}

	h.Log.WithFields(logrus.Fields{
		"function": "AddAdmin",
		"email":    r.Context().Value("Email").(string),
	}).Info("AddAdmin: Added new admin successfully")

	utils.SendJSONResponse(w, utils.H{
		"message ": "added admin successfully",
	}, http.StatusOK, r)
}

func (h *ApiHanlder) AdminDashboard(w http.ResponseWriter, r *http.Request) {
	h.Log.Info("AdminDashboard: Starting admin dashboard.....")

	var s int

	res, err := h.StreamPb.StreamerCount(context.Background(), &stream_service.Empty{})
	if err != nil {
		h.Log.WithFields(logrus.Fields{
			"function": "AdminDashboard",
			"error":    err.Error(),
			"email":    r.Context().Value("Email").(string),
		}).Error("AdminDashboard: gRPC error on admin dashboard finding stream count")

		utils.SendJSONResponse(w, err.Error(), http.StatusBadRequest, r)
		return
	}

	re, er := h.UserPb.SuperChatTotal(context.Background(), &user_service.Empty{})
	if er != nil {
		h.Log.WithFields(logrus.Fields{
			"function": "AdminDashboard",
			"error":    er.Error(),
			"email":    r.Context().Value("Email").(string),
		}).Error("AdminDashboard: gRPC error on admin dashboard finding superchat total")

		utils.SendJSONResponse(w, er.Error(), http.StatusBadRequest, r)
		return
	}

	for _, v := range h.WsConn.Conns {
		fmt.Println(v)
		s += len(v)
	}

	if res.StreamerId != 0 {
		s = s / int(res.StreamerId)
	}

	h.Log.WithFields(logrus.Fields{
		"function": "AdminDashboard",
		"email":    r.Context().Value("Email").(string),
	}).Info("AdminDashboard: admin dashboard proccess completed successfully")

	utils.SendJSONResponse(w, utils.H{
		"streams":         res.StreamerId,
		"avg_viewers":     s,
		"total_superchat": -re.Id,
	}, http.StatusOK, r)
}
