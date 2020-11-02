package solution

import "testing"

func Test_convert(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Example 1", args{s: "PAYPALISHIRING", numRows: 3}, "PAHNAPLSIIGYIR"},
		{"Example 2", args{s: "PAYPALISHIRING", numRows: 4}, "PINALSIGYAHRPI"},
		{"Example 3", args{s: "ABC", numRows: 1}, "ABC"},
		{"Unicode 1", args{s: "我是谁我在哪", numRows: 2}, "我谁在是我哪"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})

		t.Run(tt.name, func(t *testing.T) {
			if got := convert2(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("convert2() = %v, want %v", got, tt.want)
			}
		})
	}
}
