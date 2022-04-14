package godatastructure

import "fmt"

type IdInt int

type NameStr string

func boolUsage(id int, str string, p []byte) bool {
	if id == 1 || len(p) == 0 {
		return true
	}
	return len(str) != 0
}

func (i *IdInt) Write(p []byte) (int, error) {
	*i += IdInt(len(p))
	return len(p), nil
}

func mapUsage() []string {
	mp := make(map[string]int)
	mp["year"] = 2022
	mp["month"] = 04
	mp["day"] = 13

	keys := make([]string, 0, len(mp))
	for k := range mp {
		keys = append(keys, k)
	}
	return keys
}

func output() {
	buf := []byte("sre")
	res := boolUsage(10, "Young", buf)
	if res {
		fmt.Print("ok")
	}
	mapRes := mapUsage()
	fmt.Print(mapRes)
}
