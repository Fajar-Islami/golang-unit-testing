package math_test

import (
	"testing"
	"time"

	. "github.com/Fajar-Islami/golang-testing/math" // dibuat seolah2 berada dalam package yang sama
)

func TestAdd(t *testing.T) {
	testTable := []struct {
		a, b            int
		expectedOutcome int
	}{
		{a: 1, b: 2, expectedOutcome: 3},
		{a: -1, b: 2, expectedOutcome: 1},
		{a: 1, b: -2, expectedOutcome: -1},
		{a: -1, b: -2, expectedOutcome: -3},
		{a: 0, b: 0, expectedOutcome: 0},
	}

	// Seolah2 1 test yang diren
	for _, test := range testTable {
		result := Add(test.a, test.b)
		if result != test.expectedOutcome {
			t.Logf("Got: %d but expect %d", result, test.expectedOutcome)
			t.Fail()
		}
	}
}

func TestAddWithHierarchical(t *testing.T) {

	t.Run("a=positive", func(t *testing.T) {
		a := 10
		// Subtest
		t.Run("b=positive", func(t *testing.T) {
			result := Add(a, 5)
			if result != 15 {
				t.Logf("Got: %d but expect %d", result, 15)
				t.Fail()
			}

		})
		t.Run("b=negative", func(t *testing.T) {
			result := Add(a, -5)
			if result != 5 {
				t.Logf("Got: %d but expect %d", result, 5)
				t.Fail()
			}

		})
	})

}

func TestNeedsToBeSkip(t *testing.T) {
	t.Skip("This test will be skip")
}

func TestCallToDb(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip because calling db is way to long.")
	}
	<-time.After(3 * time.Second)

}
