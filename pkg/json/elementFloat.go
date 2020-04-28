package json

import "strconv"

type ElemFloat interface {
	Element
	// Float は数値表現を返します。
	Float() float64
}

// elemFloat は数値型要素です
type elemFloat struct {
	value float64
}

// NewelemFloat は要素を作成します。
func NewElemFloat(value float64) *elemFloat {
	return &elemFloat{value: value}
}

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

// AsObject は ElemObject にキャストします。
func (e *elemFloat) AsObject() (ElemObject, bool) {
	return nil, false
}

// AsArray は ElemArray にキャストします。
func (e *elemFloat) AsArray() (ElemArray, bool) {
	return nil, false
}

// AsString は ElemString にキャストします。
func (e *elemFloat) AsString() (ElemString, bool) {
	return nil, false
}

// AsFloat は elemFloat にキャストします。
func (e *elemFloat) AsFloat() (ElemFloat, bool) {
	return e, true
}

// AsBool は ElemBool にキャストします。
func (e *elemFloat) AsBool() (ElemBool, bool) {
	return nil, false
}

// IsObject は ElemObject であるか判定します。
func (e *elemFloat) IsObject() bool {
	return false
}

// IsArray は ElemArray であるか判定します。
func (e *elemFloat) IsArray() bool {
	return false
}

// AsString は ElemString であるか判定します。
func (e *elemFloat) IsString() bool {
	return false
}

// IsFloat は ElemFloat であるか判定します。
func (e *elemFloat) IsFloat() bool {
	return true
}

// IsBool は ElemBool であるか判定します。
func (e *elemFloat) IsBool() bool {
	return false
}

// IsNull は ElemNull であるか判定します。
func (e *elemFloat) IsNull() bool {
	return false
}

// Float は数値表現を返します。
func (e *elemFloat) Float() float64 {
	return e.value
}
