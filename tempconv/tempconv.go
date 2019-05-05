// Package tempconv performs Celsius to Fahrenheit conversions
package tempconv

import "fmt"

// create our custom data type
type Celsius float64
type Fahrenheit float64

// declare our variables
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

// overrite the method they implement which is from
// the float64 type
func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}
