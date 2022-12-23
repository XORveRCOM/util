// Package textlines はテキストの読み書きを行うパッケージです。
package text

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// TextLines はテキストを格納します。
type TextLines interface {
	// Lines は文字列スライスを返します。
	Lines() []string
	// Append はテキストに文字列を追加します。
	Append(string)
	// SaveTo はファイルにテキストを書き込みます。(LF,UTF-8)
	SaveTo(string) error
	// LoadFrom はファイルからテキストを読み込みます。
	LoadFrom(string) error
}

// 標準メソッドのテスト用モックのフックポイント
type iScanner interface {
	Err() error
	Scan() bool
	Text() string
}

func ifuncNewScanner(r io.Reader) iScanner {
	return bufio.NewScanner(r)
}

type hookBufio struct {
	NewScanner func(r io.Reader) iScanner
}

type iFile interface {
	Close() error
	WriteString(s string) (n int, err error)
	Read(p []byte) (n int, err error)
}

func ifuncOpen(name string) (iFile, error) {
	return os.Open(name)
}
func ifuncCreate(name string) (iFile, error) {
	return os.Create(name)
}

type hookOs struct {
	Open   func(name string) (iFile, error)
	Create func(name string) (iFile, error)
}
type hook struct {
	bufio *hookBufio
	os    *hookOs
}

// textLines はテキストを格納します。
// [TODO] 文字エンコーディング関係機能
type textLines struct {
	// 文字列スライス
	lines []string
	// フック
	hook *hook
}

// New は空のテキストを返します。
func New() TextLines {
	hook := &hook{
		bufio: &hookBufio{
			NewScanner: ifuncNewScanner,
		},
		os: &hookOs{
			Open:   ifuncOpen,
			Create: ifuncCreate,
		},
	}
	return &textLines{lines: []string{}, hook: hook}
}

// LoadFrom はファイルからテキストを読み込みます。
func LoadFrom(filename string) (TextLines, error) {
	t := New()
	return t, t.LoadFrom(filename)
}

// LoadFrom はファイルからテキストを読み込みます。
func (t *textLines) LoadFrom(filename string) error {
	f, err := t.hook.os.Open(filename)
	if err != nil {
		return fmt.Errorf("os.Open(%s) return %w", filename, err)
	}
	defer f.Close()

	s := t.hook.bufio.NewScanner(f)
	for s.Scan() {
		t.lines = append(t.lines, s.Text())
	}
	err = s.Err()
	if err != nil {
		return fmt.Errorf("scanner.Scan(%s) return %w", filename, err)
	}
	return err
}

// Lines は文字列スライスを返します。
func (t *textLines) Lines() []string {
	return t.lines
}

// Append はテキストに文字列を追加します。
func (t *textLines) Append(line string) {
	t.lines = append(t.lines, line)
}

// SaveTo はファイルにテキストを書き込みます。
func (t *textLines) SaveTo(filename string) error {
	if t == nil {
		return fmt.Errorf("nil.SaveTo(%s)", filename)
	}
	f, err := t.hook.os.Create(filename)
	if err != nil {
		return fmt.Errorf("os.Create(%s) return %w", filename, err)
	}
	defer f.Close()

	for _, line := range t.Lines() {
		_, err := f.WriteString(line + "\n")
		if err != nil {
			return fmt.Errorf("Writer(%s).WriteString() return %w", filename, err)
		}
	}
	return nil
}
