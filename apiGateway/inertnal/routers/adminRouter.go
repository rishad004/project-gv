package routers

import (
	"github.com/gorilla/mux"
	"github.com/rishad004/project-gv/apiGateway/inertnal/handler"
)

func AdminMiddle(admin *handler.ApiHanlder, r *mux.Router) {
	r.HandleFunc("/add/admin", admin.AddAdmin).Methods("POST")
	r.HandleFunc("/dashboard", admin.AdminDashboard).Methods("GET")
}

func AdminRouter(admin *handler.ApiHanlder, r *mux.Router) {
	r.HandleFunc("/", admin.LoginAdmin).Methods("POST")
}
