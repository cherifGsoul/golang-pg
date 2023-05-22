package main

import (
	"fmt"
	"interfaces/someService"
)

func main() {
	var service = someService.NewSomeServiceStruct()
	fmt.Println(service.Greet())
	fmt.Println(service.GreetingSize())
}
