package repository

type PaymentRepo interface {
	PaymentInitialize(Amount int, Type string) (int, string, error)
	PaymentVerifying(Sig, Ord, Pay string) (error, string)
}
