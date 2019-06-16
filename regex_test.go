package utils

import (
	"testing"
)

func TestIsChinese(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"纯中文", args{"你好朋友"}, true},
		{"中文加标点", args{"你好,朋友"}, false},
		{"中英文混合", args{"好的OK"}, false},
		{"纯英文", args{"hello"}, false},
		{"英文加标点", args{"hello，Jhon"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsChinese(tt.args.s); got != tt.want {
				t.Errorf("IsChinese() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsEnglish(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"纯中文", args{"你好朋友"}, false},
		{"中文加标点", args{"你好,朋友"}, false},
		{"中英文混合", args{"好的OK"}, false},
		{"纯英文", args{"hello"}, true},
		{"英文加标点", args{"hello，Jhon"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEnglish(tt.args.s); got != tt.want {
				t.Errorf("IsEnglish() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPhone(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"10086??", args{`10086`}, false},
		{"13开头电话号码", args{`13075993418`}, true},
		{"14开头电话号码", args{`14205084501`}, true},
		{"15开头电话号码", args{`15205084501`}, true},
		{"18开头电话号码", args{`18205084501`}, true},
		{"其它不合规范电话", args{`16777777777`}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPhone(tt.args.s); got != tt.want {
				t.Errorf("IsPhone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsIDCard(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"15位身份证号", args{`123456789012345`}, true},
		{"18位身份证号", args{`123456789012345678`}, true},
		{"18位身份证号带x", args{`12345678901234567X`}, true},
		{"位数不够的身份证号", args{`123456785`}, false},
		{"身份证号不符合规范", args{`12345678901234567x`}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsIDCard(tt.args.s); got != tt.want {
				t.Errorf("IsIDCard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsDigit(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"自然数", args{`123456789012345`}, true},
		{"自然数0", args{`0`}, true},
		{"负整数", args{`-123456789012345`}, false},
		{"浮点数", args{`123456.789012345`}, false},
		{"数字带英文", args{`123456xx89012345`}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDigit(tt.args.s); got != tt.want {
				t.Errorf("IsDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"email", args{`test@example.com`}, true},
		{"email", args{`single-character@b.org`}, true},
		{"email", args{`uncommon_address@test.museum`}, true},
		{"email", args{`local@sld.UPPER`}, true},
		{"email", args{`@missing.org`}, false},
		{"email", args{`missing@.com`}, false},
		{"email", args{`missing@qq.`}, false},
		{"email", args{`wrong-ip@127.1.1.1.26`}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmail(tt.args.email); got != tt.want {
				t.Errorf("IsEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsEmailRFC(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"rfc email", args{`test@example.com`}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmailRFC(tt.args.email); got != tt.want {
				t.Errorf("IsEmailRFC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsURL(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"http with www", args{"http://www.example.com"}, true},
		{"http without www", args{"http://example.com"}, true},
		{"https with www", args{"https://www.example.com"}, true},
		{"https without www", args{"https://example.com"}, true},
		{"ftp with www", args{"ftp://www.example.com"}, true},
		{"ftp without www", args{"ftp://example.com"}, true},
		{"http 带参数", args{"http://example.com?user=test&password=test"}, true},
		{"http 带特别参数", args{"http://example.com?user=test#login"}, true},
		{"http 头不正确", args{"htp://example.com"}, false},
		{"http 头缺冒号", args{"http//example.com"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsURL(tt.args.url); got != tt.want {
				t.Errorf("IsURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIsEmail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsEmail("test@example.com")
	}
}

func BenchmarkIsUrl(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsEmail("http://example.com")
	}
}
