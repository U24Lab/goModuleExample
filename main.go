package main

import (
	"fmt"

	"example/goodbye"
	"example/hello"
)

func main() {
	helloValue := hello.Hello()
	fmt.Println(helloValue)
	hello.NewHello2()
	goodbye.ByeBye()
}
