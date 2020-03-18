package integers_test

import "testing"
import "fmt"
import "github.com/devaof/learn-go-with-tests/integers"

func TestAdder(t *testing.T) {
	sum := integers.Add(4, 4)
	expected := 8

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}
