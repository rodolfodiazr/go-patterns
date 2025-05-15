package main

import (
	"fmt"

	"github.com/rodolfodiazr/go-patterns/behavioral/observer"
	"github.com/rodolfodiazr/go-patterns/behavioral/strategy"
	"github.com/rodolfodiazr/go-patterns/creational/abstractfactory"
	"github.com/rodolfodiazr/go-patterns/creational/builder"
	"github.com/rodolfodiazr/go-patterns/creational/factorymethod"
	"github.com/rodolfodiazr/go-patterns/creational/singleton"
	"github.com/rodolfodiazr/go-patterns/structural/decorator"
	"github.com/rodolfodiazr/go-patterns/structural/facade"
)

func main() {
	fmt.Println("--------------------------\n# Decorator")
	decorator.Run()

	fmt.Println("\n--------------------------\n# Singleton")
	singleton.Run()

	fmt.Println("\n--------------------------\n# Observer")
	observer.Run()

	fmt.Println("\n--------------------------\n# Strategy")
	strategy.Run()

	fmt.Println("\n--------------------------\n# Builder")
	builder.Run()

	fmt.Println("\n--------------------------\n# Factory Method")
	factorymethod.Run()

	fmt.Println("\n--------------------------\n# Abstract Factory")
	abstractfactory.Run()

	fmt.Println("\n--------------------------\n# Facade")
	facade.Run()
}
