package mysql

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/jameshwc/Million-Singer/model"
)

func Test_mysqlUserRepository_IsNameDuplicate(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"duplicate", fields{db}, args{"alice"}, true},
		{"not duplicate", fields{db}, args{"alice-2"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mysqlUserRepository{
				db: tt.fields.db,
			}
			if got := m.IsNameDuplicate(tt.args.name); got != tt.want {
				t.Errorf("mysqlUserRepository.IsNameDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlUserRepository_IsEmailDuplicate(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		email string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"duplicate", fields{db}, args{"alice@example.com"}, true},
		{"not duplicate", fields{db}, args{"alice-2@example.com"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mysqlUserRepository{
				db: tt.fields.db,
			}
			if got := m.IsEmailDuplicate(tt.args.email); got != tt.want {
				t.Errorf("mysqlUserRepository.IsEmailDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_encrypt(t *testing.T) {
	type args struct {
		pw string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"valid", args{"password"}, "5baa61e4c9b93f3f0682250b6cf8331b7ee68fd8"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encrypt(tt.args.pw); got != tt.want {
				t.Errorf("encrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlUserRepository_Add(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		name     string
		email    string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{"test1", fields{db}, args{"bob", "bob@example.com", "password"}, 2, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mysqlUserRepository{
				db: tt.fields.db,
			}
			got, err := m.Add(tt.args.name, tt.args.email, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlUserRepository.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("mysqlUserRepository.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlUserRepository_Auth(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *model.User
		err    error
	}{
		{"success", fields{db}, args{"alice", "password"}, &model.User{1, "alice", "", false, nil}, nil},
		{"fail", fields{db}, args{"alice", "password123"}, nil, sql.ErrNoRows},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mysqlUserRepository{
				db: tt.fields.db,
			}
			got, err := m.Auth(tt.args.username, tt.args.password)
			if err != tt.err {
				t.Errorf("mysqlUserRepository.Auth() error = %v, wantErr %v", err, tt.err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mysqlUserRepository.Auth() = %v, want %v", got, tt.want)
			}
		})
	}
}
