package game

import (
	"testing"
)

func Test_checkDuplicateInts(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"success", args{[]int{1, 2, 3, 4, 5}}, false},
		{"fail-1", args{[]int{1, 2, 1}}, true},
		{"fail-2", args{[]int{1, 1}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkDuplicateInts(tt.args.nums); got != tt.want {
				t.Errorf("checkDuplicateInts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lyricsJoin(t *testing.T) {
	type args struct {
		lyrics []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"success", args{[]int{1, 2, 3, 4}}, "1,2,3,4"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lyricsJoin(tt.args.lyrics); got != tt.want {
				t.Errorf("lyricsJoin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMax(t *testing.T) {
	type args struct {
		l []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"success-1", args{[]int{1, 3, 7, 5, 4, 8, 10, 2, 6}}, 10},
		{"success-2", args{[]int{1}}, 1},
		{"negative", args{[]int{3, 8, -1, 5}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMax(tt.args.l); got != tt.want {
				t.Errorf("findMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
