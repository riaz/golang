package main

import "testing"

func TestConv(t *testing.T) {
	num, err := conv("123455")
	if err != nil {
		t.Fatal(err)
	}
	if num != 123455 {
		t.Fatal("Numbers dont match")
	}
}

func TestFailConv(t *testing.T) {
	_, err := conv("")
	if err == nil {
		t.Fatal(err)
	}
}
