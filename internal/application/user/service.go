package user

type Service interface {
	Register(req RegisterRequest) error
}
