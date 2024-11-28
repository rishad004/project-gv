package di

import (
	"log"

	"github.com/rishad004/project-gv/streamer-service/internal/delivery"
	"github.com/rishad004/project-gv/streamer-service/internal/infrastructure/psql"
	"github.com/rishad004/project-gv/streamer-service/pkg/repository"
	"github.com/rishad004/project-gv/streamer-service/pkg/service"
)

func InitGRPC() {

	db, err := psql.PsqlConn()
	if err != nil {
		log.Fatal("Couldn't connect psql!")
	}

	connStream, streamSvc := InitStream()

	defer connStream.Close()

	repo := repository.NewStreamerRepo(db)
	svc := service.NewStreamerService(repo, streamSvc)
	handler := delivery.NewStreamerHandler(svc)

	if err := InitStreamer(handler); err != nil {
		log.Fatal(err)
	}
}
