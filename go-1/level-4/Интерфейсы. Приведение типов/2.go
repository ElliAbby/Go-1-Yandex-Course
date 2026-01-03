// использование type assertion

package main

import "fmt"

func PrintType(x interface{}) {
    switch x.(type) {
    case string:
        fmt.Println("string")
    case int:
        fmt.Println("int")
    case bool:
        fmt.Println("bool")
    default:
        fmt.Println("unknown")
    }
}

func main() {
	PrintType("hello") // string
	PrintType(42) // int
	PrintType(true) // bool
	PrintType(3.14) // unknown

	var x interface{} = "hello"
	s, ok := x.(string)
	if ok {
	    fmt.Println(s)
	} else {
	    fmt.Println("not a string")
	}
}
