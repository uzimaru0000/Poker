package usecase

type AuthUseCase interface {
	SignUp(string, string) error
	SignIn(string, string) error
}
