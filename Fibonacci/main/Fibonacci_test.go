package main

import "testing"

func TestFibonacci(t *testing.T)  {

	testValue := 6

	if  testValue != fibonacci(2){
		t.Errorf("bad math")
	}
	if  testValue == fibonacci(2){
		t.Errorf("good math")
	}
}
