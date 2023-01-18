package fileutil_test

import (
	"errors"
	"io"
	"io/fs"
	"os"
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
		src := fpath.Join(tempdir, "src.txt")
		dst := fpath.Join(tempdir, "dst.txt")
		txt := text.New()
		txt.SaveTo(src)

		op := hook.Os.Open
		hook.Os.Open = func(name string) (*os.File, error) {
			return nil, errors.New("Open() error")
		}
		if err := fileutil.FileCopy(src, dst); err == nil {
			t.Fail()
		}
		hook.Os.Open = op

		cr := hook.Os.Create
		hook.Os.Create = func(name string) (*os.File, error) {
			return nil, errors.New("Open() error")
		}
		if err := fileutil.FileCopy(src, dst); err == nil {
			t.Fail()
		}
		hook.Os.Create = cr

		co := hook.Io.Copy
		hook.Io.Copy = func(dst io.Writer, src io.Reader) (written int64, err error) {
			return 0, errors.New("Open() error")
		}
		if err := fileutil.FileCopy(src, dst); err == nil {
			t.Fail()
		}
		hook.Io.Copy = co

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
