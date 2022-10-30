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
		want bool
	}{
		{
			name: "funcA_1",
			args: args{a: 2, b: 4, c: 6},
			want: true,
		},
		{
			name: "funcA_2",
			args: args{a: 2, b: 5, c: 6},
			want: false,
		},

		{
			name: "funcA_3",
			args: args{a: 3, b: 2, c: 1},
			want: true,
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
