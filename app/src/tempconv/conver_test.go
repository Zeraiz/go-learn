package tempconv

import "testing"

func TestCToF(t *testing.T) {
	expected := Fahrenheit(32)
	zerC := Celsius(0)
	if ret := CToF(zerC); ret != expected {
		t.Errorf("0 celsius to fahrenheit is 32, actual: %s", ret)
	}
}
