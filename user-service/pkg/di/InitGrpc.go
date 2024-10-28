package di

import (
	"log"

	"github.com/rishad004/project-gv/user-service/internal/delivery"
	"github.com/rishad004/project-gv/user-service/internal/infrastructure/psql"
	rediss "github.com/rishad004/project-gv/user-service/internal/infrastructure/redis"
	"github.com/rishad004/project-gv/user-service/pkg/repository"
	"github.com/rishad004/project-gv/user-service/pkg/service"
)

func InitGRPC() {

	db, err := psql.PsqlConn()
	if err != nil {
		log.Fatal("Couldn't connect psql!")
	}

	rdb := rediss.RedisConn()

	connStreamer, streamerSvc := InitStreamer()
	connPayment, paymentSvc := InitPayment()

	defer connStreamer.Close()
	defer connPayment.Close()

	repo := repository.NewUserRepo(db, rdb)
	svc := service.NewUserService(repo, streamerSvc, paymentSvc)
	handler := delivery.NewUserHandler(svc)

	if err := InitUser(handler); err != nil {
		log.Fatal(err)
	}
}
