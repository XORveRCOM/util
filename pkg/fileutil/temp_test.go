package fileutil_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/xorvercom/util/pkg/fileutil"
)

func ExampleTempSpace() {
	err := fileutil.TempSpace(func(tempdir string) error {
		filename := filepath.Join(tempdir, "test.txt")
		f, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = f.WriteString("data")
		if err != nil {
			return err
		}
		f.Close()
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func TestTempSpace(t *testing.T) {
	var tdir string
	err := fileutil.TempSpace(func(tempdir string) error {
		tdir = tempdir
		data := "ABCD"
		filename := filepath.Join(tempdir, "test.txt")
		f, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = f.WriteString(data)
		if err != nil {
			return err
		}
		f.Close()
		b, err := os.ReadFile(filename)
		if err != nil {
			return err
		}
		if string(b) != data {
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

func TestRunPanic(t *testing.T) {
	err := fileutil.TempSpace(func(tempdir string) error {
		panic("err")
	})
	if err == nil {
		t.Fatal()
	}
}
