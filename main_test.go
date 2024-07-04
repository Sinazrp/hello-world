package main

import "testing"

var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isError  bool
}{
	{"valid-data", 100.0, 10, 10, false},
	{"invalid-data", -100.0, 0.0, 0.0, true},
	{"invalid-data", 50.0, 10.0, 5.0, false},
	{"invalid-data", -1.0, -777, 0.0012870013, false},
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)
		if tt.isError && err == nil {
			t.Errorf("%s: expected error", tt.name)
		}
		if !tt.isError && err != nil {
			t.Errorf("%s: expected no error, got %v", tt.name, err)
		}
		if got != tt.expected {
			t.Errorf("%s: expected %v, got %v", tt.name, tt.expected, got)
		}

	}

}

//func TestDivide(t *testing.T) {
//	_, err := divide(10.0, 1.0)
//	if err != nil {
//		t.Error("got error when we should not", err)
//	}
//
//}
//
//func TestDivide2(t *testing.T) {
//	_, err := divide(10.0, 0)
//	if err == nil {
//		t.Error("not got error when we should", err)
//	}
