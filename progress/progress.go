package progress

import (
	"fmt"
	"os"
)

const clearLineSequence = "\x1b[1G\x1b[2K"

func PrintUpdate(update string, args ...interface{}) {
	Clear()
	fmt.Fprintf(os.Stderr, update, args...)
}

func Clear() {
	fmt.Fprintf(os.Stderr, clearLineSequence)
}