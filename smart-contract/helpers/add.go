package helpers

import "fmt"

//Add adds two number checking for overflow.
func Add(b int, q int) (int, error) {

	// Check overflow
	sum := q + b

	if (sum < q || sum < b) == (b >= 0 && q >= 0) {
		return 0, fmt.Errorf("math: addition overflow occurred %d + %d", b, q)
	}

	return sum, nil
}
