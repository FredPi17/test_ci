package greeting

import "testing"

func TestHello(t *testing.T) {
	emptyResult := Hello("")

	if emptyResult != "Hello dude" {
		t.Errorf("hello(\"\") failed, expected %v, got %v", "Hello dude", emptyResult)
	}

	result := Hello("Mike")

	if result != "Hello Mike!" {
		t.Errorf("hello(\"Mike\") failed, expected %v, got %v", "Hello Mike", result)
	}
}
