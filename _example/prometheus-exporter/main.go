package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	weather "github.com/oscaromeu/openweathermap"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type OpenweatherCollector struct {
	ApiKey      string
	Location    string
	temperature *prometheus.Desc
}

func NewOpenweatherCollector(apikey string, location string) *OpenweatherCollector {
	return &OpenweatherCollector{
		ApiKey:      apikey,
		Location:    location,
		temperature: prometheus.NewDesc("openweather_temperature", "Current temperature in degrees celsius", []string{}, nil),
	}
}

func (owm *OpenweatherCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(owm, ch)
}

func (owm *OpenweatherCollector) Collect(ch chan<- prometheus.Metric) {
	conditions, _ := weather.Get(owm.Location, owm.ApiKey)
	ch <- prometheus.MustNewConstMetric(owm.temperature, prometheus.CounterValue, conditions.Temperature.Celsius())
}

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

	listenAddr := flag.String("listen-address", ":8080", "The address to listen on for metrics requests.")
	flag.Parse()

	http.Handle("/metrics", promhttp.Handler())
	log.Println("Listening on", *listenAddr)
	log.Fatal(http.ListenAndServe(*listenAddr, nil))

	weatherCollector := NewOpenweatherCollector(key, location)
	prometheus.MustRegister(weatherCollector)

}
