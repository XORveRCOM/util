package textlines

import (
	"os"
	"path"
	"testing"
)

func TestAll(t *testing.T) {
	tempdir := os.TempDir()
	var txt TextLines
	var err error
	txt = New()

	filename := path.Join(tempdir, "test.txt")
	txt.Append("abcd")
	txt.Append("0123")
	if err = txt.SaveTo(filename); err != nil {
		t.Fail()
	}

	txt, err = LoadFrom(filename)
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

	func() {
		txt = (*textLines)(nil)
		err = txt.SaveTo(filename)
		if err == nil {
			t.Fail()
		}
	}()
}
