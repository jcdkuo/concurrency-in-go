package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	salutation := "hello"

	sayHello := func() {
		defer wg.Done()
		salutation = "welcome"
	}

	wg.Add(1)
	go sayHello()
	wg.Wait()

	fmt.Println(salutation)

}
