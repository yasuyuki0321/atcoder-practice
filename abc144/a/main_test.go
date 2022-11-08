package main

import "testing"

func Test_funcA(t *testing.T) {
	type args struct {
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
			args: args{a: 2, b: 5},
			want: 10,
		},
		{
			name: "funcA_2",
			args: args{a: 5, b: 10},
			want: -1,
		},
		{
			name: "funcA_3",
			args: args{a: 9, b: 9},
			want: 81,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := funcA(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("funcA() = %v, want %v", got, tt.want)
			}
		})
	}
}
