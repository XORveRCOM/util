package json

import "strconv"

type ElemFloat interface {
	Element
	// Float は数値表現を返します。
	Float() float64
}

// elemFloat は数値型要素です
type elemFloat struct {
	element
	value float64
}

// NewelemFloat は要素を作成します。
func NewElemFloat(value float64) *elemFloat {
	return &elemFloat{value: value}
}

// JSON 表現として {...} や [...] や "string" というような文字列を返します。
func (e elemFloat) String() string {
	return e.Text()
}

// Text は文字列表現を返します。
func (e *elemFloat) Text() string {
	return strconv.FormatFloat(e.value, 'f', -1, 64)
}

// Type は要素の型を取得します。
func (e *elemFloat) Type() ElementType {
	return TypeFloat
}

// Paths は子供のパス要素の一覧を取得します。
func (e *elemFloat) Paths() []PathJSON {
	return []PathJSON{}
}

// Value はインタフェースとしての内容を取得します。
func (e *elemFloat) Value() interface{} {
	return e.value
}

// AsFloat は elemFloat にキャストします。
func (e *elemFloat) AsFloat() (ElemFloat, bool) {
	return e, true
}

// IsFloat は ElemFloat であるか判定します。
func (e *elemFloat) IsFloat() bool {
	return true
}

// Clone はディープコピーした Element を返します。
func (e *elemFloat) Clone() Element {
	return NewElemFloat(e.value)
}

// Float は数値表現を返します。
func (e *elemFloat) Float() float64 {
	return e.value
}
