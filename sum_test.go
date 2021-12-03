package main

import "testing"

func TestSum(t *testing.T) {
	arr := [5]int{1, 2, 3, 4, 5}
	got := Sum(arr)
	want := 15
	if got != want {
		t.Errorf("expect %d but got %d, %v", want, got, arr)
	}
}