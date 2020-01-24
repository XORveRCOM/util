package json

// ElemString は文字列型要素です
type ElemString struct {
	value string
}

// NewElemString は要素を作成します。
func NewElemString(value string) *ElemString {
	return &ElemString{value: value}
}

// Value はインタフェースとしての内容を取得します。
func (e *ElemString) Value() interface{} {
	return e.value
}

// Paths は要素の一覧を取得します。
func (e *ElemString) Paths() []PathJSON {
	return []PathJSON{}
}

// Type は要素の型を取得します。
func (e *ElemString) Type() ElementType {
	return TypeString
}

// Text は文字列表現を返します。
func (e *ElemString) Text() string {
	return e.value
}

// AsObject は ElemObject にキャストします。
func (e *ElemString) AsObject() (*ElemObject, bool) {
	return nil, false
}

// AsArray は ElemArray にキャストします。
func (e *ElemString) AsArray() (*ElemArray, bool) {
	return nil, false
}

// AsString は ElemString にキャストします。
func (e *ElemString) AsString() (*ElemString, bool) {
	return e, true
}

// AsFloat は ElemFloat にキャストします。
func (e *ElemString) AsFloat() (*ElemFloat, bool) {
	return nil, false
}

// AsBool は ElemBool にキャストします。
func (e *ElemString) AsBool() (*ElemBool, bool) {
	return nil, false
}

func (e ElemString) String() string {
	return "\"" + e.Text() + "\""
}
