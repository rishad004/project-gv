package routers

import (
	"github.com/gorilla/mux"
	"github.com/rishad004/project-gv/apiGateway/inertnal/handler"
)

func UserMiddle(user *handler.ApiHanlder, r *mux.Router) {
	r.HandleFunc("/profile", user.ProfileU).Methods("GET")
	r.HandleFunc("/profile/edit", user.ProfileUEdit).Methods("PUT")
	r.HandleFunc("/logout", user.LogoutU).Methods("DELETE")
	r.HandleFunc("/subscribe", user.Subscribing).Methods("POST")
	r.HandleFunc("/streamer/list", user.SubscriptionList).Methods("GET")
	r.HandleFunc("/streamer/follow", user.Following).Methods("POST")
}

func UserRouter(user *handler.ApiHanlder, r *mux.Router) {
	r.HandleFunc("/signup", user.SignupUser).Methods("POST")
	r.HandleFunc("/verify", user.EmailVerify).Methods("PATCH")
	r.HandleFunc("/login", user.LoginUser).Methods("POST")
}
