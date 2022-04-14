package goproglan

import (
	"fmt"
)

type InterfaceLearn interface {
}

func TesModule() string {
	fmt.Println("module is ok")

	// ctx, cancel := context.WithTimeout(ctx, 100*Millisecond)
	// defer cancel()

	return "ok"
}
