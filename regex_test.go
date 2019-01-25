package common

import (
    . "github.com/smartystreets/goconvey/convey"
    "testing"
)

func TestIsChinese(t *testing.T) {
    chinese := map[string]bool{
        `你好朋友`:      true,
        `好的呢`:       true,
        `不错`:        true,
        `好的OK`:      false,
        `hello 是你呀`: false,
        `好的,马上`:     false,
    }
    Convey("是否是中文", t, func() {
        for k, v := range chinese {
            So(IsChinese(k), ShouldEqual, v)
        }
    })
}

func TestIsEnglish(t *testing.T) {
    english := map[string]bool{
        `你好朋友`:        false,
        `好的呢,john`:    false,
        `ok,it is me`: false,
        `OK`:          true,
        `hellohow`:    true,
        `now`:         true,
    }
    Convey("是否是英文", t, func() {
        for k, v := range english {
            So(IsEnglish(k), ShouldEqual, v)
        }
    })
}

func TestIsPhone(t *testing.T) {
    phone := map[string]bool{
        `10086`:       false,
        `18075993418`: true,
        `15205084501`: true,
        `8434134`:     false,
        `45635624`:    false,
    }
    Convey("是否是电话号码", t, func() {
        for k, v := range phone {
            So(IsPhone(k), ShouldEqual, v)
        }
    })
}

func TestIsIDCard(t *testing.T) {
    phone := map[string]bool{
        `123456789012345`:    true,
        `12345678901234511X`: true,
        `12345678901234511x`: false,
        `45635624`:           false,
    }
    Convey("是否是身份证", t, func() {
        for k, v := range phone {
            So(IsIDCard(k), ShouldEqual, v)
        }
    })
}

func TestIsDigit(t *testing.T) {
    phone := map[string]bool{
        `123456789012345`:    true,
        `12345678901234511X`: false,
        `12345678901234511x`: false,
        `45635624`:           true,
    }
    Convey("是否是数字", t, func() {
        for k, v := range phone {
            So(IsDigit(k), ShouldEqual, v)
        }
    })
}

func TestIsEmail(t *testing.T) {
    emails := map[string]bool{
        `test@example.com`:             true,
        `single-character@b.org`:       true,
        `uncommon_address@test.museum`: true,
        `local@sld.UPPER`:              true,
        `@missing.org`:                 false,
        `missing@.com`:                 false,
        `missing@qq.`:                  false,
        `wrong-ip@127.1.1.1.26`:        false,
    }

    Convey("是否是email地址", t, func() {
        for k, v := range emails {
            So(IsEmail(k), ShouldEqual, v)
        }
    })

    Convey("是否是email地址", t, func() {
        So(IsEmailRFC(`test@example.com`), ShouldBeTrue)
    })
}

func TestIsUrl(t *testing.T) {
    urls := map[string]bool{
        "http://www.example.com":                     true,
        "http://example.com":                         true,
        "http://example.com?user=test&password=test": true,
        "http://example.com?user=test#login":         true,
        "ftp://example.com":                          true,
        "https://example.com":                        true,
        "htp://example.com":                          false,
        "http//example.com":                          false,
        "http://example":                             true,
    }
    Convey("是否是url地址", t, func() {
        for k, v := range urls {
            So(IsUrl(k), ShouldEqual, v)
        }
    })

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
