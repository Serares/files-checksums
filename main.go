package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")

	server := CreateApiServer(":3000")

	server.Run()
}
