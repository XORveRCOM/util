package json

// ElemNull はNULL値型要素です
type ElemNull struct{}

// NewElemNull は要素を作成します。
func NewElemNull() *ElemNull {
	return &ElemNull{}
}

// Value はインタフェースとしての内容を取得します。
func (e *ElemNull) Value() interface{} {
	return nil
}

// Paths は要素の一覧を取得します。
func (e *ElemNull) Paths() []PathJSON {
	return []PathJSON{}
}

// Type は要素の型を取得します。
func (e *ElemNull) Type() ElementType {
	return TypeNull
}

// Text は文字列表現を返します。
func (e *ElemNull) Text() string {
	return "null"
}

// AsObject は ElemObject にキャストします。
func (e *ElemNull) AsObject() (*ElemObject, bool) {
	return nil, false
}

// AsArray は ElemArray にキャストします。
func (e *ElemNull) AsArray() (*ElemArray, bool) {
	return nil, false
}

// AsString は ElemString にキャストします。
func (e *ElemNull) AsString() (*ElemString, bool) {
	return nil, false
}

// AsFloat は ElemFloat にキャストします。
func (e *ElemNull) AsFloat() (*ElemFloat, bool) {
	return nil, false
}

// AsBool は ElemBool にキャストします。
func (e *ElemNull) AsBool() (*ElemBool, bool) {
	return nil, false
}

func (e ElemNull) String() string {
	return e.Text()
}
