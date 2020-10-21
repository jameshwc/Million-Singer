package game

import "testing"

func TestService_AddTour(t *testing.T) {
	type args struct {
		collectsID []int
	}
	tests := []struct {
		name    string
		srv     *Service
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &Service{}
			got, err := srv.AddTour(tt.args.collectsID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.AddTour() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Service.AddTour() = %v, want %v", got, tt.want)
			}
		})
	}
}
