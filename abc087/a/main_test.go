package main

import "testing"

func Test_funcA(t *testing.T) {
	type args struct {
		x int
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "funcA_1",
			args: args{x: 1234, a: 150, b: 100},
			want: 84,
		},
		{
			name: "funcA_2",
			args: args{x: 1000, a: 108, b: 108},
			want: 28,
		},
		{
			name: "funcA_3",
			args: args{x: 579, a: 123, b: 456},
			want: 0,
		},
		{
			name: "funcA_4",
			args: args{x: 7477, a: 549, b: 593},
			want: 405,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := funcA(tt.args.x, tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("funcA() = %v, want %v", got, tt.want)
			}
		})
	}
}
