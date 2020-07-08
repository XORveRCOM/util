package json

import "strconv"

// ElemBool は論理値型要素です
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
func NewElemBool(value bool) ElemBool {
	return &elemBool{value: value}
}

// JSON 表現として {...} や [...] や "string" というような文字列を返します。
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

// Clone はディープコピーした Element を返します。
func (e *elemBool) Clone() Element {
	return NewElemBool(e.value)
}

// Bool は論理値表現を返します。
func (e *elemBool) Bool() bool {
	return e.value
}
