package variable

import (
	"fmt"
	"reflect"
)

func NumFunc() {
	// Basic type
	var integer8 int8 = -128
	var integer16 int16 = -32768
	var integer32 int32 = -2147483648
	var integer64 int64 = -9223372036854775808

	fmt.Println(integer8, "=>", reflect.TypeOf(integer8))
	fmt.Println(integer16, "=>", reflect.TypeOf(integer16))
	fmt.Println(integer32, "=>", reflect.TypeOf(integer32))
	fmt.Println(integer64, "=>", reflect.TypeOf(integer64))

	fmt.Println("")

	var unsignedInt8 uint8 = 255
	var unsignedInt16 uint16 = 65535
	var unsignedInt32 uint32 = 4294967295
	var unsignedInt64 uint64 = 18446744073709551615

	fmt.Println(unsignedInt8, "=>", reflect.TypeOf(unsignedInt8))
	fmt.Println(unsignedInt16, "=>", reflect.TypeOf(unsignedInt16))
	fmt.Println(unsignedInt32, "=>", reflect.TypeOf(unsignedInt32))
	fmt.Println(unsignedInt64, "=>", reflect.TypeOf(unsignedInt64))

	fmt.Println("")

	var float1 float32 = 3.14
	var float2 float64 = 3.14

	fmt.Println(float1, "=>", reflect.TypeOf(float1))
	fmt.Println(float2, "=>", reflect.TypeOf(float2))

	fmt.Println("")

	// Aliast type
	var bite byte = 8
	var run rune = -32
	var integer int = -64
	var unsignedInteger uint = 64

	fmt.Println(bite, "=>", "byte /", reflect.TypeOf(bite))
	fmt.Println(run, "=>", "rune /", reflect.TypeOf(run))
	fmt.Println(integer, "=>", "int /", reflect.TypeOf(integer))
	fmt.Println(unsignedInteger, "=>", "uint /", reflect.TypeOf(unsignedInteger))
}
