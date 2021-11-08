package search

import (
	"testing"
)

func TestBinary(t *testing.T) {
	for _, test := range searchTests {
		t.Run(test.name, func(t *testing.T) {
			actualIndex, actualError := Binary(test.array, test.query)

			if actualIndex != test.expectedIndex {
				t.Errorf("expected '%d', got '%d'", test.expectedIndex, actualIndex)
			}

			if actualError != test.expectedError {
				t.Errorf("expected '%s', got '%s'", test.expectedError, actualError)
			}
		})
	}
}

func BenchmarkBinary(b *testing.B) {
	array := generateBenchmarkArray()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Binary(array, i)
	}
}
