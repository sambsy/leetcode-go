package solute

import "testing"

func TestSolute(t *testing.T) {
	cases := []struct {
		arg1 string
		arg2 string
		want bool
	}{
		{
			arg1: "ab",
			arg2: "eidbaooo",
			want: true,
		},
		{
			arg1: "ab",
			arg2: "eidboaoo",
			want: false,
		},
	}

	for _, c := range cases {
		got := checkInclusion(c.arg1, c.arg2)
		if got != c.want {
			t.Errorf("checkInclusion(%s, %s) == %t, want %t", c.arg1, c.arg2, got, c.want)
		}
	}
}
