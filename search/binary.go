package search

func Binary(array []int, query int) (int, error) {
	low := 0
	high := len(array) - 1

	for low <= high {
		mid := (low + high) / 2

		switch {
		case array[mid] == query:
			return mid, nil
		case array[mid] < query:
			low = mid + 1
		case array[mid] > query:
			high = mid - 1
		}
	}

	return -1, ErrNotFound
}
