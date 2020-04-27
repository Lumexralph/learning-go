// concurrent non-blocking cache
package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func main() {
	val := runtime.GOMAXPROCS(0)
	fmt.Println("the number of cpu is: ", val)

	t := reflect.TypeOf(6)
	fmt.Println(t.String())
	fmt.Println(t)

	v := reflect.ValueOf(3) // a reflect.Value
	fmt.Println(v)          // "3"
	fmt.Printf("%v\n", v)   // "3"
	fmt.Println(v.String()) // NOTE: "<int Value>"

}
