package main

// Import format and math functions
import (
	"fmt"
	"math"
)

// Showcasing basic math operations
func main() {
	// Addition
	fmt.Println("Addition:", 5+3)

	// Subtraction
	fmt.Println("Subtraction:", 5-3)

	// Multiplication
	fmt.Println("Multiplication:", 5*3)

	// Division
	fmt.Println("Division:", 5.0/3.0)

	// Modulus (remainder)
	fmt.Println("Modulus:", 5%3)

	// Power
	fmt.Println("Power:", math.Pow(5, 3))

	// Square root
	fmt.Println("Square root:", math.Sqrt(25))

	// Absolute value
	fmt.Println("Absolute value:", math.Abs(-5))

	// Maximum
	fmt.Println("Maximum:", math.Max(5, 3))

	// Minimum
	fmt.Println("Minimum:", math.Min(5, 3))
}
