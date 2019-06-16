package utils

import (
	"testing"
)

func TestBreakUint16(t *testing.T) {
	type args struct {
		v uint16
	}
	tests := []struct {
		name       string
		args       args
		wantLoByte byte
		wantHiByte byte
	}{
		{"uint16 break to byte", args{0x1234}, 0x34, 0x12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLoByte, gotHiByte := BreakUint16(tt.args.v)
			if gotLoByte != tt.wantLoByte {
				t.Errorf("BreakUint16() gotLoByte = %v, want %v", gotLoByte, tt.wantLoByte)
			}
			if gotHiByte != tt.wantHiByte {
				t.Errorf("BreakUint16() gotHiByte = %v, want %v", gotHiByte, tt.wantHiByte)
			}
		})
	}
}

func TestBreakUint32(t *testing.T) {
	type args struct {
		v uint32
	}
	tests := []struct {
		name      string
		args      args
		wantByte0 byte
		wantByte1 byte
		wantByte2 byte
		wantByte3 byte
	}{
		{"uint32 break to byte", args{0x12345678}, 0x78, 0x56, 0x34, 0x12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotByte0, gotByte1, gotByte2, gotByte3 := BreakUint32(tt.args.v)
			if gotByte0 != tt.wantByte0 {
				t.Errorf("BreakUint32() gotByte0 = %v, want %v", gotByte0, tt.wantByte0)
			}
			if gotByte1 != tt.wantByte1 {
				t.Errorf("BreakUint32() gotByte1 = %v, want %v", gotByte1, tt.wantByte1)
			}
			if gotByte2 != tt.wantByte2 {
				t.Errorf("BreakUint32() gotByte2 = %v, want %v", gotByte2, tt.wantByte2)
			}
			if gotByte3 != tt.wantByte3 {
				t.Errorf("BreakUint32() gotByte3 = %v, want %v", gotByte3, tt.wantByte3)
			}
		})
	}
}

func TestBuildUint16(t *testing.T) {
	type args struct {
		loByte byte
		Hibyte byte
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{"build byte to uint16", args{0x34, 0x12}, 0x1234},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildUint16(tt.args.loByte, tt.args.Hibyte); got != tt.want {
				t.Errorf("BuildUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildUint32(t *testing.T) {
	type args struct {
		Byte0 byte
		Byte1 byte
		Byte2 byte
		Byte3 byte
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{"build byte to uint32", args{0x78, 0x56, 0x34, 0x12}, 0x12345678},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildUint32(tt.args.Byte0, tt.args.Byte1, tt.args.Byte2, tt.args.Byte3); got != tt.want {
				t.Errorf("BuildUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}
