package backtrack

func ambiguousCoordinates(s string) []string {
	// 坐标 输入 包含(),输出也需要用()包住，然后x，y之间需要用, 连接
	// 坐标不能是1.0之类，即最后一位不能是0
	// 暴力求解，先找分割点，然后分别前后求集合，最后笛卡尔积
	if s == "()" || len(s) <= 3 {
		return []string{}
	}
	res := []string{}
	for i := 1; i < len(s)-1; i++ {
		x, xValid := parseNum(s[1 : i+1])
		y, yValid := parseNum(s[i+1 : len(s)-1])
		if xValid && yValid {
			for _, xstr := range x {
				for _, ystr := range y {
					res = append(res, "("+xstr+", "+ystr+")")
				}
			}
		}
	}
	return res

}

// res, valid
func parseNum(s string) ([]string, bool) {
	res := []string{}
	if s == "" {
		return res, false
	}
	if len(s) == 1 {
		return []string{s}, true
	}

	for i := 1; i <= len(s); i++ {
		firstPart, valid := checkValid(s[:i], 1)
		if !valid {
			// 整数部分只有可能是头部为0的二位数即以上，如果这样，那么后续的都不用分割，全部不行
			break
		}
		if i == len(s) {
			res = append(res, firstPart)
			break
		}
		secondPart, valid := checkValid(s[i:], 2)
		if !valid {
			continue
		}
		res = append(res, firstPart+"."+secondPart)
	}
	return res, len(res) != 0
}

func checkValid(s string, part int) (string, bool) {
	if len(s) == 0 {
		return "", false
	}
	res := ""
	if part == 1 {
		if s[0] == '0' && len(s) > 1 {
			return res, false
		}
	} else {
		if s[len(s)-1] == '0' {
			return res, false
		}
	}
	res = s
	return res, true
}
