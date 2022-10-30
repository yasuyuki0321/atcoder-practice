package main

import "testing"

func Test_funcA(t *testing.T) {
	type args struct {
		o string
		e string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "funaA_1",
			args: args{o: "xyz", e: "abc"},
			want: "xaybzc",
		},
		{
			name: "funaA_2",
			args: args{o: "atcoderbeginnercontest", e: "atcoderregularcontest"},
			want: "aattccooddeerrbreeggiunlnaerrccoonntteesstt",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := funcA(tt.args.o, tt.args.e); got != tt.want {
				t.Errorf("funcA() = %v, want %v", got, tt.want)
			}
		})
	}
}
