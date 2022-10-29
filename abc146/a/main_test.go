package main

import "testing"

func TestFuncA(t *testing.T) {
	tests := []struct {
		name string
		args string
		want int
	}{
		{
			name: "normal_mon",
			args: "MON",
			want: 6,
		},
		{
			name: "normal_tue",
			args: "TUE",
			want: 5,
		},
		{
			name: "normal_wed",
			args: "WED",
			want: 4,
		},
		{
			name: "normal_aaa",
			args: "AAA",
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// tt := tt
			// t.Parallel()

			if got := funcA(tt.args); got != tt.want {
				t.Errorf("FuncA() = %v, want %v", got, tt.want)
			}
		})
	}
}
