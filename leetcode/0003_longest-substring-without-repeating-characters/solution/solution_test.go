package solution

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{"abcabcbb"}, 3},
		{"Example 2", args{"bbbbb"}, 1},
		{"Example 3", args{"pwwkew"}, 3},
		{"Example 4", args{"au"}, 2},
		{"Example 5", args{"aa"}, 1},
		{"Edge 0", args{""}, 0},
		{"Edge 1", args{" "}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})

		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring1(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring1() = %v, want %v", got, tt.want)
			}
		})

		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring2(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring2() = %v, want %v", got, tt.want)
			}
		})
	}
}
