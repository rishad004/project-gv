package di

import (
	"errors"

	"github.com/rishad004/project-gv/stream-service/internal/delivery"
	"github.com/rishad004/project-gv/stream-service/internal/infrastructure/psql"
	"github.com/rishad004/project-gv/stream-service/pkg/repository"
	"github.com/rishad004/project-gv/stream-service/pkg/service"
)

func InitGRPC() error {

	db, err := psql.PsqlConn()
	if err != nil {
		return errors.New("Couldn't connect psql!")
	}

	connStreamer, streamerSvc := InitStreamer()

	defer connStreamer.Close()

	repo := repository.NewStreamRepo(db)
	svc := service.NewStreamService(repo, streamerSvc)
	handler := delivery.NewStreamHandler(svc)

	if err := InitStream(handler); err != nil {
		return err
	}

	return nil
}
