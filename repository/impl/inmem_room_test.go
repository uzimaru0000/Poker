package impl

import (
	"github.com/uzimaru0000/poker/model"
	"reflect"
	"testing"
)

func Test_inmemRoomRepository_CreateRoom(t *testing.T) {
	type fields struct {
		table map[string]model.Room
	}
	type args struct {
		room *model.Room
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "正常に作成できる",
			fields: fields{table: map[string]model.Room{
				"error": {
					ID: "error",
					Owner: "owner",
					Member: []*model.User{},
				},
			}},
			args: args{
				room: &model.Room{
					ID:     "ROOM_ID",
					Owner:  "OWNER_ID",
					Member: []*model.User{},
				},
			},
			wantErr: false,
		},
		{
			name: "すでにあったらエラーが出る",
			fields: fields{table: map[string]model.Room{
				"error": {
					ID: "error",
					Owner: "owner",
					Member: []*model.User{},
				},
			}},
			args: args{
				room: &model.Room{
					ID:     "error",
					Owner:  "owner",
					Member: []*model.User{},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := inmemRoomRepository{
				table: tt.fields.table,
			}
			if err := i.CreateRoom(tt.args.room); (err != nil) != tt.wantErr {
				t.Errorf("CreateRoom() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_inmemRoomRepository_DeleteRoomByID(t *testing.T) {
	type fields struct {
		table map[string]model.Room
	}
	type args struct {
		s string
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
			i := inmemRoomRepository{
				table: tt.fields.table,
			}
			if err := i.DeleteRoomByID(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("DeleteRoomByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_inmemRoomRepository_GetRoomByID(t *testing.T) {
	type fields struct {
		table map[string]model.Room
	}
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Room
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := inmemRoomRepository{
				table: tt.fields.table,
			}
			got, err := i.GetRoomByID(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRoomByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRoomByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inmemRoomRepository_UpdateRoom(t *testing.T) {
	type fields struct {
		table map[string]model.Room
	}
	type args struct {
		room *model.Room
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
			i := inmemRoomRepository{
				table: tt.fields.table,
			}
			if err := i.UpdateRoom(tt.args.room); (err != nil) != tt.wantErr {
				t.Errorf("UpdateRoom() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
