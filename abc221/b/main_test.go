package main

import (
	"testing"
)

func Test_funcA(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "funcA_1",
			args: args{s: "abc", t: "abc"},
			want: true,
		},
		{
			name: "funcA_2",
			args: args{s: "aabb", t: "bbaa"},
			want: false,
		},
		{
			name: "funcA_3",
			args: args{s: "abcde", t: "abcde"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := funcA(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("funcA() = %v, want %v", got, tt.want)
			}
		})
	}
}
