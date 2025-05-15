package decorator

import "fmt"

// Run demonstrates the Decorator pattern
func Run() {
	fmt.Println("\n## OOP-style decorator")
	var notifier Notifier = &EmailNotifier{}

	notifier = &SMSNotifier{Notifier: notifier}
	notifier = &SlackNotifier{Notifier: notifier}

	notifier.Send("This is my message.")

	fmt.Println("\n## Function-based decorator")
	var baseCalc PriceCalculator = func(price float64) float64 {
		return price
	}

	result := Loggin(baseCalc)(100)
	fmt.Println("Result: ", result)
}
