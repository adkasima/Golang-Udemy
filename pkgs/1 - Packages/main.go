package main

import (
	"fmt"
	"module/helper"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Writing from main file")
	helper.Write()

	error := checkmail.ValidateFormat("devbook@gmail.com")
	fmt.Println(error)
}