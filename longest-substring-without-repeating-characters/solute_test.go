package solute

import "testing"

func TestSolute(t *testing.T) {
	cases := []struct {
		arg  string
		want int
	}{
		{
			arg:  "abcabcbb",
			want: 3,
		},
	}

	for _, c := range cases {
		got := lengthOfLongestSubstring(c.arg)
		if got != c.want {
			t.Errorf("lengthOfLongestSubstring(%s) == %d, want %d", c.arg, got, c.want)
		}
	}
}
