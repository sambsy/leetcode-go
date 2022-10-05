package solute

import (
	"strconv"
	"strings"
)

// [题目描述](https://leetcode-cn.com/problems/subdomain-visit-count/)
func subdomainVisits(cpdomains []string) []string {
	result := make(map[string]int, len(cpdomains))

	for subdomain, count := range split(cpdomains) {
		fields := strings.Split(subdomain, ".")
		for i := 0; i < len(fields); i++ {
			item := strings.Join(fields[i:], ".")
			result[item] += count
		}

	}
	return join(result)

}

func split(cpdomains []string) map[string]int {
	result := make(map[string]int, len(cpdomains))

	for _, v := range cpdomains {
		fields := strings.Split(v, " ")
		if len(fields) != 2 {
			continue
		}

		count, _ := strconv.Atoi(fields[0])
		result[fields[1]] += count
	}

	return result
}

func join(counts map[string]int) []string {
	result := make([]string, 0, len(counts))

	sb := strings.Builder{}
	for subdomain, count := range counts {
		sb.Reset()
		sb.WriteString(strconv.Itoa(count))
		sb.WriteRune(' ')
		sb.WriteString(subdomain)
		result = append(result, sb.String())
	}

	return result
}
