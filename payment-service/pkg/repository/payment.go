package repository

import (
	"github.com/rishad004/project-gv/payment-service/internal/domain"
	"github.com/rishad004/project-gv/payment-service/utils"
	"gorm.io/gorm"
)

type paymentRepo struct {
	Db *gorm.DB
}

func NewPaymentRepo(Db *gorm.DB) PaymentRepo {
	return &paymentRepo{Db: Db}
}

func (r *paymentRepo) PaymentInitialize(Amount int, Type string) (int, string, error) {

	razorId, err := utils.Executerazorpay(Type, Amount)
	if err != nil {
		return 0, "", err
	}

	payment := domain.Payment{
		RazorId: razorId,
		Amount:  Amount,
		Type:    Type,
		Status:  false,
	}
	if er := r.Db.Create(&payment).Error; er != nil {
		return 0, "", nil
	}

	return int(payment.ID), razorId, nil
}

func (r *paymentRepo) PaymentVerifying(Sig, Ord, Pay string) (error, string) {
	var payment domain.Payment

	if err := utils.RazorPaymentVerification(Sig, Ord, Pay); err != nil {
		return err, ""
	}

	if err := r.Db.First(&payment, "razor_id=?", Ord).Error; err != nil {
		return err, ""
	}

	payment.OrderId = Pay
	payment.Status = true

	if err := r.Db.Save(&payment).Error; err != nil {
		return err, ""
	}

	return nil, payment.Type
}
