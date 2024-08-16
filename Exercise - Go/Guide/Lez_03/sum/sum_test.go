package sum

import "testing"

func TestSum(t *testing.T) {
	// arrange: prepare	dati di input
	var num1, num2 int
	num1 = 3
	num2 = 2
	// act
	result := Sum(num1, num2)
	// assert
	if result != 5 {
		t.Errorf("Expected %d, got %d", 5, result)
	}
}

// go test
// go test -v
// go test -v -cover
