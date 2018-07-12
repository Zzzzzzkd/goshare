package sum

import "time"

func Sum(x, y int) int {
	return x + y
}

func SlowSum(x, y int) int {
	time.Sleep(1 * time.Microsecond)
	return x + y
}
