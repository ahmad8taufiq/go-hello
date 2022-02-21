package variable

import (
	"fmt"
	// "reflect"
)

func StringFunc() {
	fmt.Println("")
	fmt.Println(".:: String Data Type ::.")
	fmt.Println("")

	var str string
	str = "Ahmad Taufiq"

	fmt.Println(len(str))
	fmt.Println(str[0])
	fmt.Println("Ahmad Taufiq"[7])
}
