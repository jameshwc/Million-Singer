package model

import (
	"reflect"
	"testing"
)

func TestGetCollectSuccess(t *testing.T) {

	mock := setupTestDatabase(t)
	if mock == nil {
		return
	}

	mock.ExpectQuery("SELECT (.+) FROM `(.+)` WHERE `(.+)`.`(.+)_id`")
	// WithArgs(1).WillReturnRows()
}

func TestGetCollect(t *testing.T) {
	type args struct {
		collectID int
	}
	tests := []struct {
		name    string
		args    args
		want    *Collect
		wantErr bool
	}{
		{name: "success GetCollect", args: args{collectID: 1}, want: &Collect{Title: "hello world"}, wantErr: false},
	}

	setupTestDatabase(t)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCollect(tt.args.collectID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCollect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCollect() = %v, want %v", got, tt.want)
			}
		})
	}
}
