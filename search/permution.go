package search

import (
	"fmt"
)

func Permutation(str string) []string {
	// write code here
	res := []string{}
	resMark := map[string]struct{}{}
	if str == "" {
		return []string{""}
	}
	if len(str) == 1 {
		return []string{str}
	}
	// mark all index
	mark := map[int]bool{}
	for i := range str {
		mark[i] = false
	}
	tmp := ""
	var bfs func(map[int]bool)
	bfs = func(m map[int]bool) {
		for i := 0; i < len(str); i++ {
			//	check mark and add to tmp
			visited := m[i]
			if visited {
				continue
			} else {
				tmp = tmp + fmt.Sprintf("%c", str[i])
				m[i] = true
			}
			if len(tmp) == len(str) {
				if _, ok := resMark[tmp]; !ok {
					res = append(res, tmp)
					resMark[tmp] = struct{}{}
				}
			}
			bfs(m)
			// backtrack
			lenTmp := len(tmp)
			m[i] = false
			tmp = tmp[:lenTmp-1]
		}
	}
	bfs(mark)
	return res
}
