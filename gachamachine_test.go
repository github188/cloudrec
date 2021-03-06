package main

import (
	"os"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
		port string
	}{
		// TODO: Add test cases.
		{"listen-8080", "8080"},
		{"listen-x", "a"},
	}
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = append(os.Args, "")
	for _, tt := range tests {
		os.Args[1] = tt.port
		t.Run(tt.name, func(t *testing.T) {
			go main()
		})
	}
}
