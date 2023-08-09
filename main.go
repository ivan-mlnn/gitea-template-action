package main

import (
	"fmt"
	"os"
)

func main() {
	//_ = gha.GetInput("username")
	//gha.
	env := os.Environ()
	for _, s := range env {
		fmt.Println(s)
	}
}
