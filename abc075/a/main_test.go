package main

import "testing"

func Test_funcA(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "funcA_1",
			args: args{a: 5, b: 7, c: 5},
			want: 7,
		},
		{
			name: "funcA_2",
			args: args{a: 1, b: 1, c: 7},
			want: 7,
		},
		{
			name: "funcA_3",
			args: args{a: -100, b: 100, c: 100},
			want: -100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := funcA(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("funcA() = %v, want %v", got, tt.want)
			}
		})
	}
}
