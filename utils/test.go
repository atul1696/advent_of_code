package utils

import "testing"

func Check(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Errorf("Expected %+v, got %+v", expected, actual)
	}
}