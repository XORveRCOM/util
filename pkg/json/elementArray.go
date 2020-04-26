package json

import (
	"strconv"
	"strings"
)

// ElemArray は配列型要素です
type ElemArray interface {
	Element
	// Size は子供の数を返します。
	Size() int
	// Child は子供の要素を返します。
	Child(num int) Element
	// Append は要素を追加します。
	Append(child ...Element) ElemArray
}

// ElemArray は配列型要素です
type elemArray struct {
	value []Element
}

// NewElemArray は要素を作成します。
func NewElemArray() ElemArray {
	return &elemArray{value: []Element{}}
}

func (e *elemArray) String() string {
	return e.Text()
}

// Text は文字列表現を返します。
func (e *elemArray) Text() string {
	arr := []string{}
	for _, value := range e.value {
		arr = append(arr, value.String())
	}
	ret := "["
	ret += strings.Join(arr, ", ")
	ret += "]"
	return ret
}

// Type は要素の型を取得します。
func (e *elemArray) Type() ElementType {
	return TypeArray
}

// Paths は子供のパス要素の一覧を取得します。
func (e *elemArray) Paths() []PathJSON {
	arr := []PathJSON{}
	for i, item := range e.value {
		key := PathJSON("[" + strconv.Itoa(i) + "]")
		arr = append(arr, key)
		for _, sub := range item.Paths() {
			wk := key + jsonPathSeparater + sub
			arr = append(arr, wk)
		}
	}
	return arr
}

// Value は値を返します。
func (e *elemArray) Value() interface{} {
	ret := make([]interface{}, 0)
	for _, val := range e.value {
		ret = append(ret, val.Value())
	}
	return interface{}(ret)
}

// AsObject は ElemObject にキャストします。
func (e *elemArray) AsObject() (ElemObject, bool) {
	return nil, false
}

// AsArray は ElemArray にキャストします。
func (e *elemArray) AsArray() (ElemArray, bool) {
	return e, true
}

// AsString は ElemString にキャストします。
func (e *elemArray) AsString() (ElemString, bool) {
	return nil, false
}

// AsFloat は ElemFloat にキャストします。
func (e *elemArray) AsFloat() (ElemFloat, bool) {
	return nil, false
}

// AsBool は ElemBool にキャストします。
func (e *elemArray) AsBool() (ElemBool, bool) {
	return nil, false
}

// Size は子供の数を返します。
func (e *elemArray) Size() int {
	return len(e.value)
}

// Child は子供の要素を返します。
func (e *elemArray) Child(num int) Element {
	if num < 0 || e.Size() <= num {
		return NewElemNull()
	}
	return e.value[num]
}

// Append は要素を追加します。
func (e *elemArray) Append(child ...Element) ElemArray {
	e.value = append(e.value, child...)
	return e
}
