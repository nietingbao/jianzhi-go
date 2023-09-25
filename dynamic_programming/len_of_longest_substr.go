package dynamicprogramming

import "fmt"

func lengthOfLongestSubstring(s string) int {
	// write code here
	// l,r 表示加入i之后最长的字符串的边界
	// map[string][int]记录经历过的字符的位置
	// 如果是新字符，更新r，对比max
	// 如果是旧字符，从map中找l，更新l为index++
	if s == "" {
		return 0
	}
	m := map[string]int{}
	max := 1
	l, r := 0, 1
	for i := 0; i < len(s); i++ {
		c := fmt.Sprintf("%c", s[i])
		if i == 0 {
			m[c] = 0
			continue
		}
		index, ok := m[c]
		if ok {
			// 必须要在范围内才能更新
			if index >= l {
				l = index + 1
			}
		}
		// 更新当前字符串的位置
		m[c] = i
		r = i
		if r-l+1 > max {
			max = r - l + 1
		}
	}
	return max
}
