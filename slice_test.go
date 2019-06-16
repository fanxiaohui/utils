package utils

import (
	"reflect"
	"testing"
)

func TestAppendStr(t *testing.T) {
	type args struct {
		strs []string
		str  string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Append a string that does not exist in slice",
			args{[]string{"a"}, "b"}, []string{"a", "b"}},
		{"Append a string that does exist in slice",
			args{[]string{"a"}, "a"}, []string{"a"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AppendStr(tt.args.strs, tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppendUint(t *testing.T) {
	type args struct {
		s []uint
		e uint
	}
	tests := []struct {
		name string
		args args
		want []uint
	}{
		{"Append a uint that does not exist in slice",
			args{[]uint{1}, 2}, []uint{1, 2}},
		{"Append a uint that does exist in slice",
			args{[]uint{1}, 1}, []uint{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AppendUint(tt.args.s, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendUint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppendUint16(t *testing.T) {
	type args struct {
		s []uint16
		e uint16
	}
	tests := []struct {
		name string
		args args
		want []uint16
	}{
		{"Append a uint16 that does not exist in slice",
			args{[]uint16{1}, 2}, []uint16{1, 2}},
		{"Append a uint16 that does exist in slice",
			args{[]uint16{1}, 1}, []uint16{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AppendUint16(tt.args.s, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppendInt64(t *testing.T) {
	type args struct {
		s []int64
		e int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{"Append a int64 that does not exist in slice",
			args{[]int64{1, 2}, int64(3)}, []int64{1, 2, 3}},
		{"Append a int64 that does exist in slice",
			args{[]int64{1, 2}, 2}, []int64{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AppendInt64(tt.args.s, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteFromSliceStr(t *testing.T) {
	type args struct {
		s []string
		e string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"从string的切片中删除第一个指定元素, 无指定元素值",
			args{[]string{"a", "b", "b", "c"}, "e"}, []string{"a", "b", "b", "c"}},
		{"从string的切片中删除第一个指定元素, 有指定元素值",
			args{[]string{"a", "b", "b", "c"}, "b"}, []string{"a", "b", "c"}},
		{"从string的切片中删除第一个指定元素, 切片是个nil",
			args{nil, "e"}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteFromSliceStr(tt.args.s, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteFromSliceStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteFromSliceStrAll(t *testing.T) {
	type args struct {
		s []string
		e string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"从string的切片中删除所有指定元素, 无指定元素值",
			args{[]string{"a", "b", "b", "c"}, "e"}, []string{"a", "b", "b", "c"}},
		{"从string的切片中删除所有指定元素, 有指定元素值",
			args{[]string{"a", "b", "b", "c"}, "b"}, []string{"a", "c"}},
		{"从string的切片中删除所有指定元素, 切片是个nil",
			args{nil, "e"}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteFromSliceStrAll(tt.args.s, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteFromSliceStrAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteFromSliceUint16(t *testing.T) {
	type args struct {
		s []uint16
		e uint16
	}
	tests := []struct {
		name string
		args args
		want []uint16
	}{
		{"从uint16的切片中删除第一个指定元素, 无指定元素值",
			args{[]uint16{1, 2, 2, 3}, 4}, []uint16{1, 2, 2, 3}},
		{"从uint16的切片中删除第一个指定元素, 有指定元素值",
			args{[]uint16{1, 2, 2, 3}, 2}, []uint16{1, 2, 3}},
		{"从uint16的切片中删除第一个指定元素, 切片是个nil",
			args{nil, 2}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteFromSliceUint16(tt.args.s, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteFromSliceUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteFromSliceUint16All(t *testing.T) {
	type args struct {
		s []uint16
		e uint16
	}
	tests := []struct {
		name string
		args args
		want []uint16
	}{
		{"从uint16的切片中删除所有指定元素, 无指定元素值",
			args{[]uint16{1, 2, 2, 3}, 4}, []uint16{1, 2, 2, 3}},
		{"从uint16的切片中删除所有指定元素, 有指定元素值",
			args{[]uint16{1, 2, 2, 3}, 2}, []uint16{1, 3}},
		{"从uint16的切片中删除所有指定元素, 切片是个nil",
			args{nil, 2}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteFromSliceUint16All(tt.args.s, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteFromSliceUint16All() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteFromSliceUint(t *testing.T) {
	type args struct {
		s []uint
		e uint
	}
	tests := []struct {
		name string
		args args
		want []uint
	}{
		{"从uint的切片中删除第一个指定元素, 无指定元素值",
			args{[]uint{1, 2, 2, 3}, 4}, []uint{1, 2, 2, 3}},
		{"从uint的切片中删除第一个指定元素, 有指定元素值",
			args{[]uint{1, 2, 2, 3}, 2}, []uint{1, 2, 3}},
		{"从uint的切片中删除第一个指定元素, 切片是个nil",
			args{nil, 2}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteFromSliceUint(tt.args.s, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteFromSliceUint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteFromSliceUintAll(t *testing.T) {
	type args struct {
		s []uint
		e uint
	}
	tests := []struct {
		name string
		args args
		want []uint
	}{
		{"从uint的切片中删除所有指定元素, 无指定元素值",
			args{[]uint{1, 2, 2, 3}, 4}, []uint{1, 2, 2, 3}},
		{"从uint的切片中删除所有指定元素, 有指定元素值",
			args{[]uint{1, 2, 2, 3}, 2}, []uint{1, 3}},
		{"从uint的切片中删除所有指定元素, 切片是个nil",
			args{nil, 2}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteFromSliceUintAll(tt.args.s, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteFromSliceUintAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteFromSliceInt64(t *testing.T) {
	type args struct {
		s []int64
		e int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{"从int64的切片中删除第一个指定元素, 无指定元素值",
			args{[]int64{1, 2, 2, 3}, 4}, []int64{1, 2, 2, 3}},
		{"从int64的切片中删除第一个指定元素, 有指定元素值",
			args{[]int64{1, 2, 2, 3}, 2}, []int64{1, 2, 3}},
		{"从int64的切片中删除第一个指定元素, 切片是个nil",
			args{nil, 2}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteFromSliceInt64(tt.args.s, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteFromSliceInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteFromSliceInt64All(t *testing.T) {
	type args struct {
		s []int64
		e int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{"从int64的切片中删除所有指定元素, 无指定元素值",
			args{[]int64{1, 2, 2, 3}, 4}, []int64{1, 2, 2, 3}},
		{"从int64的切片中删除所有指定元素, 有指定元素值",
			args{[]int64{1, 2, 2, 3}, 2}, []int64{1, 3}},
		{"从int64的切片中删除所有指定元素, 切片是个nil",
			args{nil, 2}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteFromSliceInt64All(tt.args.s, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteFromSliceInt64All() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompareSliceStr(t *testing.T) {
	type args struct {
		s1 []string
		s2 []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Compare two slices that do have same elements and order",
			args{[]string{"1", "2", "3"}, []string{"1", "2", "3"}}, true},
		{"Compare two slices that do have same elements but does not have same order",
			args{[]string{"2", "1", "3"}, []string{"1", "2", "3"}}, false},
		{"Compare two slices that have different number of elements",
			args{[]string{"1", "2"}, []string{"1", "2", "3"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareSliceStr(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("CompareSliceStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompareSliceStrU(t *testing.T) {
	type args struct {
		s1 []string
		s2 []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Compare two slices that do have same elements and order",
			args{[]string{"1", "2", "3"}, []string{"1", "2", "3"}}, true},
		{"Compare two slices that do have same elements but does not have same order",
			args{[]string{"2", "1", "3"}, []string{"1", "2", "3"}}, true},
		{"Compare two slices that do have different elements but has same count",
			args{[]string{"2", "1", "4"}, []string{"1", "2", "3"}}, false},
		{"Compare two slices that have different number of elements",
			args{[]string{"1", "2"}, []string{"1", "2", "3"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareSliceStrU(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("CompareSliceStrU() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSliceContainsStr(t *testing.T) {
	type args struct {
		sl  []string
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"字符串slice含有指定字符串,大小写敏感,一样", args{[]string{"A", "b", "c"}, "A"}, true},
		{"字符串slice不含有指定字符串,大小写敏感,因为小写的", args{[]string{"A", "b", "c"}, "a"}, false},
		{"字符串slice不含有指定字符串,大小写敏感", args{[]string{"A", "b", "c"}, "d"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSliceContainsStr(tt.args.sl, tt.args.str); got != tt.want {
				t.Errorf("IsSliceContainsStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSliceContainsStrNocase(t *testing.T) {
	type args struct {
		sl  []string
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"字符串slice含有指定字符串,忽略大小写,一样", args{[]string{"A", "b", "c"}, "A"}, true},
		{"字符串slice不含有指定字符串,忽略大小写,因为小写的", args{[]string{"A", "b", "c"}, "a"}, true},
		{"字符串slice不含有指定字符串,忽略大小写", args{[]string{"A", "b", "c"}, "d"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSliceContainsStrNocase(tt.args.sl, tt.args.str); got != tt.want {
				t.Errorf("IsSliceContainsStrNocase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSliceContainsInt64(t *testing.T) {
	type args struct {
		sl []int64
		i  int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"int64 slice含有指定值", args{[]int64{1, 2}, 2}, true},
		{"int64 slice不含有指定值", args{[]int64{1, 2}, 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSliceContainsInt64(tt.args.sl, tt.args.i); got != tt.want {
				t.Errorf("IsSliceContainsInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSliceContainsUint(t *testing.T) {
	type args struct {
		sl []uint
		i  uint
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"uint slice含有指定值", args{[]uint{1, 2}, 2}, true},
		{"uint slice不含有指定值", args{[]uint{1, 2}, 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSliceContainsUint(tt.args.sl, tt.args.i); got != tt.want {
				t.Errorf("IsSliceContainsUint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSliceContainsUint16(t *testing.T) {
	type args struct {
		sl []uint16
		i  uint16
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"uint16 slice含有指定值", args{[]uint16{1, 2}, 2}, true},
		{"uint16 slice不含有指定值", args{[]uint16{1, 2}, 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSliceContainsUint16(tt.args.sl, tt.args.i); got != tt.want {
				t.Errorf("IsSliceContainsUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseSliceBytes(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"Reverse bytes", args{[]byte{1, 2, 3, 4, 5}}, []byte{5, 4, 3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseSliceBytes(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseSliceBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Reverse String", args{"hello"}, "olleh"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseString(tt.args.s); got != tt.want {
				t.Errorf("ReverseString() = %v, want %v", got, tt.want)
			}
		})
	}
}
