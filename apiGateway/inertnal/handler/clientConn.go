package handler

import (
	payment_pb "github.com/rishad004/Gv_protofiles/payment"
	stream_pb "github.com/rishad004/Gv_protofiles/stream"
	streamer_pb "github.com/rishad004/Gv_protofiles/streamer"
	user_pb "github.com/rishad004/Gv_protofiles/user"
)

type ApiHanlder struct {
	UserPb     user_pb.UserServiceClient
	StreamerPb streamer_pb.StreamerServiceClient
	PaymentPb  payment_pb.PaymentServiceClient
	StreamPb   stream_pb.StreamServiceClient
}

func ClientConnect(userSvc user_pb.UserServiceClient, streamerSvc streamer_pb.StreamerServiceClient, paymentSvc payment_pb.PaymentServiceClient, streamSvc stream_pb.StreamServiceClient) *ApiHanlder {
	return &ApiHanlder{UserPb: userSvc, StreamerPb: streamerSvc, PaymentPb: paymentSvc, StreamPb: streamSvc}
}
