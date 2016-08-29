package output

import (
	"fmt"
	"github.com/fatih/color"
	"os"
)

func PrintErrorAndExit(message string) {
	color.Set(color.FgRed)
	fmt.Println(message)
	color.Unset()
	os.Exit(1)
}
