package usecase

type AuthUsecase interface {
	SignUp(string, string) error
	SignIn(string, string) error
}
