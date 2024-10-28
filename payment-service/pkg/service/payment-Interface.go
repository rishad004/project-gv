package service

type PaymentRepo interface {
	PaymentInitialize(Amount int, Type string) (int, string, error)
	PaymentVerifying(Sig, Ord, Pay string) (error, string)
}

type PaymentService interface {
	PaymentInitialize(Amount int, Type string) (int, string, error)
	PaymentVerifying(id int, Sig, Ord, Pay string) error
}
