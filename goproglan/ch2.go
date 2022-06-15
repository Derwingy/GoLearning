package goproglan

import (
	"flag"
	"fmt"
	"strings"
)

func FlagUse() {
	var n = flag.Bool("n", false, "de")
	var sep = flag.String("s", " ", "ds")
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
