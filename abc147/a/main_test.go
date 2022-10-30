package main

import "testing"

func TestFuncA(t *testing.T) {
	type args struct {
		a1 int
		a2 int
		a3 int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestFuncA_win",
			args: args{a1: 5, a2: 7, a3: 9},
			want: "win",
		},
		{
			name: "TestFuncA_bust",
			args: args{a1: 13, a2: 7, a3: 2},
			want: "bust",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := funcA(tt.args.a1, tt.args.a2, tt.args.a3); got != tt.want {
				t.Errorf("funcA() = %v, want %v", got, tt.want)
			}
		})
	}
}
