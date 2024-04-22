package main

import "testing"

func TestHitungLuasPersegi(t *testing.T) {
	tests := []struct {
		name   string
		sisi   float64
		expect float64
	}{
		{"Sisi 4", 4.0, 16.0},
		{"Sisi 5", 5.0, 25.0},
		{"Sisi 7.5", 7.5, 56.25},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HitungLuasPersegi(tt.sisi)
			if result != tt.expect {
				t.Errorf("Got %.2f, expected %.2f", result, tt.expect)
			}
		})
	}
}
