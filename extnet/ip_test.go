// 扩展net包, 提供ip到数值的转换

package extnet

import (
	"net"
	"reflect"
	"testing"
)

func TestIP2Numer(t *testing.T) {
	type args struct {
		p net.IP
	}
	tests := []struct {
		name    string
		args    args
		want    uint32
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IP2Numer(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("IP2Numer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IP2Numer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumer2IP(t *testing.T) {
	type args struct {
		l uint32
	}
	tests := []struct {
		name string
		args args
		want net.IP
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Numer2IP(tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Numer2IP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDot2Numer(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    uint32
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Dot2Numer(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Dot2Numer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Dot2Numer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumber2Dot(t *testing.T) {
	type args struct {
		l uint32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Number2Dot(tt.args.l); got != tt.want {
				t.Errorf("Number2Dot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMask2Dot(t *testing.T) {
	type args struct {
		mask net.IPMask
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mask2Dot(tt.args.mask); got != tt.want {
				t.Errorf("Mask2Dot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDot2Mask(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want net.IPMask
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Dot2Mask(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dot2Mask() = %v, want %v", got, tt.want)
			}
		})
	}
}
