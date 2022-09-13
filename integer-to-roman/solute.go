package solute

var mm = map[int]string{
	1:    "I",
	4:    "IV",
	5:    "V",
	9:    "IX",
	10:   "X",
	40:   "XL",
	50:   "L",
	90:   "XC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "CM",
	1000: "M",
}
var mmm = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

// [题目描述](https://leetcode.cn/problems/integer-to-roman/)
func intToRoman(num int) string {
	result := ""
	for num > 0 {
		for _, v := range mmm {
			if num >= v {
				result += mm[v]
				num -= v
				break
			}
		}
	}
	return result
}
