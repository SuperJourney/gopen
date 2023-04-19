package common

import (
	"fmt"
)

func Info(format string, a ...interface{}) {
	fmt.Println("log begin----")
	fmt.Printf("\x1b[31m"+format+"\x1b[0m\n", a...)
	fmt.Println("-----end----")
}
