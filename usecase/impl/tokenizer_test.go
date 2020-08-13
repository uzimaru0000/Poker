package impl

import (
	"reflect"
	"testing"
)

func Test_tokenizer_CreateToken(t *testing.T) {
	type args struct {
		claim interface{}
	}
	tests := []struct {
		name    string
		t       *tokenizer
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.t.CreateToken(tt.args.claim)
			if (err != nil) != tt.wantErr {
				t.Errorf("tokenizer.CreateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("tokenizer.CreateToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tokenizer_VerifyToken(t *testing.T) {
	type args struct {
		tokenString string
	}
	tests := []struct {
		name    string
		t       *tokenizer
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.t.VerifyToken(tt.args.tokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("tokenizer.VerifyToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tokenizer.VerifyToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
