package variable

import (
	"fmt"
	"reflect"
)

func BoolFunc() {
	fmt.Println("")
	fmt.Println(".:: Boolean Data Type ::.")
	fmt.Println("")

	var benar bool
	benar = true

	var salah bool
	salah = false

	fmt.Println(benar, "=>", reflect.TypeOf(benar))
	fmt.Println(salah, "=>", reflect.TypeOf(salah))
}
