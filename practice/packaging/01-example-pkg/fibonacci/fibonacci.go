// Package fibonacci is an example package that calculates the fibonacci sequence
package fibonacci

// Each number in the sequence is the sum of the two numbers that precede it.
// So, the sequence goes: 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, and so on.
// The mathematical equation describing it is Xn+2= Xn+1 + Xn

// CalcSlow function will take an integer as input, then calculate the fibonacci
// sequence of that number using recursion.
func CalcSlow(n int) int {
	// If we are equal-to or less-than one, we can return the original value
	// to prevent further calculations.
	if n <= 1 {
		return n
	}

	// We will continue to recurse through the function adding numbers until we get back
	// to the original fibonacci number or 1.
	return CalcSlow(n-1) + CalcSlow(n-2)
}

// CalcFaster function will take an integer as input, then calculate the fibonacci
// sequence of that number using iteration.
func CalcFaster(n int) int {
	// If we are equal-to or less-than one, we can return the original value
	// to prevent further calculations.
	if n <= 1 {
		return n
	}

	var n1, n2 = 0, 1

	// Range through our loop until our index reaches the original
	// fibonacci number passed in.
	for i := 2; i < n; i++ {
		// For each iteration, we shift the n2 variable down to the n1 variable, and set
		// n2 to the value of n1 + n2.
		n1, n2 = n2, n1+n2
	}

	// Once we reach the original fibonacci number, n1 and n2 should be at the values of the previous
	// two fibonacci numbers. We can return the sum of these two numbers as the new fibonacci number.
	return n1 + n2
}
