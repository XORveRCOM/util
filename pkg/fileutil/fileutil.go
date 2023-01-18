package fileutil

import (
	"io"
	"io/fs"
	"os"
	"sort"
)

type HookOs struct {
	ReadFile  func(name string) ([]byte, error)
	WriteFile func(name string, data []byte, perm fs.FileMode) error
	Rename    func(oldpath string, newpath string) error
	Stat      func(name string) (fs.FileInfo, error)
	Remove    func(name string) error
	ReadDir   func(name string) ([]fs.DirEntry, error)
	IsExist   func(err error) bool
	MkdirAll  func(path string, perm fs.FileMode) error
	Open      func(name string) (*os.File, error)
	Create    func(name string) (*os.File, error)
}
type HookIo struct {
	Copy func(dst io.Writer, src io.Reader) (written int64, err error)
}
type Hook struct {
	Os *HookOs
	Io *HookIo
}

var hook *Hook

func init() {
	hook = &Hook{}
	hook.Os = &HookOs{}
	hook.Os.ReadFile = os.ReadFile
	hook.Os.WriteFile = os.WriteFile
	hook.Os.Rename = os.Rename
	hook.Os.Stat = os.Stat
	hook.Os.Remove = os.Remove
	hook.Os.ReadDir = os.ReadDir
	hook.Os.IsExist = os.IsExist
	hook.Os.MkdirAll = os.MkdirAll
	hook.Os.Open = os.Open
	hook.Os.Create = os.Create
	hook.Io = &HookIo{}
	hook.Io.Copy = io.Copy
}

// Deprecated: should not be used for anything other than testing
func HookForTest() *Hook {
	return hook
}

// FileCopy ファイルコピー
func FileCopy(src, dst string) error {
	// コピー先を消去
	err := FileIfDelete(dst)
	if err != nil {
		return err
	}
	s, err := hook.Os.Open(src)
	if err != nil {
		return err
	}
	defer s.Close()
	d, err := hook.Os.Create(dst)
	if err != nil {
		return err
	}
	defer d.Close()
	_, err = hook.Io.Copy(d, s)
	if err != nil {
		return err
	}
	return nil
}

// FileIfMove ファイルを移動 (移動元が無くても正常終了)
func FileIfMove(src, dst string) error {
	// 移動先を消去
	err := FileIfDelete(dst)
	if err != nil {
		return err
	}
	// 移動
	return hook.Os.Rename(src, dst)
}

// FileIfDelete ファイル削除 (削除元が無くても正常終了)
func FileIfDelete(file string) error {
	_, err := hook.Os.Stat(file)
	if err != nil {
		if hook.Os.IsExist(err) {
			// ファイルがあるのにstatが失敗
			return err
		}
		// ファイルが無かったのでnoop
		return nil
	}
	// 削除
	return hook.Os.Remove(file)
}

// フアイルが存在するかをチェック
func FileExists(filename string) bool {
	f, err := hook.Os.Stat(filename)
	return err == nil && false == f.IsDir()
}

// ディレクトリが存在するかをチェック
func DirExists(filename string) bool {
	f, err := hook.Os.Stat(filename)
	return err == nil && f.IsDir()
}

// FilesList は folder パスの全ファイル名の文字列配列を返します
func FilesList(folder string) []string {
	res := []string{}
	if files, err := hook.Os.ReadDir(folder); err == nil {
		for _, file := range files {
			if false == file.IsDir() {
				res = append(res, file.Name())
			}
		}
	}
	sort.Sort(sort.StringSlice(res))
	return res
}

// DirsList は
func DirsList(folder string) []string {
	res := []string{}
	if files, err := hook.Os.ReadDir(folder); err == nil {
		for _, file := range files {
			if file.IsDir() {
				res = append(res, file.Name())
			}
		}
	}
	sort.Sort(sort.StringSlice(res))
	return res
}
