package util

import (
	"fmt"
)

func HasError(err error) bool {
	if err != nil {
		Logger("has error:", err)

	}
	return err != nil
}

func Logger(any ...any) {
	fmt.Println("\n" + fmt.Sprint(any...) + "\n")
}
