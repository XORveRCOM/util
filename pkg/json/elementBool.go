package json

import "strconv"

type ElemBool interface {
	Element
	// Bool は論理値表現を返します。
	Bool() bool
}

// elemBool は論理値型要素です
type elemBool struct {
	element
	value bool
}

// NewElemBool は要素を作成します。
func NewElemBool(value bool) *elemBool {
	return &elemBool{value: value}
}

func (e elemBool) String() string {
	return e.Text()
}

// Text は文字列表現を返します。
func (e *elemBool) Text() string {
	return strconv.FormatBool(e.value)
}

// Type は要素の型を取得します。
func (e *elemBool) Type() ElementType {
	return TypeBool
}

// Paths は子供のパス要素の一覧を取得します。
func (e *elemBool) Paths() []PathJSON {
	return []PathJSON{}
}

// Value はインタフェースとしての内容を取得します。
func (e *elemBool) Value() interface{} {
	return e.value
}

// AsBool は elemBool にキャストします。
func (e *elemBool) AsBool() (ElemBool, bool) {
	return e, true
}

// IsBool は ElemBool であるか判定します。
func (e *elemBool) IsBool() bool {
	return true
}

// Bool は論理値表現を返します。
func (e *elemBool) Bool() bool {
	return e.value
}
