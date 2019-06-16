package utils

import "testing"

func TestFormatBaseTypes(t *testing.T) {
	type args struct {
		val  interface{}
		args []int
	}
	tests := []struct {
		name  string
		args  args
		wantS string
	}{
		{"bool", args{true, []int{}}, "true"},
		{"float32", args{float32(100.2), []int{}}, "100.2"},
		{"float64", args{float64(100.2), []int{}}, "100.2"},
		{"int", args{int(8), []int{}}, "8"},
		{"int8", args{int8(-3), []int{}}, "-3"},
		{"int16", args{int16(-3000), []int{}}, "-3000"},
		{"int32", args{int32(50000), []int{}}, "50000"},
		{"int64", args{int64(111111), []int{}}, "111111"},
		{"uint", args{uint(8), []int{}}, "8"},
		{"uint8", args{uint8(3), []int{}}, "3"},
		{"uint16", args{uint16(3000), []int{}}, "3000"},
		{"uint32", args{uint32(50000), []int{}}, "50000"},
		{"uint64", args{uint64(111111), []int{}}, "111111"},
		{"string", args{"hello", []int{}}, "hello"},
		{"[]byte", args{[]byte{'1', '2', '3'}, []int{}}, "123"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotS := FormatBaseTypes(tt.args.val, tt.args.args...); gotS != tt.wantS {
				t.Errorf("FormatBaseTypes() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}
