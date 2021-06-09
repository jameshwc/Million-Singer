package stat

import (
	"reflect"
	"testing"
)

func TestCalDiskUsage(t *testing.T) {
	tests := []struct {
		name string
		want Disk
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalDiskUsage(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalDiskUsage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetServer(t *testing.T) {
	tests := []struct {
		name string
		want *Server
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetServer(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetServer() = %v, want %v", got, tt.want)
			}
		})
	}
}
