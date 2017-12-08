package util

import (
	"fmt"
	"os"
)

func MaybeExitWithError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
