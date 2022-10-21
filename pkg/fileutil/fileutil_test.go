package fileutil_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/xorvercom/util/pkg/fileutil"
)

func TestDirsList(t *testing.T) {
	var tdir string
	err := fileutil.TempSpace(func(tempdir string) error {
		tdir = tempdir

		var dirs []string
		dirs = fileutil.DirsList(tempdir)
		if len(dirs) != 0 {
			return fmt.Errorf("err")
		}

		var dirname string
		dirname = filepath.Join(tempdir, "dir2")
		if fileutil.DirExists(dirname) {
			return fmt.Errorf("err")
		}
		os.Mkdir(dirname, os.ModeDir)
		dirs = fileutil.DirsList(tempdir)
		if len(dirs) != 1 {
			return fmt.Errorf("err")
		}
		if dirs[0] != "dir2" {
			return fmt.Errorf("err")
		}

		dirname = filepath.Join(tempdir, "dir1")
		os.Mkdir(dirname, os.ModeDir)
		dirs = fileutil.DirsList(tempdir)
		if len(dirs) != 2 {
			return fmt.Errorf("err")
		}
		if dirs[1] != "dir2" {
			return fmt.Errorf("err")
		}
		if dirs[0] != "dir1" {
			return fmt.Errorf("err")
		}

		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if fileutil.FileExists(tdir) {
		t.Fatal(tdir)
	}
}

func TestFileList(t *testing.T) {
	var tdir string
	err := fileutil.TempSpace(func(tempdir string) error {
		tdir = tempdir

		var dirs []string
		var dirname string
		dirname = filepath.Join(tempdir, "dir2")
		os.Mkdir(dirname, os.ModeDir)
		dirs = fileutil.DirsList(tempdir)
		if len(dirs) != 1 {
			return fmt.Errorf("err")
		}

		var files []string
		files = fileutil.FilesList(tempdir)
		if len(files) != 0 {
			return fmt.Errorf("err")
		}

		var filename2 string
		filename2 = filepath.Join(tempdir, "file2.txt")
		f, err := os.Create(filename2)
		if err != nil {
			return err
		}
		f.Close()
		files = fileutil.FilesList(tempdir)
		if len(files) != 1 {
			return fmt.Errorf("err")
		}
		if files[0] != "file2.txt" {
			return fmt.Errorf("err")
		}

		var filename1 string
		filename1 = filepath.Join(tempdir, "file1.txt")
		if fileutil.FileCopy(filename2, filename1) != nil {
			return fmt.Errorf("err")
		}
		files = fileutil.FilesList(tempdir)
		if len(files) != 2 {
			return fmt.Errorf("err")
		}
		if files[1] != "file2.txt" {
			return fmt.Errorf("err")
		}
		if files[0] != "file1.txt" {
			return fmt.Errorf("err")
		}

		err = fileutil.FileIfMove(filename2, filename1)
		if err!=nil {
			return err
		}
		files = fileutil.FilesList(tempdir)
		if len(files) != 1 {
			return fmt.Errorf("err")
		}
		if files[0] != "file1.txt" {
			return fmt.Errorf("err")
		}

		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if fileutil.FileExists(tdir) {
		t.Fatal(tdir)
	}
}
