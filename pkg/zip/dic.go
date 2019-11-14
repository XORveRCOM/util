package zip

import (
	azip "archive/zip"
	"fmt"
)

// Dictionary は zip 内のファイルの辞書です。
type Dictionary interface {
	// Contains は、そのファイルが存在する場合に true を返します。
	Contains(filepath string) bool
	// Get は、そのファイルの中身を取得します。
	Get(filepath string) ([]byte, error)
	// FilePaths はファイルの一覧を取得します。
	FilePaths() []string
	// Close は Dictionary をクローズします。
	Close()
}

type dic struct {
	arc   *azip.ReadCloser
	dic   map[string]*azip.File
	files []string
}

// OpenDictionary は Dictionary を取得します。
func OpenDictionary(zippath string) (Dictionary, error) {
	r, err := azip.OpenReader(zippath)
	if err != nil {
		return nil, err
	}
	dic := &dic{arc: r, dic: map[string]*azip.File{}, files: []string{}}

	for _, f := range dic.arc.File {
		if f.FileInfo().IsDir() {
			continue
		}
		dic.dic[f.Name] = f
		dic.files = append(dic.files, f.Name)
	}
	return dic, nil
}

func (d *dic) Contains(filepath string) bool {
	_, ok := d.dic[filepath]
	return ok
}
func (d *dic) Get(filepath string) ([]byte, error) {
	f, ok := d.dic[filepath]
	if !ok {
		return []byte{}, fmt.Errorf("filepath not found: %s", filepath)
	}
	rc, err := f.Open()
	if err != nil {
		return []byte{}, err
	}
	defer rc.Close()
	if f.UncompressedSize64 == 0 {
		// サイズ0のファイルをRead()するとEOFエラーとなるため
		return []byte{}, nil
	}
	arr := make([]byte, f.UncompressedSize64)
	_, rerr := rc.Read(arr)
	return arr, rerr
}
func (d *dic) FilePaths() []string {
	return d.files
}
func (d *dic) Close() {
	d.arc.Close()
}
