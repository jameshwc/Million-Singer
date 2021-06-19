package prom

import (
	"testing"
	"time"
)

func Test_recordServerMetrics(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"success"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go recordServerMetrics()
		})
	}
	time.Sleep(5 * time.Second)
}
