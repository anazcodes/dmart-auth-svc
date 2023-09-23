package util

import (
	"context"
	"fmt"
)

func HasError(ctx context.Context, err error) bool {
	if err != nil {
		Logger("has error:", err)

	}
	return err != nil
}

func Logger(any ...any) {
	fmt.Println("\n" + fmt.Sprint(any...) + "\n")
}
