package six

import "testing"

func TestFirst(t *testing.T) {
	ans, err := First()
	if err != nil {
		t.Fatal(err)
	}

	if ans != 6259 {
		t.Fail()
	}
}

func TestSecond(t *testing.T) {
	ans, err := Second()
	if err != nil {
		t.Fatal(err)
	}

	if ans != 534 {
		t.Fail()
	}
}
