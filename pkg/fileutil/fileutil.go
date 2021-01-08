package fileutil

import (
	"io/ioutil"
	"os"
)

// FileCopy ファイルコピー
func FileCopy(src, dst string) error {
	// コピー先を消去
	err := FileIfDelete(dst)
	if err != nil {
		return err
	}
	// 読み出し
	b, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	// 書き出し
	return ioutil.WriteFile(dst, b, 0644)
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
