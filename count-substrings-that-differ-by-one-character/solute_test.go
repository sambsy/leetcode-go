package solute

import "testing"

func Test_countSubstrings(t *testing.T) {
	cases := []struct {
		s    string
		t    string
		want int
	}{
		{
			s:    "aba",
			t:    "baba",
			want: 6,
		}, {
			s:    "ab",
			t:    "bb",
			want: 3,
		}, {
			s:    "a",
			t:    "a",
			want: 0,
		}, {
			s:    "abe",
			t:    "bbc",
			want: 10,
		},
	}

	for _, cc := range cases {
		result := countSubstrings(cc.s, cc.t)
		if result != cc.want {
			t.Errorf("s[%v],t[%v],want[%v],got[%v]", cc.s, cc.t, cc.want, result)
		}
	}
}
