package routers

import (
	"github.com/gorilla/mux"
	"github.com/rishad004/project-gv/apiGateway/inertnal/handler"
)

func StreamRouter(handle *handler.ApiHanlder, r *mux.Router) {
	r.HandleFunc("/live/start/{id}", handle.StreamStart).Methods("POST")
	r.HandleFunc("/live/end/{id}", handle.StreamEnd).Methods("POST")
	r.HandleFunc("/live/{id}", handle.LiveStream).Methods("GET")
}

func StreamMiddle(handle *handler.ApiHanlder, r *mux.Router) {
	r.HandleFunc("/stream/details", handle.StreamDetailing).Methods("POST")
}
