package json

type ElemString interface {
	Element
	// Length は文字列長を返します。
	Length() int
}

// ElemString は文字列型要素です
type elemString struct {
	value string
}

// NewElemString は要素を作成します。
func NewElemString(value string) ElemString {
	return &elemString{value: value}
}

func (e *elemString) String() string {
	return "\"" + e.Text() + "\""
}

// Text は文字列表現を返します。
func (e *elemString) Text() string {
	return e.value
}

// Type は要素の型を取得します。
func (e *elemString) Type() ElementType {
	return TypeString
}

// Paths は子供のパス要素の一覧を取得します。
func (e *elemString) Paths() []PathJSON {
	return []PathJSON{}
}

// Value はインタフェースとしての内容を取得します。
func (e *elemString) Value() interface{} {
	return e.value
}

// AsObject は ElemObject にキャストします。
func (e *elemString) AsObject() (ElemObject, bool) {
	return nil, false
}

// AsArray は ElemArray にキャストします。
func (e *elemString) AsArray() (ElemArray, bool) {
	return nil, false
}

// AsString は ElemString にキャストします。
func (e *elemString) AsString() (ElemString, bool) {
	return e, true
}

// AsFloat は ElemFloat にキャストします。
func (e *elemString) AsFloat() (ElemFloat, bool) {
	return nil, false
}

// AsBool は ElemBool にキャストします。
func (e *elemString) AsBool() (ElemBool, bool) {
	return nil, false
}

// IsObject は ElemObject であるか判定します。
func (e *elemString) IsObject() bool {
	return false
}

// IsArray は ElemArray であるか判定します。
func (e *elemString) IsArray() bool {
	return false
}

// AsString は ElemString であるか判定します。
func (e *elemString) IsString() bool {
	return true
}

// IsFloat は ElemFloat であるか判定します。
func (e *elemString) IsFloat() bool {
	return false
}

// IsBool は ElemBool であるか判定します。
func (e *elemString) IsBool() bool {
	return false
}

// IsNull は ElemNull であるか判定します。
func (e *elemString) IsNull() bool {
	return false
}

// Length は文字列長を返します。
func (e *elemString) Length() int {
	return len(e.value)
}
