package main

import (
	"fmt"

	"vgoawesomeProject/subpkg"
)

func main() {
	r := subpkg.GetRouter()
	fmt.Printf("%v\n", r)
}
