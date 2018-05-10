package environment

import (
	"testing"
)

func TestSingleMissing(t *testing.T) {
	_, err := ReadOptions(Options{
		"FOO": "",
	})
	if err == nil {
		t.Log("no error on missing environment variable")
		t.Fail()
	}
	t.Log(err)
}

func TestTwoMissing(t *testing.T) {
	_, err := ReadOptions(Options{
		"FOO": "",
		"BAR": "",
	})
	if err == nil {
		t.Log("no error on missing environment variable")
		t.Fail()
	}
	t.Log(err)
}

func TestMultipleMissing(t *testing.T) {
	_, err := ReadOptions(Options{
		"FOO":     "",
		"BAR":     "",
		"FLOOGOO": "",
	})
	if err == nil {
		t.Log("no error on missing environment variable")
		t.Fail()
	}
	t.Log(err)
}
