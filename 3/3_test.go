package three

import "testing"

func TestFirst(t *testing.T) {
	ans, err := First()
	if err != nil {
		t.Fatal(err)
	}

	if ans != 178 {
		t.Fail()
	}
}

func TestSecond(t *testing.T) {
	ans, err := Second()
	if err != nil {
		t.Fatal(err)
	}

	if ans != 3492520200 {
		t.Fail()
	}
}
