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

func TestService_AddTour(t *testing.T) {
	type args struct {
		collectsID []int
	}
	tourBase := []int{1}
	collectBase := []int{1, 2, 3}
	tests := []struct {
		name    string
		args    args
		want    int
		err     error
		tour    repo.TourRepo
		collect repo.CollectRepo
	}{
		{"success", args{collectBase}, len(tourBase) + 1, nil, newRepoMockTourBase(tourBase), newRepoMockCollectBase(collectBase)},
		{"collect nil fail", args{[]int{}}, 0, C.ErrTourAddFormatIncorrect, newRepoMockTourBase(tourBase), newRepoMockCollectBase(collectBase)},
		{"collect record not found", args{[]int{15, 20, 22}}, 0, C.ErrTourAddCollectsRecordNotFound, newRepoMockTourBase(tourBase), newRepoMockCollectBase(collectBase)},
		{"partly collect record not found", args{[]int{1, 20}}, 0, C.ErrTourAddCollectsRecordNotFound, newRepoMockTourBase(tourBase), newRepoMockCollectBase(collectBase)},
		{"collect record duplicate", args{[]int{1, 1, 2}}, 0, C.ErrTourAddCollectsDuplicate, newRepoMockTourBase(tourBase), newRepoMockCollectBase(collectBase)},
		{"tour database error", args{collectBase}, 0, C.ErrDatabase, newRepoMockTourServerError(), newRepoMockCollectBase(collectBase)},
		{"collect database error", args{collectBase}, 0, C.ErrDatabase, newRepoMockTourBase(tourBase), newRepoMockCollectServerError()},
	}

	log.Logger = logrus.New() // mock logger

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo.Collect = tt.collect
			repo.Tour = tt.tour

			srv := &Service{}
			got, err := srv.AddTour(tt.args.collectsID)
			if err != tt.err {
				t.Errorf("Service.AddTour() error = %v, want err %v", err, tt.err)
				return
			}
			if got != tt.want {
				t.Errorf("Service.AddTour() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetTotalTours(t *testing.T) {
	tourBase := []int{1, 2, 3}
	tests := []struct {
		name string
		want int
		err  error
		tour repo.TourRepo
	}{
		{"success", 3, nil, newRepoMockTourBase(tourBase)},
		{"tour database error", 0, C.ErrDatabase, newRepoMockTourServerError()},
	}

	log.Logger = logrus.New()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo.Tour = tt.tour
			srv := &Service{}
			got, err := srv.GetTotalTours()
			if err != tt.err {
				t.Errorf("Service.GetTotalTours() error = %v, wantErr %v", err, tt.err)
				return
			}
			if got != tt.want {
				t.Errorf("Service.GetTotalTours() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetTour(t *testing.T) {
	type args struct {
		param string
	}
	tests := []struct {
		name  string
		args  args
		want  *model.Tour
		err   error
		tour  repo.TourRepo
		cache repo.CacheRepo
	}{
		{"success", args{"1"}, &model.Tour{1, nil}, nil, newRepoMockTourBase([]int{1}), newRepoMockCacheServerError()},
		{"param fail", args{"1u"}, nil, C.ErrTourIDNotNumber, newRepoMockTourBase([]int{1}), newRepoMockCacheServerError()},
		{"param fail-2", args{"u"}, nil, C.ErrTourIDNotNumber, newRepoMockTourBase([]int{1}), newRepoMockCacheServerError()},
		{"tour id not found", args{"2"}, nil, C.ErrTourNotFound, newRepoMockTourBase([]int{1}), newRepoMockCacheServerError()},
		{"tour id empty", args{""}, nil, C.ErrTourIDNotNumber, newRepoMockTourBase([]int{1}), newRepoMockCacheServerError()},
		{"tour database error", args{"1"}, nil, C.ErrDatabase, newRepoMockTourServerError(), newRepoMockCacheServerError()},
		{"cache used but parse error", args{"1"}, &model.Tour{1, nil}, nil, newRepoMockTourBase([]int{1}), newRepoMockCacheBase()},
		// TODO: cache success test case
	}

	log.Logger = logrus.New()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &Service{}
			repo.Cache = tt.cache
			repo.Tour = tt.tour
			got, err := srv.GetTour(tt.args.param)
			if err != tt.err {
				t.Errorf("Service.GetTour() error = %v, wantErr %v", err, tt.err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetTour() = %v, want %v", got, tt.want)
			}
		})
	}
}
