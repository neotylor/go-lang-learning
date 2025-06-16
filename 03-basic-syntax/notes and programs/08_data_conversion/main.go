package main

import (
	"fmt"
	"strconv"
)

func main() {
	println("Learn Data conversion in GO Lang.")

	var num int = 32

	println("Number is", num)
	fmt.Printf("Number variable type is: %T\n", num)

	var float float64 = float64(num)

	fmt.Printf("Number variable type and value is: %T => %f\n", float, float)

	var str string = strconv.Itoa(num)

	fmt.Printf("Str variable type and value is: %T => %s\n", str, str)

	var strNumb string = "15989"

	numb_string, err := strconv.Atoi(strNumb)

	fmt.Printf("strNumb variable type and value is: %T => %s\n", strNumb, strNumb)
	fmt.Printf("numb_string variable type and value is: %T => %d and any error is: %s\n", numb_string, numb_string, err)

	var strFloat string = "5.35"

	float_str, err := strconv.ParseFloat(strFloat, 64)

	fmt.Printf("strFloat variable type and value is: %T => %s\n", strFloat, strFloat)
	fmt.Printf("float_str variable type and value is: %T => %f and any error is: %s\n", float_str, float_str, err)

}
