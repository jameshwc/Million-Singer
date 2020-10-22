package game

import (
	"reflect"
	"testing"

	"github.com/jameshwc/Million-Singer/model"
	C "github.com/jameshwc/Million-Singer/pkg/constant"
	"github.com/jameshwc/Million-Singer/pkg/log"
	"github.com/jameshwc/Million-Singer/repo"
	"github.com/sirupsen/logrus"
)

func TestService_GetCollect(t *testing.T) {
	type args struct {
		param string
	}
	tests := []struct {
		name    string
		args    args
		collect repo.CollectRepo
		cache   repo.CacheRepo
		want    *model.Collect
		err     error
	}{
		{"success", args{"1"}, newRepoMockCollectBase([]int{1}), newRepoMockCacheServerError(), &model.Collect{1, "collect-1", nil}, nil},
		{"param fail", args{"1u"}, newRepoMockCollectBase([]int{1}), newRepoMockCacheServerError(), nil, C.ErrCollectIDNotNumber},
		{"param fail-2", args{"u"}, newRepoMockCollectBase([]int{1}), newRepoMockCacheServerError(), nil, C.ErrCollectIDNotNumber},
		{"collect id not found", args{"2"}, newRepoMockCollectBase([]int{1}), newRepoMockCacheServerError(), nil, C.ErrCollectNotFound},
		{"collect id empty", args{""}, newRepoMockCollectBase([]int{1}), newRepoMockCacheServerError(), nil, C.ErrCollectIDNotNumber},
		{"collect database error", args{"1"}, newRepoMockCollectServerError(), newRepoMockCacheServerError(), nil, C.ErrDatabase},
		{"cache used but parse error", args{"1"}, newRepoMockCollectBase([]int{1}), newRepoMockCacheBase(), &model.Collect{1, "collect-1", nil}, nil},
		// TODO: cache success test case
	}

	log.Logger = logrus.New()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &Service{}
			repo.Collect = tt.collect
			repo.Cache = tt.cache
			got, err := srv.GetCollect(tt.args.param)
			if err != tt.err {
				t.Errorf("Service.GetCollect() error = %v, wantErr %v", err, tt.err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetCollect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_AddCollect(t *testing.T) {
	type args struct {
		songs []int
		title string
	}
	tests := []struct {
		name    string
		args    args
		collect repo.CollectRepo
		song    repo.SongRepo
		want    int
		err     error
	}{
		{"success", args{[]int{1, 2, 3}, "collect-2"}, newRepoMockCollectBase([]int{1}), newRepoMockSongBase([]int{1, 2, 3}), 2, nil},
		{"song nil fail", args{[]int{}, "collect-2"}, newRepoMockCollectBase([]int{1}), newRepoMockSongBase([]int{1, 2, 3}), 0, C.ErrCollectAddFormatIncorrect},
		{"song record not found", args{[]int{15, 21, 22}, "collect-2"}, newRepoMockCollectBase([]int{1}), newRepoMockSongBase([]int{1, 2, 3}), 0, C.ErrCollectAddSongsRecordNotFound},
		{"partly song record not found", args{[]int{1, 20}, "collect-2"}, newRepoMockCollectBase([]int{1}), newRepoMockSongBase([]int{1, 2, 3}), 0, C.ErrCollectAddSongsRecordNotFound},
		{"song record duplicate", args{[]int{1, 1, 2}, "collect-2"}, newRepoMockCollectBase([]int{1}), newRepoMockSongBase([]int{1, 2, 3}), 0, C.ErrCollectAddSongsDuplicate},
		{"collect database error", args{[]int{1, 2, 3}, "collect-2"}, newRepoMockCollectServerError(), newRepoMockSongBase([]int{1, 2, 3}), 0, C.ErrDatabase},
		{"song database error", args{[]int{1, 2, 3}, "collect-2"}, newRepoMockCollectBase([]int{1}), newRepoMockSongServerError(), 0, C.ErrDatabase},
	}

	log.Logger = logrus.New()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &Service{}
			repo.Collect = tt.collect
			repo.Song = tt.song
			got, err := srv.AddCollect(tt.args.songs, tt.args.title)
			if err != tt.err {
				t.Errorf("Service.AddCollect() error = %v, wantErr %v", err, tt.err)
				return
			}
			if got != tt.want {
				t.Errorf("Service.AddCollect() = %v, want %v", got, tt.want)
			}
		})
	}
}
