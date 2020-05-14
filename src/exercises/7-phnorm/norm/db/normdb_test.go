package db

import (
	"testing"

	_ "github.com/lib/pq"
)

func TestSetupDB(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "can connect",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetupDB()
		})
	}
}

//GetAllPNs
func TestGetAllPNs(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "can query",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if nums := GetAllPNs(); len(nums) == 0 {
				t.Errorf("GetAllPNs() no values returned")
			}
		})
	}
}
