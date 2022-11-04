package main

import "testing"

func Test_funcA(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "funcA_1",
			args: args{x: 3, y: 4},
			want: "3333",
		},
		{
			name: "funcA_2",
			args: args{x: 7, y: 7},
			want: "7777777",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := funcA(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("funcA() = %v, want %v", got, tt.want)
			}
		})
	}
}
