package main

import (
	"fmt"

	"example/main/goodbye"
	"example/main/hello"
)

func main() {
	helloValue := hello.Hello()
	fmt.Println(helloValue)
	hello.NewHello2()
	goodbye.ByeBye()
}
