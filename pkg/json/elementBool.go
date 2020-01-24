package json

import "strconv"

// ElemBool は論理値型要素です
type ElemBool struct {
	value bool
}

// NewElemBool は要素を作成します。
func NewElemBool(value bool) *ElemBool {
	return &ElemBool{value: value}
}

// Bool は論理値表現を返します。
func (e *ElemBool) Bool() bool {
	return e.value
}

// Value はインタフェースとしての内容を取得します。
func (e *ElemBool) Value() interface{} {
	return e.value
}

// Paths は要素の一覧を取得します。
func (e *ElemBool) Paths() []PathJSON {
	return []PathJSON{}
}

// Type は要素の型を取得します。
func (e *ElemBool) Type() ElementType {
	return TypeBool
}

// Text は文字列表現を返します。
func (e *ElemBool) Text() string {
	return strconv.FormatBool(e.value)
}

// AsObject は ElemObject にキャストします。
func (e *ElemBool) AsObject() (*ElemObject, bool) {
	return nil, false
}

// AsArray は ElemArray にキャストします。
func (e *ElemBool) AsArray() (*ElemArray, bool) {
	return nil, false
}

// AsString は ElemString にキャストします。
func (e *ElemBool) AsString() (*ElemString, bool) {
	return nil, false
}

// AsFloat は ElemFloat にキャストします。
func (e *ElemBool) AsFloat() (*ElemFloat, bool) {
	return nil, false
}

// AsBool は ElemBool にキャストします。
func (e *ElemBool) AsBool() (*ElemBool, bool) {
	return e, true
}

func (e ElemBool) String() string {
	return e.Text()
}
