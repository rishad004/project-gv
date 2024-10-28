package di

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rishad004/project-gv/apiGateway/inertnal/handler"
	"github.com/spf13/viper"
)

func ClientConnect() {
	r := mux.NewRouter()

	connStreamer, streamerSvc := InitStreamer()
	connUser, userSvc := InitUser()
	connPayment, paymentSvc := InitPayment()
	connStream, streamSvc := InitStream()

	handle := handler.ClientConnect(userSvc, streamerSvc, paymentSvc, streamSvc)

	StreamerRouting(r, handle)
	UserRouting(r, handle)
	PaymentRouting(r, handle)
	StreamRouting(r, handle)

	defer connStreamer.Close()
	defer connUser.Close()
	defer connPayment.Close()
	defer connStream.Close()

	log.Println("Server listening on port :8080")
	http.ListenAndServe(viper.GetString("PORT"), r)
}
