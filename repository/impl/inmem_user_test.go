package impl

import (
	"errors"
	"reflect"
	"testing"

	"github.com/uzimaru0000/poker/model"
	"github.com/uzimaru0000/poker/repository"
)

func TestGetUser(t *testing.T) {
	tests := []struct {
		description string
		arg         *model.User
		expected    *model.User
		expectedErr error
	}{
		{
			description: "正常に取得出来る",
			arg:         &model.User{ID: "hoge"},
			expected:    &model.User{ID: "hoge", Name: "Hoge Fuga", Email: "hoge@example.com"},
			expectedErr: nil,
		},
		{
			description: "存在しないIDだとerrorになる",
			arg:         &model.User{ID: "error"},
			expected:    nil,
			expectedErr: &repository.NotFound{ID: "error"},
		},
		{
			description: "Emailアドレスでも取得できる",
			arg:         &model.User{Email: "hoge@example.com"},
			expected:    &model.User{ID: "hoge", Name: "Hoge Fuga", Email: "hoge@example.com"},
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			repo := NewInMemUserRepository()
			repo.(*inmemUserRepository).table = map[string]model.User{
				"hoge": {
					ID:    "hoge",
					Name:  "Hoge Fuga",
					Email: "hoge@example.com",
				},
			}

			actual, err := repo.GetUser(tt.arg)
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf(`
expected:	err = %v
actual:	err = %v
				`, tt.expectedErr, err)
				return
			}

			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf(`
expected:	result = %v
actual:	result = %v
				`, tt.expected, actual)
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	tests := []struct {
		description string
		arg         *model.User
		expected    error
	}{
		{
			description: "正常にユーザーを作成できる",
			arg:         &model.User{ID: "hoge", Name: "Hoge Fuga", Email: "hoge@example.com"},
			expected:    nil,
		},
		{
			description: "すでにIDがあったらエラーになる",
			arg:         &model.User{ID: "error", Name: "Hoge Fuga", Email: "hoge@example.com"},
			expected:    &repository.AlreadyRegistered{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			repo := NewInMemUserRepository()
			repo.(*inmemUserRepository).table = map[string]model.User{
				"error": {
					ID:    "error",
					Name:  "Hoge Fuga",
					Email: "hoge@example.com",
				},
			}

			actual := repo.CreateUser(tt.arg)

			if !errors.Is(actual, tt.expected) {
				t.Errorf(`
expected:	result = %v
actual:		result = %v
				`, tt.expected, actual)
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	tests := []struct {
		description string
		arg         *model.User
		expected    *model.User
		expectedErr error
	}{
		{
			description: "名前が正常にアップデート出来る",
			arg:         &model.User{ID: "test-data", Name: "Hoge Huga", Email: "hoge@example.com"},
			expected:    &model.User{ID: "test-data", Name: "Hoge Huga", Email: "hoge@example.com"},
			expectedErr: nil,
		},
		{
			description: "Emailアドレスがアップデート出来る",
			arg:         &model.User{ID: "test-data", Name: "hoge huga", Email: "hoge@example.co.jp"},
			expected:    &model.User{ID: "test-data", Name: "hoge huga", Email: "hoge@example.co.jp"},
			expectedErr: nil,
		},
		{
			description: "Userがアップデート出来る",
			arg:         &model.User{ID: "test-data", Name: "Hoge Huga", Email: "hoge@example.co.jp"},
			expected:    &model.User{ID: "test-data", Name: "Hoge Huga", Email: "hoge@example.co.jp"},
			expectedErr: nil,
		},
		{
			description: "存在しないユーザーを更新しようとしたらエラー",
			arg:         &model.User{ID: "error", Name: "Hoge Huga", Email: "hoge@example.com"},
			expected:    nil,
			expectedErr: &repository.NotFound{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			repo := NewInMemUserRepository()
			repo.(*inmemUserRepository).table = map[string]model.User{
				"test-data": {
					ID:    "test-data",
					Name:  "hoge huga",
					Email: "hoge@example.com",
				},
			}

			actual, err := repo.UpdateUser(tt.arg)

			if !reflect.DeepEqual(actual, tt.expected) || !errors.Is(err, tt.expectedErr) {
				t.Errorf(`
expected:	result = %v, err = %v
actual:		result = %v, err = %v
				`, tt.expected, tt.expectedErr, actual, err)
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	tests := []struct {
		description string
		arg         *model.User
		expected    error
	}{
		{
			description: "正常に削除できる",
			arg:         &model.User{ID: "test-data"},
			expected:    nil,
		},
		{
			description: "存在しないUserだったらエラー",
			arg:         &model.User{ID: "error"},
			expected:    &repository.NotFound{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			repo := NewInMemUserRepository()
			repo.(*inmemUserRepository).table = map[string]model.User{
				"test-data": {
					ID:    "test-data",
					Name:  "hoge huga",
					Email: "hoge@example.com",
				},
			}

			actual := repo.DeleteUser(tt.arg)

			if !errors.Is(actual, tt.expected) {
				t.Errorf(`
expected:	result = %v
actual:		result = %v
				`, tt.expected, actual)
			}
		})
	}
}
