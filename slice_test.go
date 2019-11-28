package utils

import (
	"reflect"
	"testing"
)

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
