package main

import "testing"

func Test_funcA(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "funcA_1",
			args: []int{1, 6, 3},
			want: 5,
		},
		{
			name: "funcA_2",
			args: []int{11, 5, 5},
			want: 6,
		},
		{
			name: "funcA_3",
			args: []int{100, 100, 100},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := funcA(tt.args); got != tt.want {
				t.Errorf("funcA() = %v, want %v", got, tt.want)
			}
		})
	}
}
