package main

import (
	_ "os"
)

type Celsius int
type Fahrenheit int

func (c Celsius) convertToFahrenheit() Fahrenheit {
	return Fahrenheit(100 * c)
}

func (c *Celsius) convertToFahrenheitWithPtr() Fahrenheit {
	return Fahrenheit(100 * *c)
}

func main() {
	celsius := Celsius(100)
	celsius.convertToFahrenheit() // this can be accessible with method (receiver)

	celsiusPtr := Celsius(100)
	celsiusPtr.convertToFahrenheitWithPtr() // compiler will add a * itself
}
