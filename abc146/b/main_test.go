package main

import "testing"

func TestFuncA(t *testing.T) {
	type args = struct {
		s string
		n int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "normal",
			args: args{s: "ABCXYZ", n: 2},
			want: "CDEZAB",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := funcA(tt.args.s, tt.args.n); got != tt.want {
				t.Errorf("FuncA() = %v, want %v", got, tt.want)
			}
		})
	}
}
