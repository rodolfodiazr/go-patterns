package singleton

import (
	"fmt"
	"sync"
)

type Config struct {
	AppName   string
	Port      int
	DebugMode bool
}

var (
	configInstance *Config
	once           sync.Once
)

// GetConfig returns the singleton config instance
func GetConfig() *Config {
	once.Do(
		func() {
			fmt.Println("Load configuration")
			// You could read from a file/env
			configInstance = &Config{
				AppName:   "Sample App",
				Port:      8080,
				DebugMode: true,
			}
		},
	)
	return configInstance
}

// Run demonstrates the Singleton Pattern
func Run() {
	cfg1 := GetConfig()
	fmt.Printf("Config 1: %+v\n", cfg1)

	cfg2 := GetConfig()
	fmt.Printf("Config 2: %+v\n", cfg2)
	fmt.Println("Same instance?", cfg1 == cfg2)
}
