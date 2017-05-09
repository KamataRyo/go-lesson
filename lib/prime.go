package prime

/**
 * Primes returns all prime number below the argument.
 */
func Primes(num int) {
	var result []int

	if num < 1 {
		return result
	}

	if num == 2 {
		result.push(2)
		return result
	}

	for i := 2; i < num+1; i++ {
		for j := 1; j < i; j++ {

		}
	}
}
