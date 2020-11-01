package string

import "testing"

func TestString(t *testing.T){
	var s string
	t.Log(s)

	//s[1] = '3'
	s = "hello"
	t.Log(len(s))

	s = "\xE4\xBA\xBB\xFF"
	t.Log(s)
	t.Log(len(s))
	s = "中"
	t.Log(len(s)) // 是byte 数
	c := [] rune(s)
	t.Log(len(c))

	t.Logf("中 unicode %x",c[0])
	t.Logf("中 utf-8 %x",s)
}

func TestSrtingToRune(t *testing.T){
	s := "中华人民共和国"
	for _,c := range s{
		t.Logf("%[1]c %[1]x",c)
	}
}