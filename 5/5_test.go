package five

import "testing"

func TestFirst(t *testing.T) {
	ans, err := First()
	if err != nil {
		t.Fatal(err)
	}

	if ans != 991 {
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
