package go_grammer

import "strings"

func no_repeat_longest_sub_seq(s string) string {
	var res string
	if len(s) == 0 {
		res = ""
	}
	if len(s) == 1 {
		res = s
	}
	for idx := 0; idx < len(s)-1; idx++ {
		var re_str string
		re_ch := s[idx]
		re_str = string(re_ch)
		tmp_str := string(re_ch)
		for pointer_pos := idx + 1; pointer_pos < len(s); pointer_pos++ {
			ch := s[pointer_pos]
			if ch != s[idx] && !strings.Contains(re_str, string(ch)) {
				tmp_str += string(ch)
				if len(tmp_str) >= len(re_str) {
					re_str = tmp_str
				}
			} else {
				break
			}
		}
		if len(res) <= len(re_str) {
			res = re_str
		}
	}
	return res
}

func longest_sub_str(str1 string, str2 string, len_str1 int, len_str2 int) int {
	table := make([][]int, len_str1+1)
	for k := 0; k < len_str1+1; k++ {
		table[k] = make([]int, len_str2+1)
	}
	for i := 0; i < len_str1+1; i++ {
		for j := 0; j < len_str2+1; j++ {
			if i == 0 || j == 0 {
				table[i][j] = 0
			} else if str1[i-1] == str2[j-1] {
				table[i][j] = longest_sub_str(str1, str2, len_str1-1, len_str2-1) + 1
			} else {
				if longest_sub_str(str1, str2, len_str1-1, len_str2) > longest_sub_str(str1, str2, len_str1, len_str2-1) {
					table[i][j] = longest_sub_str(str1, str2, len_str1-1, len_str2)
				} else {
					table[i][j] = longest_sub_str(str1, str2, len_str1, len_str2-1)
				}
			}
		}
	}
	return table[len_str1][len_str2]
}
