package main

import (
	"testing"
)

// check if function gives root or nothing in case of root is absent
func TestSearchOneRoot(t *testing.T)  {

	if SearchOneRoot != nil {
		t.Error("Function works properly")
	}
}

// check if function gives roots or nothing in case of root is absent
func TestSearchAllRoots(t *testing.T)  {

}
