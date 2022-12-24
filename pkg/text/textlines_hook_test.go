package text_test

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/xorvercom/util/pkg/text"
)

type scanner struct {
}

func (s *scanner) Scan() bool {
	return false
}
func (s *scanner) Err() error {
	return errors.New("Scanner error")
}
func (s *scanner) Text() string {
	return ""
}
func TestHookOs(t *testing.T) {
	filename := filepath.Join(os.TempDir(), "test.txt")
	txt := text.New()

	txt.Append("abcd")
	txt.Append("0123")
	if err := txt.SaveTo(filename); err != nil {
		t.Fail()
	}

	// Scanner エラーテスト
	hook := txt.HookForTest()
	ns := hook.Bufio.NewScanner
	hook.Bufio.NewScanner = func(r io.Reader) text.InterScanner {
		return &scanner{}
	}
	err := txt.LoadFrom(filename)
	if err == nil {
		t.Fail()
	}
	hook.Bufio.NewScanner = ns

	// WriteString エラーテスト
	c := hook.Os.Create
	hook.Os.Create = func(name string) (text.InterFile, error) {
		return &dummyFile{}, nil
	}
	txt.Append("")
	if err := txt.SaveTo(filename); err == nil {
		t.Fail()
	}
	hook.Os.Create = c
}

type dummyFile struct{}
func (f *dummyFile) Close() error {
	return nil
}
func (d *dummyFile)WriteString (s string) (n int, err error) {
	return 0, errors.New("WriteString() error")
}
func (f *dummyFile) Read(p []byte) (n int, err error) {
	return 0, nil
}