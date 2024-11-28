package routers

import (
	"github.com/gorilla/mux"
	"github.com/rishad004/project-gv/apiGateway/inertnal/handler"
)

func WsRouting(r *mux.Router, handle *handler.ApiHanlder) {
	r.HandleFunc("/user/chat", handle.UserChat)
}
