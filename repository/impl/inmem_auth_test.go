package impl

import (
	"errors"
	"testing"

	"github.com/uzimaru0000/poker/model"
	"github.com/uzimaru0000/poker/repository"
)

func TestGetHash(t *testing.T) {
	tests := []struct {
		description string
		arg         *model.Auth
		expected    string
		expectedErr error
	}{
		{
			description: "正常に取得出来る",
			arg:         &model.Auth{ID: "hoge"},
			expected:    "HASH",
			expectedErr: nil,
		},
		{
			description: "存在しないIDだとerrorになる",
			arg:         &model.Auth{ID: "error"},
			expected:    "",
			expectedErr: &repository.NotFound{ID: "error"},
		},
	}

	for _, tt := range tests {
		repo := NewInMemAuthRepository()
		repo.(*inmemAuthRepository).table = map[string]model.Auth{
			"hoge": {
				ID:   "hoge",
				Hash: "HASH",
			},
		}

		t.Run(tt.description, func(t *testing.T) {
			actual, err := repo.GetHash(tt.arg)
			if actual != tt.expected || !errors.Is(err, tt.expectedErr) {
				t.Errorf(`
expected:	result = %s, err = %v
actual:	result = %s, err = %v
				`, tt.expected, tt.expectedErr, actual, err)
			}
		})
	}
}

func TestStoreHash(t *testing.T) {
	tests := []struct {
		description string
		arg         *model.Auth
		expected    error
	}{
		{
			description: "正常に保存出来る",
			arg:         &model.Auth{ID: "hoge", Hash: "HASH"},
			expected:    nil,
		},
		{
			description: "すでに存在しているIDだとErrorになる",
			arg:         &model.Auth{ID: "error", Hash: "HASH"},
			expected:    &repository.AlreadyRegistered{ID: "error"},
		},
	}

	for _, tt := range tests {
		repo := NewInMemAuthRepository()
		repo.(*inmemAuthRepository).table = map[string]model.Auth{
			"error": {
				ID:   "error",
				Hash: "HASH",
			},
		}

		t.Run(tt.description, func(t *testing.T) {
			actual := repo.StoreHash(tt.arg)
			if !errors.Is(actual, tt.expected) {
				t.Errorf(`
expected:	result = %s
actual:	result = %s
				`, tt.expected, actual)
			}
		})
	}
}
