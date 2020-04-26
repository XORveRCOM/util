package json

import "strconv"

type ElemBool interface {
	Element
	// Bool は論理値表現を返します。
	Bool() bool
}

// elemBool は論理値型要素です
type elemBool struct {
	value bool
}

// NewelemBool は要素を作成します。
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

// AsObject は ElemObject にキャストします。
func (e *elemBool) AsObject() (ElemObject, bool) {
	return nil, false
}

// AsArray は ElemArray にキャストします。
func (e *elemBool) AsArray() (ElemArray, bool) {
	return nil, false
}

// AsString は ElemString にキャストします。
func (e *elemBool) AsString() (ElemString, bool) {
	return nil, false
}

// AsFloat は ElemFloat にキャストします。
func (e *elemBool) AsFloat() (ElemFloat, bool) {
	return nil, false
}

// AsBool は elemBool にキャストします。
func (e *elemBool) AsBool() (ElemBool, bool) {
	return e, true
}

// Bool は論理値表現を返します。
func (e *elemBool) Bool() bool {
	return e.value
}
