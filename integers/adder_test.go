package integers

import (
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 0
	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}
