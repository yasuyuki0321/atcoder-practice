package main

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

func Test_funcA(t *testing.T) {
	tests := []struct {
		name  string
		input io.Reader
		want  string
	}{
		{
			name:  "funcA_1",
			input: bytes.NewBufferString("1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26"),
			want:  "abcdefghijklmnopqrstuvwxyz",
		},
		{
			name:  "funcA_2",
			input: bytes.NewBufferString("2 1 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26"),
			want:  "bacdefghijklmnopqrstuvwxyz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := funcA(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("funcA() = %v, want %v", got, tt.want)
			}
		})
	}
}
