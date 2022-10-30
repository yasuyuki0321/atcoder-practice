package main

import "testing"

func Test_funcA(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "funcA_1",
			args: args{s: "redcoder"},
			want: 1,
		},
		{
			name: "funcA_2",
			args: args{s: "vvvvvv"},
			want: 0,
		},
		{
			name: "funcA_3",
			args: args{s: "abcdabc"},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := funcA(tt.args.s); got != tt.want {
				t.Errorf("funcA() = %v, want %v", got, tt.want)
			}
		})
	}
}
