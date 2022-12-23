package text_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/xorvercom/util/pkg/text"
)

func Example() {
	filename := filepath.Join(os.TempDir(), "test.txt")

	// 書き出し
	write := text.New()
	write.Append("abcd")
	write.Append("0123")
	if err := write.SaveTo(filename); err != nil {
		panic(err)
	}

	// 読み込み
	// read, err := LoadFrom(filename) でも同じ挙動となります
	read := text.New()
	if err := read.LoadFrom(filename); err != nil {
		panic("load error")
	}

	txt := read.Lines()
	fmt.Println(len(txt))
	fmt.Println(txt[0])
	fmt.Println(txt[1])

	// output:
	// 2
	// abcd
	// 0123
}

func TestAll(t *testing.T) {
	tempdir := os.TempDir()
	var txt text.TextLines
	var err error
	txt = text.New()

	filename := filepath.Join(tempdir, "test.txt")
	txt.Append("abcd")
	txt.Append("0123")
	if err = txt.SaveTo(filename); err != nil {
		t.Fail()
	}

	txt, err = text.LoadFrom(filename)
	if err != nil {
		t.Fail()
	}
	if txt.Lines()[0] != "abcd" {
		t.Fatal(txt.Lines()[0])
	}
	if txt.Lines()[1] != "0123" {
		t.Fatal(txt.Lines()[1])
	}

	func() {
		err = txt.SaveTo("\tfile")
		if err == nil {
			t.Fail()
		}
	}()
}
