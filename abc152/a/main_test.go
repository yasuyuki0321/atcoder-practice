package main

import "testing"

func Test_funcA(t *testing.T) {
	type args struct {
		n int
		m int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "funcA_1",
			args: args{n: 3, m: 3},
			want: "Yes",
		},
		{
			name: "funcA_2",
			args: args{n: 3, m: 2},
			want: "No",
		},
		{
			name: "funcA_1",
			args: args{n: 1, m: 1},
			want: "Yes",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := funcA(tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("funcA() = %v, want %v", got, tt.want)
			}
		})
	}
}
