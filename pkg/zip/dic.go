package zip

import (
	azip "archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"golang.org/x/text/encoding/japanese"
)

// ZippedFile はzipのエントリとしてのファイルです。
type ZippedFile interface {
	NameUTF8() string
	File() *azip.File
}
type file struct {
	file *azip.File
	dic  *dic
}

// Dictionary は zip 内のファイルの辞書です。
type Dictionary interface {
	// Contains は、そのファイルが存在する場合に true を返します。
	Contains(filepath string) bool
	// Get は、そのファイルの中身を取得します。
	Get(filepath string) ([]byte, error)
	// GetReader は、そのファイルのReaderを取得します。
	GetReader(filepath string) (io.ReadCloser, error)
	// FileHeader は、そのファイルの情報を取得します。
	File(filepath string) ZippedFile
	// FilePaths はファイルの一覧を取得します。
	FilePaths() []string
	// Close は Dictionary をクローズします。
	Close()
	// 大文字小文字を区別する場合は真を返します。
	IsIgnoreCapital() bool
}

type dic struct {
	arc           *azip.ReadCloser
	dic           map[string]ZippedFile
	files         []string
	ignoreCapital bool
}

// OpenDictionary は Dictionary を取得します。
func OpenDictionary(zippath string, ignoreCapital bool) (Dictionary, error) {
	r, err := azip.OpenReader(zippath)
	if err != nil {
		return nil, err
	}
	dic := &dic{arc: r, dic: map[string]ZippedFile{}, files: []string{}, ignoreCapital: ignoreCapital}

	for _, f := range dic.arc.File {
		// if f.FileInfo().IsDir() {
		// 	continue
		// }
		file := &file{file: f, dic: dic}
		name := file.NameUTF8()
		dic.dic[name] = file
		dic.files = append(dic.files, name)
	}
	return dic, nil
}

// Contains は、そのファイルが存在する場合に true を返します。
func (d *dic) Contains(filepath string) bool {
	if d.ignoreCapital {
		filepath = strings.ToLower(filepath)
	}
	_, ok := d.dic[filepath]
	return ok
}

// Get は、そのファイルの中身を取得します。
func (d *dic) Get(filepath string) ([]byte, error) {
	rc, err := d.GetReader(filepath)
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	return ioutil.ReadAll(rc)
}

// GetReader は、そのファイルの中身を取得します。
func (d *dic) GetReader(filepath string) (io.ReadCloser, error) {
	f := d.File(filepath)
	if f == nil {
		return nil, fmt.Errorf("filepath not found: %s", filepath)
	}
	return f.File().Open()
}

// FileHeader は、そのファイルの情報を取得します。
func (d *dic) File(filepath string) ZippedFile {
	if d.ignoreCapital {
		filepath = strings.ToLower(filepath)
	}
	f, ok := d.dic[filepath]
	if !ok {
		return nil
	}
	return f
}

// FilePaths はファイルの一覧を取得します。
func (d *dic) FilePaths() []string {
	return d.files
}

// Close は Dictionary をクローズします。
func (d *dic) Close() {
	d.arc.Close()
}

// 大文字小文字を区別する場合は真を返します。
func (d *dic) IsIgnoreCapital() bool {
	return d.ignoreCapital
}

// NameUTF8 はUTF8に正規化したファイル名を取得します。
func (f *file) NameUTF8() string {
	name := f.file.Name
	if f.file.NonUTF8 {
		if n, err := japanese.ShiftJIS.NewDecoder().String(name); err == nil {
			name = n
		}
	}
	if f.dic.ignoreCapital {
		// キーは小文字で統一
		name = strings.ToLower(name)
	}
	return name
}

func (f *file) File() *azip.File {
	return f.file
}
