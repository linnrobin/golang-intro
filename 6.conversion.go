package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int32 = 10
	var y int64 = int64(x)
	fmt.Println(y)

	var z float64 = float64(y)
	fmt.Println(z)

	var a float64 = 3.9
	var b int32 = int32(a)
	fmt.Println(b)

	var thisString string = "100"
	thisStringInt, _ := strconv.Atoi(thisString)
	fmt.Println(thisStringInt)

	thisInt := -100
	thisIntString := strconv.Itoa(thisInt)
	fmt.Println(thisIntString)

	thisBoolString, _ := strconv.ParseBool("true")
	fmt.Println(thisBoolString)
	thisStringBool := strconv.FormatBool(false)
	fmt.Println(thisStringBool)

	thisFloatString, _ := strconv.ParseFloat("3.1415", 64)
	fmt.Println(thisFloatString)
	thisStringFloat := strconv.FormatFloat(3.1415, 'E', -1, 64)
	fmt.Println(thisStringFloat)
}
