package main

import "testing"

func TestDivide(t *testing.T) {
	_, err := divide(10.0, 1.0)
	if err != nil {
		t.Error("got error when we should not", err)
	}

}

func TestDivide2(t *testing.T) {
	_, err := divide(10.0, 0)
	if err == nil {
		t.Error("not got error when we should", err)
	}

}
