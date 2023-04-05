package solute

import "testing"

func Test_isSubsequence(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want bool
	}{
		{"abc", "ahbgdc", true},
		{"axc", "ahbgdc", false},
	}
	for _, tt := range tests {
		got := isSubsequence(tt.s, tt.t)
		if got != tt.want {
			t.Errorf("s %v, t %v, got %v, want %v", tt.s, tt.t, got, tt.want)
		}
	}
}
