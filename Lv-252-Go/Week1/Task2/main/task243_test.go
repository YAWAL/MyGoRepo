package main

import "testing"

//TestSearchOneRoot check if function gives root or nothing in case of root is absent
func TestSearchOneRoot(t *testing.T)  {

	if !SearchOneRoot(5) {
		t.Error("Function works properly")
	}
}

//TestSearchAllRoots check if function gives roots or nothing in case of root is absent
func TestSearchAllRoots(t *testing.T)  {

}
