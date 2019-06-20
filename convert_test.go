package utils

import (
	"reflect"
	"testing"
)

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

func TestStr2Bytes(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"string to bytes no copy", args{"hello world"}, []byte("hello world")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Str2Bytes(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Str2Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytes2Str(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"bytes to string no copy", args{[]byte{'h', 'e', 'l', 'l', 'o'}}, "hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bytes2Str(tt.args.b); got != tt.want {
				t.Errorf("Bytes2Str() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkStr2Bytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Str2Bytes("hello world")
	}
}

func BenchmarkStr2BytesCommon(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = []byte("hello world")
	}
}

func BenchmarkBytes2Str(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Bytes2Str([]byte{'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a'})
	}
}

func BenchmarkBytes2StrCommon(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = string([]byte{'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a'})
	}
}
