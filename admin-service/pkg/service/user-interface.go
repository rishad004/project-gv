package service

type AdminRepo interface {
	Login(Email, Password string) (string, error)
	AddAdmin(email, password string) error
}

type AdminService interface {
	Login(Email, Password string) (string, error)
	AddAdmin(email, password string) error
}
