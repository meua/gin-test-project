package string_

import "testing"

func TestStringRune(t *testing.T) {
	s := "中华人民共和国"
	for _, v := range s {
		t.Logf("%[1]c %[1]d", v)
	}
}
