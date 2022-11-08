package main

import "testing"

func Test_funcA(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "funcA_1",
			args: args{n: 10},
			want: "Yes",
		},
		{
			name: "funcA_2",
			args: args{n: 50},
			want: "No",
		},
		{
			name: "funcA_3",
			args: args{n: 81},
			want: "Yes",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := funcA(tt.args.n); got != tt.want {
				t.Errorf("funcA() = %v, want %v", got, tt.want)
			}
		})
	}
}
