package main

import (
	"fmt"
	"os"

	weather "github.com/oscaromeu/openweathermap"
)

func main() {
	key := os.Getenv("OPENWEATHERMAP_API_KEY")
	if key == "" {
		fmt.Fprintln(os.Stderr, "Please set the environment variable OPENWEATHERMAP_API_KEY.")
		os.Exit(1)
	}
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s LOCATION\n\nExample: %[1]s Barcelona,ES\n", os.Args[0])
		os.Exit(1)
	}
	location := os.Args[1]
	conditions, err := weather.Get(location, key)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("%s %.1fºC \n",
		conditions.Summary,
		conditions.Temperature.Celsius(),
	)
}
