package routers

import (
	"github.com/gorilla/mux"
	"github.com/rishad004/project-gv/apiGateway/inertnal/handler"
)

func PaymentRouter(handle *handler.ApiHanlder, r *mux.Router) {
	r.HandleFunc("/page", handle.PaymentRendering).Methods("GET")
	r.HandleFunc("/verify", handle.PaymentVerifying).Methods("POST")
}
