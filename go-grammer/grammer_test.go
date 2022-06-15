package go_grammer

import (
	"fmt"
	"testing"
)

// func TestLongestSubSeq(t *testing.T) {
// 	test_str := "wekndew"
// 	res := no_repeat_longest_sub_seq(test_str)
// 	if res != "knedw" {
// 		fmt.Printf("wrong")
// 	}
// }

func TestLongestSubStr(t *testing.T) {
	str1 := "abc"
	str2 := "dkabs"
	len := longest_sub_str(str1, str2, len(str1), len(str2))
	fmt.Println(len)
}
