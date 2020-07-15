package singleton

import (
	"testing"
)

func TestGetInstance(t *testing.T) {
	counter1 := GetInstance()
	if counter1 == nil {
		t.Error("Expected pointer to Singleton after calling GetInstance(), not nil")
	}
	expectedCounter := counter1

	currentcount := counter1.AddOne()
	if currentcount != 1 {
		t.Error("Expected AddOne to increase the counter by 1 but got ", currentcount)
	}

	counter2 := GetInstance()
	if counter2 != expectedCounter {
		t.Error("expected to get the same counter but got ", counter2)
	}

	currentcount = counter2.AddOne()
	if currentcount != 2 {
		t.Error("Expected that after calling AddOne the 2nd time the value of currentcount would be 2 and not ", currentcount)
	}
}
