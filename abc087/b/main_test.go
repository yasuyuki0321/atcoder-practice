package main

import "testing"

func Test_funcA(t *testing.T) {
	type args struct {
		a int
		b int
		c int
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "funcA_1",
			args: args{a: 1, b: 2, c: 2, x: 100},
			want: 2,
		},
		{
			name: "funcA_2",
			args: args{a: 5, b: 1, c: 0, x: 150},
			want: 0,
		},
		{
			name: "funcA_3",
			args: args{a: 30, b: 40, c: 50, x: 6000},
			want: 213,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := funcA(tt.args.a, tt.args.b, tt.args.c, tt.args.x); got != tt.want {
				t.Errorf("funcA() = %v, want %v", got, tt.want)
			}
		})
	}
}
