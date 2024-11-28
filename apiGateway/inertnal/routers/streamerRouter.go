package routers

import (
	"github.com/gorilla/mux"
	"github.com/rishad004/project-gv/apiGateway/inertnal/handler"
)

func StreamerRouter(handle *handler.ApiHanlder, r *mux.Router) {
	r.HandleFunc("/register", handle.Registration).Methods("POST")
	r.HandleFunc("/channel", handle.ChannelView).Methods("GET")
	r.HandleFunc("/channel/edit", handle.ChannelEdit).Methods("PUT")
	r.HandleFunc("/channel/subscription", handle.SubscriptionSet).Methods("POST")
}
