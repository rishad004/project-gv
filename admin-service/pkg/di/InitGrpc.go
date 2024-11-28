package di

import (
	"log"

	"github.com/rishad004/project-gv/admin-service/internal/delivery"
	"github.com/rishad004/project-gv/admin-service/internal/infrastructure/psql"
	rediss "github.com/rishad004/project-gv/admin-service/internal/infrastructure/redis"
	"github.com/rishad004/project-gv/admin-service/pkg/repository"
	"github.com/rishad004/project-gv/admin-service/pkg/service"
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

	repo := repository.NewAdminRepo(db, rdb)
	svc := service.NewAdminService(repo, streamerSvc, paymentSvc)
	handler := delivery.NewAdminHandler(svc)

	if err := InitAdmin(handler); err != nil {
		log.Fatal(err)
	}
}
