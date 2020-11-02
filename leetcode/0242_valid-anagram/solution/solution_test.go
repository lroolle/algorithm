package solution

import "testing"

func Test_isAnagram(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{"anagram", "nagaram"}, true},
		{"Example 2", args{"rat", "car"}, false},
		{"Example 3", args{"aacc", "ccac"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAnagram() = %v, want %v", got, tt.want)
			}

			if got := isAnagramBitMap(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}

var s = "aeivwpijhnfutngezpjplenrigsmwniribupuogkndgcbtwnszewkfjxkciivrkgfzcshgywwzsiynkvybiwwtuvjiexhvxjeehh"
var t = "aeivwpijhnfutngezpjplenrigsmwniribupuogkndgcbtwnszewkfjxkciivrkgfzcshgywwzsiynkvybiwwtuvjiexhvxjeehh"

func Benchmark_isAnagramBitMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		isAnagramBitMap(s, t)
	}
}

func Benchmark_isAnagramBitMapGoroutine(b *testing.B) {
	for n := 0; n < b.N; n++ {
		isAnagramBitMapGoroutine(s, t)
	}
}
