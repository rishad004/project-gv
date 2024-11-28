package di

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/rishad004/project-gv/apiGateway/inertnal/handler"
	"github.com/rishad004/project-gv/apiGateway/inertnal/routers"
	"github.com/rishad004/project-gv/apiGateway/pkg/middleware"
)

func NewConn() *handler.Socket {
	return &handler.Socket{
		Conns: make(map[string][]*websocket.Conn),
	}
}

func WsRouting(r *mux.Router, handle *handler.ApiHanlder) {
	wsRoute := r.PathPrefix("/ws").Subrouter()
	wsRoute.Use(middleware.MiddlewareU)
	routers.WsRouting(wsRoute, handle)
}
