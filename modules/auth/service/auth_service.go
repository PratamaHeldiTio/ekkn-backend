package service

type AuthService interface {
	GenerateTokenJwt(nim string) (string, error)
}
