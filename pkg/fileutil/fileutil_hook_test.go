package fileutil_test

import (
	"errors"
	"io/fs"
	fpath "path/filepath"
	"testing"

	"github.com/xorvercom/util/pkg/fileutil"
	"github.com/xorvercom/util/pkg/text"
)

func TestHookOs(t *testing.T) {
	hook := fileutil.HookForTest()

	func() {
		ie := hook.Os.IsExist
		defer func() {
			hook.Os.IsExist = ie
		}()

		hook.Os.IsExist = func(err error) bool {
			return true
		}

		if err := fileutil.FileIfDelete(""); err == nil {
			t.Fail()
		}

		if err := fileutil.FileCopy("", ""); err == nil {
			t.Fail()
		}

		if err := fileutil.FileIfMove("", ""); err == nil {
			t.Fail()
		}
	}()

	fileutil.TempSpace(func(tempdir string) error {
		ie := hook.Os.ReadFile
		defer func() {
			hook.Os.ReadFile = ie
		}()
		hook.Os.ReadFile = func(name string) ([]byte, error) {
			return nil, errors.New("ReadFile() error")
		}

		src := fpath.Join(tempdir, "src.txt")
		dst := fpath.Join(tempdir, "dst.txt")
		txt := text.New()
		txt.SaveTo(src)
		if err := fileutil.FileCopy(src, dst); err == nil {
			t.Fail()
		}
		return nil
	})

	func() {
		ie := hook.Os.MkdirAll
		defer func() {
			hook.Os.MkdirAll = ie
		}()
		hook.Os.MkdirAll = func(path string, perm fs.FileMode) error {
			return errors.New("MkdirAll() error")
		}
		err := fileutil.TempSpace(func(tempdir string) error {
			return nil
		})
		if err == nil {
			t.Fail()
		}
	}()
}
