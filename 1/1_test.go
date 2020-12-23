package one

import "testing"

func TestFirst(t *testing.T) {
	v, err := First()
	if err != nil {
		t.Fatal(err)
	}

	if v != 542619 {
		t.Fail()
	}
}

func TestSecond(t *testing.T) {
	v, err := Second()
	if err != nil {
		t.Fatal(err)
	}

	if v != 32858450 {
		t.Fail()
	}
}
