package singleton

import "testing"

func TestGetInstance(t *testing.T) {
	counter1 := GetInstance()

	if counter1 == nil {
		// Test 1: First time GetInstance call
		t.Error("Expected pointer to Singleton after calling GetInsance(), not nil")
	}

	expectedCounter := counter1

	currentCount := counter1.AddOne()

	if currentCount != 1 {
		// Test 2: Test AddOne using counter1
		t.Errorf(`After calling for the first time to count, the count must be 1
			but it is %d\n`,
			currentCount,
		)
	}

	counter2 := GetInstance()

	if counter2 != expectedCounter {
		// Test 3: counter1 and counter2 should have the same instance
		t.Error("Expected same instance in counter2 but it got a different instance")
	}

	currentCount = counter2.AddOne()

	if currentCount != 2 {
		// Test 4: Test AddOne using counter2
		t.Errorf(`After calling 'AddOne' using counter2, the currentCount must
			be 2, but was %d\n`,
			currentCount,
		)
	}
}