package godatastructure

import "fmt"

type IdInt int

type NameStr string

func funcBool(id int, str string, p []byte) bool {
	if id == 1 || len(p) == 0 {
		return true
	}
	return len(str) != 0
}

func (i *IdInt) Write(p []byte) (int, error) {
	*i += IdInt(len(p))
	return len(p), nil
}

func main() {
	buf := []byte("opssre")
	res := funcBool(10, "Young", buf)
	if res {
		fmt.Print("ok")
	}
}
