package impl

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/uzimaru0000/poker/mock_repository"
	"github.com/uzimaru0000/poker/model"
)

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	userRepo := mock_repository.NewMockUserRepository(ctrl)
	usecase := NewUserUsecase(userRepo, &MockIDGenerator{})

	in := &model.User{Name: "uzimaru", Email: "shuji365630@gmail.com"}
	out := &model.User{ID: "MOCK_ID", Name: "uzimaru", Email: "shuji365630@gmail.com"}

	userRepo.
		EXPECT().
		CreateUser(out).
		Return(nil)

	user, err := usecase.CreateUser(in)

	if err != nil {
		t.Fatalf("ユーザーが正常に作成されませんでした。(message: %s)", err.Error())
	}

	if user.ID != out.ID || user.Name != out.Name || user.Email != out.Email {
		t.Fatalf("与えたUser情報と違う")
	}
}

func TestGetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	userRepo := mock_repository.NewMockUserRepository(ctrl)
	usecase := NewUserUsecase(userRepo, &MockIDGenerator{})

	in := &model.User{ID: "MOCK_ID"}
	out := &model.User{ID: "MOCK_ID", Name: "uzimaru", Email: "shuji365630@gmail.com"}

	userRepo.
		EXPECT().
		GetUser(in).
		Return(out, nil)

	user, err := usecase.GetUser(in)

	if err != nil {
		t.Fatalf("ユーザーが正常に取得出来なかった(message: %s)", err.Error())
	}

	if *user != *out {
		t.Fatalf("正しいUserを取得出来ていない")
	}
}

type MockIDGenerator struct{}

func (*MockIDGenerator) Generate() (string, error) {
	return "MOCK_ID", nil
}
