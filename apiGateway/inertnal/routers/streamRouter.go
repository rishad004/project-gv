package routers

import (
	"github.com/gorilla/mux"
	"github.com/rishad004/project-gv/apiGateway/inertnal/handler"
)

func StreamRouter(handle *handler.ApiHanlder, r *mux.Router) {
	r.HandleFunc("/live/start", handle.StreamStart).Methods("POST")
	r.HandleFunc("/live/end", handle.StreamEnd).Methods("POST")
	r.HandleFunc("/live/{id}", handle.LiveStream).Methods("GET")
}

func StreamMiddle(handle *handler.ApiHanlder, r *mux.Router) {
	r.HandleFunc("/stream/details", handle.StreamDetailing).Methods("POST")
}
