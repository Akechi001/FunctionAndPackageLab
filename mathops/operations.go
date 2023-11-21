package mathops

// Add performs addition of two numbers.
func Add(a, b int) int {
	return a + b
}

// Subtract performs subtraction of two numbers.
func Subtract(a, b int) int {
	return a - b
}

// Multiply performs multiplication of two numbers.
func Multiply(a, b int) int {
	return a * b
}

// Divide performs division of two numbers.
// It returns both the quotient and remainder.
func Divide(a, b int) (int, int) {
	return a / b, a % b
}
