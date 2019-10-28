package main

import "testing"

func Test_logic(t *testing.T) {

	tests := []struct {
		name string
		s    string
		want string
	}{
		{name: "01B0", s: "01B0", want: "00"},
		{name: "0BB1", s: "0BB1", want: "1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := logic(tt.s); got != tt.want {
				t.Errorf("logic() = %v, want %v", got, tt.want)
			}
		})
	}
}
