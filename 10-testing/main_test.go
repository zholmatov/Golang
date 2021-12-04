package main

import "testing"

var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
}

func TestDivision(t *testing.T) {

	for _, tt := range tests {

		got, err := divide(tt.dividend, tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error("expected error but got none")
			}
		} else {
			if err != nil {
				t.Error("got the error, but wasn't expecting any", err.Error())
			}
		}

		if got != tt.expected {
			t.Errorf("expected %f, but got %f", tt.expected, got)
		}
	}
}

func TestDivide(t *testing.T) {
	_, err := divide(10.0, 1.0)

	if err != nil {
		t.Error("Having error while it should not have")
	}
}

func TestBadDivide(t *testing.T) {
	_, err := divide(10.0, 0)

	if err == nil {
		t.Error("Did not throw an error when it should have")
	}
}
