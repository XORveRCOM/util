package json

type ElemNull interface {
	Element
}

// ElemNull はNULL値型要素です
type elemNull struct{}

// NewElemNull は要素を作成します。
func NewElemNull() ElemNull {
	return &elemNull{}
}

func (e *elemNull) String() string {
	return e.Text()
}

// Type は要素の型を取得します。
func (e *elemNull) Type() ElementType {
	return TypeNull
}

// Paths は子供のパス要素の一覧を取得します。
func (e *elemNull) Paths() []PathJSON {
	return []PathJSON{}
}

// Value はインタフェースとしての内容を取得します。
func (e *elemNull) Value() interface{} {
	return nil
}

// Text は文字列表現を返します。
func (e *elemNull) Text() string {
	return "null"
}

// AsObject は ElemObject にキャストします。
func (e *elemNull) AsObject() (ElemObject, bool) {
	return nil, false
}

// AsArray は ElemArray にキャストします。
func (e *elemNull) AsArray() (ElemArray, bool) {
	return nil, false
}

// AsString は ElemString にキャストします。
func (e *elemNull) AsString() (ElemString, bool) {
	return nil, false
}

// AsFloat は ElemFloat にキャストします。
func (e *elemNull) AsFloat() (ElemFloat, bool) {
	return nil, false
}

// AsBool は ElemBool にキャストします。
func (e *elemNull) AsBool() (ElemBool, bool) {
	return nil, false
}

// IsObject は ElemObject であるか判定します。
func (e *elemNull) IsObject() bool {
	return false
}

// IsArray は ElemArray であるか判定します。
func (e *elemNull) IsArray() bool {
	return false
}

// AsString は ElemString であるか判定します。
func (e *elemNull) IsString() bool {
	return false
}

// IsFloat は ElemFloat であるか判定します。
func (e *elemNull) IsFloat() bool {
	return false
}

// IsBool は ElemBool であるか判定します。
func (e *elemNull) IsBool() bool {
	return false
}

// IsNull は ElemNull であるか判定します。
func (e *elemNull) IsNull() bool {
	return true
}
