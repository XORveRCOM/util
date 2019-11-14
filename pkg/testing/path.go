package testing

import (
	path "path/filepath"
	"runtime"
)

// GetFilepath は、そのソースファイルの置いてあるフォルダのパスを起点としたファイルのパスを生成します。
// os.Getwd() でも良いですが、カレントフォルダを使うのはちょっと気持ち悪かったので。
// testing パッケージに配置してある理由は、ファイルのパスを起点としたファイルアクセスはテスト以外では使うべきでないという理由です。
func GetFilepath(filepath string) string {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		return ""
	}
	return path.Join(path.Dir(file), filepath)
}
