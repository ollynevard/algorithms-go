package search

type searchTest struct {
	name          string
	array         []int
	query         int
	expectedIndex int
	expectedError error
}

var searchTests = []searchTest{
	{"ItemExists", []int{2, 3, 5, 7, 11}, 5, 2, nil},
	{"ItemNotFound", []int{2, 3, 5, 7, 11}, 13, -1, ErrNotFound},
	{"EmptyArray", []int{}, 2, -1, ErrNotFound},
}

func generateBenchmarkArray() []int {
	var array []int

	for i := 0; i < 1000; i++ {
		array = append(array, i)
	}

	return array
}
