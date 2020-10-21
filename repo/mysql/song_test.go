package mysql

import (
	"database/sql"
	"testing"
)

func Test_mysqlSongRepository_CheckManyExist(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		songsID []int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{"success-1", fields{db}, args{[]int{1}}, 1, false},
		{"success-2", fields{db}, args{[]int{1, 2, 3, 4}}, 4, false},
		{"partly fail", fields{db}, args{[]int{1, 2, 8, 10}}, 2, false},
		{"totally fail", fields{db}, args{[]int{10, 11, 12, 13}}, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mysqlSongRepository{
				db: tt.fields.db,
			}
			got, err := m.CheckManyExist(tt.args.songsID)
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlSongRepository.CheckManyExist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("mysqlSongRepository.CheckManyExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlSongRepository_QueryByVideoID(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		videoID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantId  int64
		wantErr bool
	}{
		{"success", fields{db}, args{"VDvr08sCPOc"}, 3, false},
		{"fail", fields{db}, args{"VDvr08sCPOa"}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mysqlSongRepository{
				db: tt.fields.db,
			}
			gotId, err := m.QueryByVideoID(tt.args.videoID)
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlSongRepository.QueryByVideoID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("mysqlSongRepository.QueryByVideoID() = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}
