package mysql

import (
	"reflect"
	"testing"

	"github.com/jameshwc/Million-Singer/model"
)

func Test_mysqlTourRepository_Get(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Tour
		wantErr bool
	}{
		{"success", fields{db}, args{1}, &model.Tour{1, []*model.Collect{
			{1, "collect-1", nil}, {2, "collect-2", nil}, {3, "collect-3", nil}, {4, "collect-4", nil}, {5, "collect-5", nil},
		}}, false},
		{"fail", fields{db}, args{2}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mysqlTourRepository{
				db: tt.fields.db,
			}
			got, err := m.Get(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlTourRepository.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mysqlTourRepository.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlTourRepository_GetTotal(t *testing.T) {
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		{"success", fields{db}, 1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mysqlTourRepository{
				db: tt.fields.db,
			}
			got, err := m.GetTotal()
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlTourRepository.GetTotal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("mysqlTourRepository.GetTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlTourRepository_Add(t *testing.T) {
	type args struct {
		collectsID []int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{"success", fields{db}, args{[]int{1, 3, 5}}, 2, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mysqlTourRepository{
				db: tt.fields.db,
			}
			got, err := m.Add(tt.args.collectsID)
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlTourRepository.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("mysqlTourRepository.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
