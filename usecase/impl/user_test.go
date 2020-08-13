package impl

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/uzimaru0000/poker/lib"
	"github.com/uzimaru0000/poker/mock_repository"
	"github.com/uzimaru0000/poker/model"
	"github.com/uzimaru0000/poker/repository"
	"github.com/uzimaru0000/poker/usecase"
)

type MockIDGenerator struct{}

func (*MockIDGenerator) Generate() (string, error) {
	return "MOCK_ID", nil
}

func Test_userUsecase_CreateUser(t *testing.T) {
	type fields struct {
		userRepo      repository.UserRepository
		uuidGenerator lib.UUIDGenerator
	}
	type args struct {
		user *model.User
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		{
			name: "正しく作成される",
			args: args{
				user: &model.User{Name: "uzimaru", Email: "uzimaru@example.com"},
			},
			want:    &model.User{ID: "MOCK_ID", Name: "uzimaru", Email: "uzimaru@example.com"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			repo := mock_repository.NewMockUserRepository(ctrl)
			repo.
				EXPECT().
				CreateUser(tt.args.user).
				Return(nil)

			fields := fields{
				userRepo:      repo,
				uuidGenerator: &MockIDGenerator{},
			}
			uu := &userUsecase{
				userRepo:      fields.userRepo,
				uuidGenerator: fields.uuidGenerator,
			}

			got, err := uu.CreateUser(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUsecase.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUsecase_GetUserByID(t *testing.T) {
	type fields struct {
		userRepo      repository.UserRepository
		uuidGenerator lib.UUIDGenerator
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uu := &userUsecase{
				userRepo:      tt.fields.userRepo,
				uuidGenerator: tt.fields.uuidGenerator,
			}
			got, err := uu.GetUserByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.GetUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUsecase.GetUserByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUsecase_GetUserByEmail(t *testing.T) {
	type fields struct {
		userRepo      repository.UserRepository
		uuidGenerator lib.UUIDGenerator
	}
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uu := &userUsecase{
				userRepo:      tt.fields.userRepo,
				uuidGenerator: tt.fields.uuidGenerator,
			}
			got, err := uu.GetUserByEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.GetUserByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUsecase.GetUserByEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUsecase_UpdateUser(t *testing.T) {
	type fields struct {
		userRepo      repository.UserRepository
		uuidGenerator lib.UUIDGenerator
	}
	type args struct {
		user *model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uu := &userUsecase{
				userRepo:      tt.fields.userRepo,
				uuidGenerator: tt.fields.uuidGenerator,
			}
			got, err := uu.UpdateUser(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUsecase.UpdateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUsecase_DeleteUser(t *testing.T) {
	type fields struct {
		userRepo      repository.UserRepository
		uuidGenerator lib.UUIDGenerator
	}
	type args struct {
		user *model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uu := &userUsecase{
				userRepo:      tt.fields.userRepo,
				uuidGenerator: tt.fields.uuidGenerator,
			}
			if err := uu.DeleteUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewUserUsecase(t *testing.T) {
	type args struct {
		userRepo      repository.UserRepository
		uuidGenerator lib.UUIDGenerator
	}
	tests := []struct {
		name string
		args args
		want usecase.UserUsecase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserUsecase(tt.args.userRepo, tt.args.uuidGenerator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}
