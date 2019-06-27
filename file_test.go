// Copyright 2013 com authors
// authors: Unknwon
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package utils

import (
	"testing"
)

func TestHumaneFileSize(t *testing.T) {
	type args struct {
		s uint64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"B", args{5}, "5B"},
		{"KB", args{100 * 1024}, "100KB"},
		{"MB", args{100 * 1024 * 1024}, "100MB"},
		{"GB", args{100 * 1024 * 1024 * 1024}, "100GB"},
		{"TB", args{100 * 1024 * 1024 * 1024 * 1024}, "100TB"},
		{"PB", args{100 * 1024 * 1024 * 1024 * 1024 * 1024}, "100PB"},
		{"EB", args{10 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024}, "10EB"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HumaneFileSize(tt.args.s); got != tt.want {
				t.Errorf("HumaneFileSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsFile(t *testing.T) {
	if !IsFile("file.go") {
		t.Errorf("IsFile:\n Expect => %v\n Got => %v\n", true, false)
	}

	if IsFile("testdata") {
		t.Errorf("IsFile:\n Expect => %v\n Got => %v\n", false, true)
	}

	if IsFile("files.go") {
		t.Errorf("IsFile:\n Expect => %v\n Got => %v\n", false, true)
	}
}

func TestIsExist(t *testing.T) {
	t.Run("Pass a file name that exists", func(t *testing.T) {
		if !IsExist("file.go") {
			t.Errorf("IsExist:\n Expect => %v\n Got => %v\n", true, false)
		}
	})

	t.Run("Pass a directory name that exists", func(t *testing.T) {
		if !IsExist("testdata") {
			t.Errorf("IsExist:\n Expect => %v\n Got => %v\n", true, false)
		}
	})

	t.Run("Pass a directory name that does not exist", func(t *testing.T) {
		if IsExist(".hg") {
			t.Errorf("IsExist:\n Expect => %v\n Got => %v\n", false, true)
		}
	})
}

func BenchmarkIsFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsFile("file.go")
	}
}

func BenchmarkIsExist(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsExist("file.go")
	}
}
