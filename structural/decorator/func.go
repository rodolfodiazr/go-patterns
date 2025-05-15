package decorator

import "fmt"

// Function type
type PriceCalculator func(price float64) float64

// Loggin adds logging before and after execution
func Loggin(pc PriceCalculator) PriceCalculator {
	return func(price float64) float64 {
		fmt.Println("Starting price calculation: ", price)
		result := pc(price)
		fmt.Println("Ending price calculation: ", result)
		return result
	}
}
