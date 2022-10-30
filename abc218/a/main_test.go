package main

import "testing"

func Test_funcA(t *testing.T) {
	type args struct {
		s string
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "funcA_1",
			args: args{s: "oooxoox", n: 4},
			want: "No",
		},
		{
			name: "funcA_2",
			args: args{s: "ooooooo", n: 7},
			want: "Yes",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := funcA(tt.args.s, tt.args.n); got != tt.want {
				t.Errorf("funcA() = %v, want %v", got, tt.want)
			}
		})
	}
}
