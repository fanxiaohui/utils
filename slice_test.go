package common

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAppendStr(t *testing.T) {
	Convey("Append a string to a slice with no duplicates", t, func() {
		s := []string{"a"}

		Convey("Append a string that does not exist in slice", func() {
			s = AppendStr(s, "b")
			So(len(s), ShouldEqual, 2)
		})

		Convey("Append a string that does exist in slice", func() {
			s = AppendStr(s, "a")
			So(len(s), ShouldEqual, 1)
		})
	})
}

func TestAppendUint(t *testing.T) {
	Convey("appends uint to slice with no duplicates", t, func() {
		s := []uint{1}

		Convey("Append a uint that does not exist in slice", func() {
			s = AppendUint(s, 2)
			So(len(s), ShouldEqual, 2)
		})

		Convey("Append a uint that does exist in slice", func() {
			s = AppendUint(s, 1)
			So(len(s), ShouldEqual, 1)
		})
	})
}

func TestAppendUint16(t *testing.T) {
	Convey("appends uint16 to slice with no duplicates", t, func() {
		s := []uint16{1}

		Convey("Append a uint16 that does not exist in slice", func() {
			s = AppendUint16(s, 2)
			So(len(s), ShouldEqual, 2)
		})

		Convey("Append a uint16 that does exist in slice", func() {
			s = AppendUint16(s, 1)
			So(len(s), ShouldEqual, 1)
		})
	})
}

func TestDeleteFromSliceUint16(t *testing.T) {
	Convey("从uint16的切片中删除第一个指定元素", t, func() {
		s := []uint16{1, 2, 2, 3}

		Convey("从uint16的切片中删除第一个指定元素, 无指定元素值", func() {
			s = DeleteFromSliceUint16(s, 4)
			So(len(s), ShouldEqual, 4)
		})

		Convey("从uint16的切片中删除第一个指定元素, 有指定元素值", func() {
			s = DeleteFromSliceUint16(s, 2)
			So(len(s), ShouldEqual, 3)
		})

		Convey("从uint16的切片中删除第一个指定元素, 切片是个nil", func() {
			s = DeleteFromSliceUint16(nil, 2)
			So(s, ShouldBeNil)
		})

	})
}

func TestDeleteFromSliceUint16All(t *testing.T) {
	Convey("从uint16的切片中删除所有指定元素", t, func() {
		s := []uint16{1, 2, 2, 3}

		Convey("从uint16的切片中删除所有指定元素, 无指定元素值", func() {
			s = DeleteFromSliceUint16All(s, 4)
			So(len(s), ShouldEqual, 4)
		})

		Convey("从uint16的切片中删除所有指定元素, 有指定元素值", func() {
			s = DeleteFromSliceUint16All(s, 2)
			So(len(s), ShouldEqual, 2)
		})

		Convey("从uint16的切片中删除所有指定元素, 切片是个nil", func() {
			s = DeleteFromSliceUint16All(nil, 2)
			So(s, ShouldBeNil)
		})

	})
}

func TestDeleteFromSliceUint(t *testing.T) {
	Convey("从uint的切片中删除第一个指定元素", t, func() {
		s := []uint{1, 2, 2, 3}

		Convey("从uint的切片中删除第一个指定元素, 无指定元素值", func() {
			s = DeleteFromSliceUint(s, 4)
			So(len(s), ShouldEqual, 4)
		})

		Convey("从uint的切片中删除第一个指定元素, 有指定元素值", func() {
			s = DeleteFromSliceUint(s, 2)
			So(len(s), ShouldEqual, 3)
		})

		Convey("从uint的切片中删除第一个指定元素, 切片是个nil", func() {
			s = DeleteFromSliceUint(nil, 2)
			So(s, ShouldBeNil)
		})

	})
}

func TestDeleteFromSliceUintAll(t *testing.T) {
	Convey("从uint的切片中删除所有指定元素", t, func() {
		s := []uint{1, 2, 2, 3}

		Convey("从uint的切片中删除所有指定元素, 无指定元素值", func() {
			s = DeleteFromSliceUintAll(s, 4)
			So(len(s), ShouldEqual, 4)
		})

		Convey("从uint的切片中删除所有指定元素, 有指定元素值", func() {
			s = DeleteFromSliceUintAll(s, 2)
			So(len(s), ShouldEqual, 2)
		})

		Convey("从uint的切片中删除所有指定元素, 切片是个nil", func() {
			s = DeleteFromSliceUintAll(nil, 2)
			So(s, ShouldBeNil)
		})

	})
}
func TestCompareSliceStr(t *testing.T) {
	Convey("Compares two 'string' type slices with elements and order", t, func() {
		Convey("Compare two slices that do have same elements and order", func() {
			So(CompareSliceStr(
				[]string{"1", "2", "3"}, []string{"1", "2", "3"}), ShouldBeTrue)
		})

		Convey("Compare two slices that do have same elements but does not have same order", func() {
			So(!CompareSliceStr(
				[]string{"2", "1", "3"}, []string{"1", "2", "3"}), ShouldBeTrue)
		})

		Convey("Compare two slices that have different number of elements", func() {
			So(!CompareSliceStr(
				[]string{"2", "1"}, []string{"1", "2", "3"}), ShouldBeTrue)
		})
	})
}

func TestCompareSliceStrU(t *testing.T) {
	Convey("Compare two 'string' type slices with elements and ignore the order", t, func() {
		Convey("Compare two slices that do have same elements and order", func() {
			So(CompareSliceStrU(
				[]string{"1", "2", "3"}, []string{"1", "2", "3"}), ShouldBeTrue)
		})

		Convey("Compare two slices that do have same elements but does not have same order", func() {
			So(CompareSliceStrU(
				[]string{"2", "1", "3"}, []string{"1", "2", "3"}), ShouldBeTrue)
		})

		Convey("Compare two slices that do have different elements but has same count", func() {
			So(CompareSliceStrU(
				[]string{"2", "1", "4"}, []string{"1", "2", "3"}), ShouldBeFalse)
		})

		Convey("Compare two slices that have different number of elements", func() {
			So(!CompareSliceStrU(
				[]string{"2", "1"}, []string{"1", "2", "3"}), ShouldBeTrue)
		})

	})
}

func TestSliceContainsStr(t *testing.T) {
	Convey("字符串slice是否含有指定字符串", t, func() {
		s := []string{"A", "b", "c"}

		Convey("字符串slice含有指定字符串", func() {
			So(IsSliceContainsStr(
				s, "a"), ShouldBeTrue)
		})

		Convey("字符串slice不含有指定字符串", func() {
			So(IsSliceContainsStr(
				s, "f"), ShouldBeFalse)
		})
	})
}

func TestSliceContainsInt64(t *testing.T) {
	Convey("int64 slice是否含有指定值", t, func() {
		v := []int64{1, 2, 3, 4, 5}

		Convey("int64 slice含有指定值", func() {
			So(IsSliceContainsInt64(
				v, 2), ShouldBeTrue)
		})

		Convey("int64 slice不含有指定值", func() {
			So(IsSliceContainsInt64(
				v, 9), ShouldBeFalse)
		})
	})
}
func TestSliceContainsUint(t *testing.T) {
	Convey("uint slice是否含有指定值", t, func() {
		v := []uint{1, 2, 3, 4, 5}

		Convey("uint slice含有指定值", func() {
			So(IsSliceContainsUint(
				v, 2), ShouldBeTrue)
		})

		Convey("uint slice不含有指定值", func() {
			So(IsSliceContainsUint(
				v, 9), ShouldBeFalse)
		})
	})
}

func TestSliceContainsUint16(t *testing.T) {
	Convey("uint16 slice是否含有指定值", t, func() {
		v := []uint16{1, 2, 3, 4, 5}

		Convey("uint16 slice含有指定值", func() {
			So(IsSliceContainsUint16(
				v, 2), ShouldBeTrue)
		})

		Convey("uint16 slice不含有指定值", func() {
			So(IsSliceContainsUint16(
				v, 9), ShouldBeFalse)
		})
	})
}
