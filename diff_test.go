// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package diff

import (
	"strings"
	"testing"
)

var formatTests = []struct {
	text1 string
	text2 string
	diff  string
}{
	{"a b c", "a b d e f", "a b -c +d +e +f"},
	{"", "a b c", "+a +b +c"},
	{"a b c", "", "-a -b -c"},
	{"a b c", "d e f", "-a -b -c +d +e +f"},
	{"a b c d e f", "a b d e f", "a b -c d e f"},
	{"a b c e f", "a b c d e f", "a b c +d e f"},
}

func TestFormat(t *testing.T) {
	for _, tt := range formatTests {
		// Turn spaces into \n.
		text1 := strings.ReplaceAll(tt.text1, " ", "\n")
		if text1 != "" {
			text1 += "\n"
		}
		text2 := strings.ReplaceAll(tt.text2, " ", "\n")
		if text2 != "" {
			text2 += "\n"
		}
		out := Format(text1, text2)
		// Cut final \n, cut spaces, turn remaining \n into spaces.
		out = strings.ReplaceAll(strings.ReplaceAll(strings.TrimSuffix(out, "\n"), " ", ""), "\n", " ")
		if out != tt.diff {
			t.Errorf("diff(%q, %q) = %q, want %q", text1, text2, out, tt.diff)
		}
	}
}
