package json

type ElemString interface {
	Element
	// Length は文字列長を返します。
	Length() int
}

// ElemString は文字列型要素です
type _ElemString struct {
	value string
}

// NewElemString は要素を作成します。
func NewElemString(value string) ElemString {
	return &_ElemString{value: value}
}

func (e *_ElemString) String() string {
	return "\"" + e.Text() + "\""
}

// Text は文字列表現を返します。
func (e *_ElemString) Text() string {
	return e.value
}

// Type は要素の型を取得します。
func (e *_ElemString) Type() ElementType {
	return TypeString
}

// Paths は子供のパス要素の一覧を取得します。
func (e *_ElemString) Paths() []PathJSON {
	return []PathJSON{}
}

// Value はインタフェースとしての内容を取得します。
func (e *_ElemString) Value() interface{} {
	return e.value
}

// AsObject は ElemObject にキャストします。
func (e *_ElemString) AsObject() (ElemObject, bool) {
	return nil, false
}

// AsArray は ElemArray にキャストします。
func (e *_ElemString) AsArray() (ElemArray, bool) {
	return nil, false
}

// AsString は ElemString にキャストします。
func (e *_ElemString) AsString() (ElemString, bool) {
	return e, true
}

// AsFloat は ElemFloat にキャストします。
func (e *_ElemString) AsFloat() (ElemFloat, bool) {
	return nil, false
}

// AsBool は ElemBool にキャストします。
func (e *_ElemString) AsBool() (ElemBool, bool) {
	return nil, false
}

// Length は文字列長を返します。
func (e *_ElemString) Length() int {
	return len(e.value)
}
