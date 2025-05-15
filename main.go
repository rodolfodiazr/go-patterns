package main

import (
	"fmt"

	"github.com/rodolfodiazr/go-patterns/creational/singleton"
	"github.com/rodolfodiazr/go-patterns/structural/decorator"
)

func main() {
	fmt.Println("--------------------------\n# Decorator")
	decorator.Run()

	fmt.Println("\n--------------------------\n# Singleton")
	singleton.Run()
}
