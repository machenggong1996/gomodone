package gomodone

import "fmt"

// say Hi to someone
func SayHi(name string, age int) string {
	fmt.Println("age is :", age)
	return fmt.Sprintf("Hi, %s", name)
}
