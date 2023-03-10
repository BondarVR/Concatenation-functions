package main

import (
	"Concatenation-functions/pkg"
	"fmt"
)

func main() {
	strings := pkg.GenerateStrings(40)
	fmt.Println(pkg.Concat(strings))
	fmt.Println(pkg.СoncatOptimized(strings))
}
