package json

type ElemString interface {
	Element
	// Length は文字列長を返します。
	Length() int
}

// ElemString は文字列型要素です
type elemString struct {
	element
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

// AsString は ElemString にキャストします。
func (e *elemString) AsString() (ElemString, bool) {
	return e, true
}

// AsString は ElemString であるか判定します。
func (e *elemString) IsString() bool {
	return true
}

// Clone はディープコピーした Element を返します。
func (e *elemString) Clone() Element {
	return NewElemString(e.value)
}

// Length は文字列長を返します。
func (e *elemString) Length() int {
	return len(e.value)
}
