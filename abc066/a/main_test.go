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
			args: args{a: 700, b: 600, c: 780},
			want: 1300,
		},
		{
			name: "funcA_2",
			args: args{a: 10000, b: 10000, c: 10000},
			want: 20000,
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
