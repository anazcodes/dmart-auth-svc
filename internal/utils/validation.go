package utils

import (
	"context"
	"fmt"
)

func HasError(ctx context.Context, err error) bool {

	return err != nil
}

func HighlightError(str string) {
	fmt.Printf("\n\nerror:\n %s \n\n", str)
}
