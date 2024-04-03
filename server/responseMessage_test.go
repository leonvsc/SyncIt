package main

import (
	"reflect"
	"testing"
)

func TestSendFile(t *testing.T) {
	tests := []struct {
		name      string
		headerMap map[string]string
		want      []byte
	}{
		{
			name:      "Invalid file path",
			headerMap: map[string]string{"Path": "nonexistent.txt"},
			want:      nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sendFile(tt.headerMap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sendFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendOkToClient(t *testing.T) {
	want := []byte("...") // Replace "..." with expected byte slice
	if got := sendOkToClient(); !reflect.DeepEqual(got, want) {
		t.Errorf("sendOkToClient() = %v, want %v", got, want)
	}
}
