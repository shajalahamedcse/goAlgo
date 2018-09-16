package utilities

import (
	"testing"
)

func TestIntegerComparator(t *testing.T) {

	tests := [][]interface{}{
		{23, 23, 0},
		{11, 12, -1},
		{12, 11, 1},
		{0, 0, 0},
	}

	for _, test := range tests {
		result := IntegerComparator(test[0], test[1])
		expected := test[2]
		if result != expected {
			t.Errorf("Got %v expected %v", result, expected)
		}
	}
}

func TestStringComparator(t *testing.T) {

	tests := [][]interface{}{
		{"shajal", "shajal", 1},
		{"shajal", "shajal", 0},
	}

	for _, test := range tests {
		result := StringComparator(test[0], test[1])
		expected := test[2]

		if result != expected {
			t.Errorf("Got %v expected %v", result, expected)
		}
	}
}
