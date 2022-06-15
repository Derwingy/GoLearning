package go_grammer

import (
	"log"
	"testing"
)

func TestNameStyle(t *testing.T) {
	val := NameStyle(32)
	if val == 0 {
		log.Printf("value %g is wrong", val)
	}
}
