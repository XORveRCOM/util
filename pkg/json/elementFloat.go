package json

import "strconv"

// ElemFloat は数値型要素です
type ElemFloat struct {
	value float64
}

// NewElemFloat は要素を作成します。
func NewElemFloat(value float64) *ElemFloat {
	return &ElemFloat{value: value}
}

// Float は数値表現を返します。
func (e *ElemFloat) Float() float64 {
	return e.value
}

// Value はインタフェースとしての内容を取得します。
func (e *ElemFloat) Value() interface{} {
	return e.value
}

// Paths は要素の一覧を取得します。
func (e *ElemFloat) Paths() []PathJSON {
	return []PathJSON{}
}

// Type は要素の型を取得します。
func (e *ElemFloat) Type() ElementType {
	return TypeFloat
}

// Text は文字列表現を返します。
func (e *ElemFloat) Text() string {
	return strconv.FormatFloat(e.value, 'f', -1, 64)
}

// AsObject は ElemObject にキャストします。
func (e *ElemFloat) AsObject() (*ElemObject, bool) {
	return nil, false
}

// AsArray は ElemArray にキャストします。
func (e *ElemFloat) AsArray() (*ElemArray, bool) {
	return nil, false
}

// AsString は ElemString にキャストします。
func (e *ElemFloat) AsString() (*ElemString, bool) {
	return nil, false
}

// AsFloat は ElemFloat にキャストします。
func (e *ElemFloat) AsFloat() (*ElemFloat, bool) {
	return e, true
}

// AsBool は ElemBool にキャストします。
func (e *ElemFloat) AsBool() (*ElemBool, bool) {
	return nil, false
}

func (e ElemFloat) String() string {
	return e.Text()
}
