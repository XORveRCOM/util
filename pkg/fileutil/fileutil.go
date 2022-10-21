package fileutil

import (
	"os"
	"sort"
)

// FileCopy ファイルコピー
func FileCopy(src, dst string) error {
	// コピー先を消去
	err := FileIfDelete(dst)
	if err != nil {
		return err
	}
	// 読み出し
	b, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	// 書き出し
	return os.WriteFile(dst, b, 0644)
}

// FileIfMove ファイルを移動 (移動元が無くても正常終了)
func FileIfMove(src, dst string) error {
	// 移動先を消去
	err := FileIfDelete(dst)
	if err != nil {
		return err
	}
	// 移動
	return os.Rename(src, dst)
}

// FileIfDelete ファイル削除 (削除元が無くても正常終了)
func FileIfDelete(file string) error {
	_, err := os.Stat(file)
	if err != nil {
		if os.IsExist(err) {
			// ファイルがあるのにstatが失敗
			return err
		}
		// ファイルが無かったのでnoop
		return nil
	}
	// 削除
	return os.Remove(file)
}

// フアイルが存在するかをチェック
func FileExists(filename string) bool {
	f, err := os.Stat(filename)
	return err == nil && false == f.IsDir()
}

// ディレクトリが存在するかをチェック
func DirExists(filename string) bool {
	f, err := os.Stat(filename)
	return err == nil && f.IsDir()
}

// FilesList は folder パスの全ファイル名の文字列配列を返します
func FilesList(folder string) []string {
	res := []string{}
	if files, err := os.ReadDir(folder); err == nil {
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
	if files, err := os.ReadDir(folder); err == nil {
		for _, file := range files {
			if file.IsDir() {
				res = append(res, file.Name())
			}
		}
	}
	sort.Sort(sort.StringSlice(res))
	return res
}
