package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Conditions struct {
	Summary     string
	Temperature Temperature
	FeelsLike   Temperature
	TempMin     Temperature
	TempMax     Temperature
}

type Temperature float64

func (t Temperature) Celsius() float64 {
	return float64(t) - 273.15
}

type OWMResponse struct {
	Weather []struct {
		Main string
	}
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  float64 `json:"pressure"`
		Humidity  float64 `json:"humidity"`
	}
}

func ParseResponse(data []byte) (Conditions, error) {
	var resp OWMResponse
	err := json.Unmarshal(data, &resp)
	if err != nil {
		return Conditions{}, fmt.Errorf("invalid API response %s: %w", data, err)
	}
	if len(resp.Weather) < 1 {
		return Conditions{}, fmt.Errorf("invalid API response %s: want at least one Weather element", data)
	}

	conditions := Conditions{
		Summary:     resp.Weather[0].Main,
		Temperature: Temperature(resp.Main.Temp),
		FeelsLike:   Temperature(resp.Main.FeelsLike),
	}
	return conditions, nil
}

type Client struct {
	APIKey     string
	BaseURL    string
	HTTPClient *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		APIKey:     apiKey,
		BaseURL:    "https://api.openweathermap.org",
		HTTPClient: http.DefaultClient,
	}
}

func (c Client) FormatURL(location string) string {
	return fmt.Sprintf("%s/data/2.5/weather?q=%s&appid=%s", c.BaseURL, location, c.APIKey)
}

func (c *Client) GetWeather(location string) (Conditions, error) {
	URL := c.FormatURL(location)
	resp, err := c.HTTPClient.Get(URL)
	if err != nil {
		return Conditions{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return Conditions{}, fmt.Errorf("unexpected response status %q", resp.Status)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Conditions{}, err
	}
	conditions, err := ParseResponse(data)
	if err != nil {
		return Conditions{}, fmt.Errorf("invalid API response %s: %w", data, err)
	}
	return conditions, nil
}

func Get(location, key string) (Conditions, error) {
	c := NewClient(key)
	conditions, err := c.GetWeather(location)
	if err != nil {
		return Conditions{}, err
	}
	return conditions, nil
}
