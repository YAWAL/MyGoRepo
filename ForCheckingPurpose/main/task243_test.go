package main

import "testing"

//TestSearchOneRoot check if function gives root or nothing in case of root is absent
func TestSearchOneRoot(t *testing.T)  {

	value := SearchOneRoot(5)
	if value{
		t.Log("Function works properly")
	}
}

//TestSearchAllRoots check if function gives roots or nothing in case of root is absent
func TestSearchAllRoots(t *testing.T)  {

}
