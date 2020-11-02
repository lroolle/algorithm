package solution

import "testing"

func Test_solveSudoku(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
	}{
		{"Example 1", args{}, },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solveSudoku(tt.args.board)
		})
	}
}
