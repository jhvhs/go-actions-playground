package main

import (
	"fmt"
	"os"

	"github.com/jhvhs/go-actions-playground/pkg"
)

func main() {
	fmt.Println("It's okay")
	pkg.Application{
		Name:   "some name",
		Output: os.Stdout,
	}.Run()
}
