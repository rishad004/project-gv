package di

import (
	"errors"

	"github.com/rishad004/project-gv/payment-service/internal/delivery"
	"github.com/rishad004/project-gv/payment-service/internal/infrastructure/psql"
	"github.com/rishad004/project-gv/payment-service/pkg/repository"
	"github.com/rishad004/project-gv/payment-service/pkg/service"
)

func InitGRPC() error {

	db, err := psql.PsqlConn()
	if err != nil {
		return errors.New("Couldn't connect psql!")
	}

	connUser, userSvc := InitUser()
	defer connUser.Close()

	repo := repository.NewPaymentRepo(db)
	svc := service.NewPaymentService(repo, userSvc)
	handler := delivery.NewPaymentHandler(svc)

	if err := InitPayment(handler); err != nil {
		return err
	}

	return nil
}
