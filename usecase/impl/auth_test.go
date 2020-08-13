package impl

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/uzimaru0000/poker/mock_repository"
	"github.com/uzimaru0000/poker/model"
)

func Test_authUsecase_SignUp(t *testing.T) {
	type args struct {
		id       string
		password string
	}

	tests := []struct {
		name    string
		au      *authUsecase
		args    args
		wantErr bool
	}{
		{
			name: "SignUpが出来る",
			args: args{
				id:       "HogeHoge",
				password: "password",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			repo := mock_repository.NewMockAuthRepository(ctrl)

			repo.EXPECT().StoreHash(&model.Auth{ID: tt.args.id, Hash: tt.args.password}).Return(nil)

			tt.au = NewAuthUsecase(repo, &MockHashGen{}).(*authUsecase)

			if err := tt.au.SignUp(tt.args.id, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("authUsecase.SignUp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_authUsecase_SignIn(t *testing.T) {
	type args struct {
		id       string
		password string
	}
	tests := []struct {
		name    string
		au      *authUsecase
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.au.SignIn(tt.args.id, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("authUsecase.SignIn() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

type MockHashGen struct{}

func (*MockHashGen) Create(pass string) (string, error) {
	return pass, nil
}

func (*MockHashGen) Verify(hash string, pass string) error {
	if hash != pass {
		return fmt.Errorf("Error")
	}

	return nil
}
