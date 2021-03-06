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
		name   string
		fields fields
		args   args
		wantId int64
		err    error
	}{
		{"success", fields{db}, args{"VDvr08sCPOc"}, 3, nil},
		{"fail", fields{db}, args{"VDvr08sCPOa"}, 0, sql.ErrNoRows},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mysqlSongRepository{
				db: tt.fields.db,
			}
			gotId, err := m.QueryByVideoID(tt.args.videoID)
			if err != tt.err {
				t.Errorf("mysqlSongRepository.QueryByVideoID() error = %v, wantErr %v", err, tt.err)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("mysqlSongRepository.QueryByVideoID() = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}

func Test_mysqlSongRepository_Gets(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"success", fields{db}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mysqlSongRepository{
				db: tt.fields.db,
			}
			gotSongs, err := m.Gets()
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlSongRepository.Gets() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for _, song := range gotSongs {
				if len(song.Name) == 0 || len(song.Singer) == 0 || song.ID == 0 {
					t.Errorf("mysqlSongRepository.Gets() return nil song")
					return
				}
			}
		})
	}
}
