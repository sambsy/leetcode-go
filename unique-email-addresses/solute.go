package solute

import (
	"bytes"
	"strings"
)

func numUniqueEmails(emails []string) int {
	result := make(map[string]bool)
	buf := bytes.Buffer{}
	for _, v := range emails {
		buf.Reset()
		fields := strings.Split(v, "@")
		if len(fields) != 2 {
			continue
		}

		for _, c := range []byte(fields[0]) {
			if c == '.' {
				continue
			}
			if c == '+' {
				break
			}
			buf.WriteByte(c)
		}
		buf.WriteByte('@')
		buf.WriteString(fields[1])

		result[buf.String()] = true
	}
	return len(result)
}
