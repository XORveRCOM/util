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
	// テスト用フック
	HookForTest() *Hook
}

// 標準メソッドのテスト用モックのフックポイント
type InterScanner interface {
	Err() error
	Scan() bool
	Text() string
}

func ifuncNewScanner(r io.Reader) InterScanner {
	return bufio.NewScanner(r)
}

type HookBufio struct {
	NewScanner func(r io.Reader) InterScanner
}

type InterFile interface {
	Close() error
	WriteString(s string) (n int, err error)
	Read(p []byte) (n int, err error)
}

func ifuncOpen(name string) (InterFile, error) {
	return os.Open(name)
}
func ifuncCreate(name string) (InterFile, error) {
	return os.Create(name)
}

type HookOs struct {
	Open   func(name string) (InterFile, error)
	Create func(name string) (InterFile, error)
}
type Hook struct {
	Bufio *HookBufio
	Os    *HookOs
}

// textLines はテキストを格納します。
// [TODO] 文字エンコーディング関係機能
type textLines struct {
	// 文字列スライス
	lines []string
	// フック
	hook *Hook
}

// New は空のテキストを返します。
func New() TextLines {
	hook := &Hook{
		Bufio: &HookBufio{
			NewScanner: ifuncNewScanner,
		},
		Os: &HookOs{
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
	f, err := t.hook.Os.Open(filename)
	if err != nil {
		return fmt.Errorf("os.Open(%s) return %w", filename, err)
	}
	defer f.Close()

	s := t.hook.Bufio.NewScanner(f)
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
	f, err := t.hook.Os.Create(filename)
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

// Deprecated: should not be used for anything other than testing
func (t *textLines) HookForTest() *Hook {
	return t.hook
}
