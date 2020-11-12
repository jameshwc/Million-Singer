package game

import (
	"reflect"
	"testing"

	"github.com/jameshwc/Million-Singer/model"
	C "github.com/jameshwc/Million-Singer/pkg/constant"
	"github.com/jameshwc/Million-Singer/pkg/log"
	"github.com/jameshwc/Million-Singer/pkg/subtitle"
	"github.com/sirupsen/logrus"
)

func TestService_AddSong(t *testing.T) {
	type args struct {
		s *Song
	}
	tests := []struct {
		name     string
		args     args
		subtitle subtitle.Subtitler
		weber    subtitle.Weber
		want     int
		err      error
	}{
		{"fail url incorrect", args{
			&Song{[]byte{}, "youtube", "", "beautiful", "eminem", []int{1, 2, 3}, "hip-hop", "en"},
		}, NewMockSubtitler(), NewMockWeber(), 0, C.ErrSongFormatIncorrect},
		{"fail url incorrect-2", args{
			&Song{[]byte{}, "youtube", "https://youtube.com/", "beautiful", "eminem", []int{1, 2, 3}, "hip-hop", "en"},
		}, NewMockSubtitler(), NewMockWeber(), 0, C.ErrSongAddURLIncorrect},
		{"fail url incorrect-3", args{
			&Song{[]byte{}, "youtube", "https://youtube.com/", "beautiful", "eminem", []int{1, 2, 3}, "hip-hop", "en"},
		}, NewMockSubtitler(), NewMockWeber(), 0, C.ErrSongAddURLIncorrect},
	}

	log.Logger = logrus.New()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &Service{}
			subtitle.Youtube = tt.weber
			subtitle.SRT = tt.subtitle
			subtitle.LRC = tt.subtitle
			got, err := srv.AddSong(tt.args.s)
			if err != tt.err {
				t.Errorf("Service.AddSong() error = %v, want err %v", err, tt.err)
				return
			}
			if got != tt.want {
				t.Errorf("Service.AddSong() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetSongs(t *testing.T) {
	tests := []struct {
		name    string
		srv     *Service
		want    []*model.Song
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &Service{}
			got, err := srv.GetSongs()
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetSongs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetSongs() = %v, want %v", got, tt.want)
			}
		})
	}
}
