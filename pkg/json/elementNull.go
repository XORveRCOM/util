package json

type ElemNull interface {
	Element
}

// ElemNull はNULL値型要素です
type elemNull struct {
	element
}

// NewElemNull は要素を作成します。
func NewElemNull() ElemNull {
	return &elemNull{}
}

func (e *elemNull) String() string {
	return e.Text()
}

// Text は文字列表現を返します。
func (e *elemNull) Text() string {
	return "null"
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

// IsNull は ElemNull であるか判定します。
func (e *elemNull) IsNull() bool {
	return true
}
