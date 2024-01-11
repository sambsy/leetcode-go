package solute

import (
	"log"
	"testing"
)

func TestSolute(t *testing.T) {
	cases := []struct {
		arg  string
		want int
	}{
		{
			arg:  "bbbbbb",
			want: 1,
		},
		{
			arg:  "pwwkew",
			want: 3,
		},
		{
			arg:  "abcabcbb",
			want: 3,
		},
	}

	for _, c := range cases {
		got := lengthOfLongestSubstring(c.arg)
		log.Println(got)
		if got != c.want {
			t.Errorf("lengthOfLongestSubstring(%s) == %d, want %d", c.arg, got, c.want)
		}
	}
}
