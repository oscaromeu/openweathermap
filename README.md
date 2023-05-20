# OpenWeatherMap Go API

This project is a simple Go (golang) package for use with openweathermap.org's API. Note that it only outputs a few metrics. For a more detailed cover of the OpenWeatherMap API, you can use [briandowns/openweathermap](https://github.com/briandowns/openweathermap).

To use the OpenWeatherMap API, you need to obtain an API key. Sign up [here](http://home.openweathermap.org/users/sign_up). Once you have your key, create an environment variable called `OPENWEATHERMAP_API_KEY`. Start coding!

```
export OPENWEATHERMAP_API_KEY=<API_KEY_VALUE>
curl -s "https://api.openweathermap.org/data/2.5/weather?q=Barcelona,ES&appid=$OPENWEATHERMAP_API_KEY" | jq '.'
```

## Installation

Install the OpenWeatherMap Go API package by running the following command:

```shell
go get github.com/oscaromeu/openweathermap
```

This will download the package and its dependencies to your Go workspace.

## Requirements

+ Go (1.13 or later)

## Features

The OpenWeatherMap Go API package allows you to fetch current weather data by city, including the following metrics:

+ Temperature
+ Summary status


__Note:__ Data is only available in SI metric units.

## Examples

Here's an example code snippet that demonstrates how to use the OpenWeatherMap Go API package:

```go
package main

import (
	"fmt"
	"os"

	"github.com/oscaromeu/openweathermap"
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
	conditions, err := openweathermap.Get(location, key)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("%s %.1fºC\n",
		conditions.Summary,
		conditions.Temperature.Celsius(),
	)
}
```

See the `_examples` directory for more examples of usage.

## Contributions

This is a small personal project, and contributions from others are not currently accepted.

## License

This work is licensed under a [Creative Commons Attribution 4.0 International
License][cc-by].

[![CC BY 4.0][cc-by-image]][cc-by]

[cc-by]: http://creativecommons.org/licenses/by/4.0/

[cc-by-image]: https://i.creativecommons.org/l/by/4.0/88x31.png