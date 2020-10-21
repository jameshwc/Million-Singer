package mysql

import (
	"reflect"
	"testing"

	"github.com/jameshwc/Million-Singer/model"
)

func Test_mysqlCollectRepository_Get(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Collect
		wantErr bool
	}{
		{"success", fields{db}, args{1}, &model.Collect{1, "collect-1", []*model.Song{
			{1, nil, "avLxcVkPgug", "", "", "en", "Beautiful", "Eminem", "rap,hip-hop", ""},
			{2, nil, "JxzKNHfNRdI", "", "", "en", "No Sleep", "Martin Garrix feat. Bonn", "edm", ""},
			{3, nil, "VDvr08sCPOc", "", "", "en", "Remember The Name", "Fort Minor", "rap,hip-hop", ""},
		}}, false},
		{"fail", fields{db}, args{6}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mysqlCollectRepository{
				db: tt.fields.db,
			}
			got, err := m.Get(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlCollectRepository.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mysqlCollectRepository.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_mysqlCollectRepository_Add(t *testing.T) {
	type args struct {
		title   string
		songsID []int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{"success", fields{db}, args{"collect-6", []int{2, 3, 4}}, 6, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mysqlCollectRepository{
				db: tt.fields.db,
			}
			got, err := m.Add(tt.args.title, tt.args.songsID)
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlCollectRepository.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("mysqlCollectRepository.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlCollectRepository_CheckManyExist(t *testing.T) {
	type args struct {
		collectsID []int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{"success-1", fields{db}, args{[]int{1, 2, 3}}, 3, false},
		{"success-2", fields{db}, args{[]int{1, 2, 3, 4, 5}}, 5, false},
		{"partly fail-1", fields{db}, args{[]int{1, 3, 5, 8}}, 3, false},
		{"partly fail-2", fields{db}, args{[]int{1, 8}}, 1, false},
		{"totally fail", fields{db}, args{[]int{8, 10, 12}}, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mysqlCollectRepository{
				db: tt.fields.db,
			}
			got, err := m.CheckManyExist(tt.args.collectsID)
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlCollectRepository.CheckManyExist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("mysqlCollectRepository.CheckManyExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
