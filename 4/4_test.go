package four

import "testing"

func TestFirst(t *testing.T) {
	ans, err := First()
	if err != nil {
		t.Fatal(err)
	}

	if ans != 239 {
		t.Fail()
	}
}

func TestSecond(t *testing.T) {
	ans, err := Second()
	if err != nil {
		t.Fatal(err)
	}

	if ans != 188 {
		t.Fail()
	}
}
