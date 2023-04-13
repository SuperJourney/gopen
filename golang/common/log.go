package common

import (
	"fmt"
)

func Info(format string, a ...any) {
	fmt.Printf(format, a...)
}
