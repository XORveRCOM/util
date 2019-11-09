// Package textlines はテキストの読み書きを行うパッケージです。
package textlines

import (
	"bufio"
	"fmt"
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

// textLines はテキストを格納します。
// [TODO] 文字エンコーディング関係機能
type textLines struct {
	// 文字列スライス
	lines []string
}

// New は空のテキストを返します。
func New() TextLines {
	return &textLines{lines: []string{}}
}

// LoadFrom はファイルからテキストを読み込みます。
func LoadFrom(filename string) (TextLines, error) {
	t := &textLines{lines: []string{}}
	return t, t.LoadFrom(filename)
}

// LoadFrom はファイルからテキストを読み込みます。
func (t *textLines) LoadFrom(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("os.Open(%s) return %v", filename, err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		t.lines = append(t.lines, s.Text())
	}
	err = s.Err()
	if err != nil {
		err = fmt.Errorf("scanner.Scan(%s) return %v", filename, err)
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
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("os.Create(%s) return %v", filename, err)
	}
	defer f.Close()

	for _, line := range t.Lines() {
		_, err := f.WriteString(line + "\n")
		if err != nil {
			return fmt.Errorf("Writer(%s).WriteString() return %v", filename, err)
		}
	}
	return nil
}
