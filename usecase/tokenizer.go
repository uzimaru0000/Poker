package usecase

type Tokenizer interface {
	CreateToken(interface{}) (string, error)
	VerifyToken(string) (interface{}, error)
}
