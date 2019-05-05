package tempconv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

/*
NB: the Celsius and Fahrenheit names could be visible because
1. they start with Capital letters, they are exported
2. because the files belong to same package, the exported names
could be seen in other files that belong to same package name
*/