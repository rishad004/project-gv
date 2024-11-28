package handler

import (
	"net/http"

	"github.com/gorilla/websocket"
	admin_pb "github.com/rishad004/Gv_protofiles/admin"
	payment_pb "github.com/rishad004/Gv_protofiles/payment"
	stream_pb "github.com/rishad004/Gv_protofiles/stream"
	streamer_pb "github.com/rishad004/Gv_protofiles/streamer"
	user_pb "github.com/rishad004/Gv_protofiles/user"
	"github.com/sirupsen/logrus"
)

var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type ApiHanlder struct {
	UserPb     user_pb.UserServiceClient
	StreamerPb streamer_pb.StreamerServiceClient
	PaymentPb  payment_pb.PaymentServiceClient
	StreamPb   stream_pb.StreamServiceClient
	AdminPb    admin_pb.AdminServiceClient
	WsConn     *Socket
	Log        *logrus.Logger
}

func ClientConnect(userSvc user_pb.UserServiceClient, streamerSvc streamer_pb.StreamerServiceClient,
	paymentSvc payment_pb.PaymentServiceClient, streamSvc stream_pb.StreamServiceClient, adminSvc admin_pb.AdminServiceClient, ws *Socket, lg *logrus.Logger) *ApiHanlder {
	return &ApiHanlder{UserPb: userSvc,
		StreamerPb: streamerSvc,
		PaymentPb:  paymentSvc,
		StreamPb:   streamSvc,
		AdminPb:    adminSvc,
		WsConn:     ws,
		Log:        lg,
	}
}
