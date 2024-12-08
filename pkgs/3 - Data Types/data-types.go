package main

import (
	"errors"
	"fmt"
)

func main() {

	//BEGIN INTEGER NUMBERS

	// int8, int16, int32, int64

	number := -10000000
	fmt.Println(number)

	var number2 uint32 = 10000
	fmt.Println(number2)

	//nickname
	//RUNE = INT32
	var number3 rune = 12456
	fmt.Println(number3)

	//BYTE = UINT8
	var number4 byte = 123
	fmt.Println(number4)

	//END INTEGER NUMBERS

	//BEGIN FLOAT NUMBERS

	// float32, float64
	var realNumber1 float32 = 123.45
	fmt.Println(realNumber1)

	var realNumber2 float32 = 1230000000.45
	fmt.Println(realNumber2)

	realNumber3 := 12345.45
	fmt.Println(realNumber3)

	//END FLOAT NUMBERS

	//BEGIN STRINGS
	var str string = "Text"
	fmt.Println(str)

	str2 := "Text2"
	fmt.Println(str2)

	char := "a"//OUT: string a
	fmt.Println(char)

	char2 := 'a'//OUT: rune 97 | Which is the position of 'a' at ascii table
	fmt.Println(char2)

	//END STRINGS

	//BEGIN EXTRAS

	var text string
	fmt.Println(text) //OUT: ' ' print a whitespace, which is the initial state of a string variable

	var text2 int16
	fmt.Println(text2) //OUT: 0 | Which is the initial state of a int16 variable

	var boolean bool = false
	fmt.Println(boolean)

	var e error
	fmt.Println(e) //OUT: <nil>

	var outOfCoffe error = errors.New("No more coffe available!!!")
	fmt.Println(outOfCoffe)

	var ErrOutOfTea = fmt.Errorf("no more tea available")
	fmt.Println(ErrOutOfTea)
}