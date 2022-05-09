package utils

import (
	"fmt"
	"os"
)

func Dd(values ...any) {
	for _, value := range values {
		fmt.Println(value)
	}
	os.Exit(500)
}
