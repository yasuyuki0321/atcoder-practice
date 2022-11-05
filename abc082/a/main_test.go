package main

import "testing"

func Test_funcA(t *testing.T) {
	type args struct {
		a float64
		b float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "funcA_1",
			args: args{a: 1, b: 3},
			want: 2,
		},
		{
			name: "funcA_2",
			args: args{a: 7, b: 4},
			want: 6,
		},
		{
			name: "funcA_3",
			args: args{a: 5, b: 5},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			funcA(tt.args.a, tt.args.b)
		})
	}
}
