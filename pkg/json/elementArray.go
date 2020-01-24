package json

import (
	"strconv"
	"strings"
)

// ElemArray は配列型要素です
type ElemArray struct {
	value []Element
}

// NewElemArray は要素を作成します。
func NewElemArray() *ElemArray {
	return &ElemArray{value: []Element{}}
}

// Value は値を返します。
func (e *ElemArray) Value() interface{} {
	ret := make([]interface{}, 0)
	for _, val := range e.value {
		ret = append(ret, val.Value())
	}
	return interface{}(ret)
}

// Length は子供の数を返します。
func (e *ElemArray) Length() int {
	return len(e.value)
}

// Child は子供の要素を返します。
func (e *ElemArray) Child(num int) Element {
	if num < 0 || e.Length() <= num {
		return NewElemNull()
	}
	return e.value[num]
}

// Append は要素を追加します。
func (e *ElemArray) Append(child ...Element) *ElemArray {
	e.value = append(e.value, child...)
	return e
}

// Paths は要素の一覧を取得します。
func (e *ElemArray) Paths() []PathJSON {
	arr := []PathJSON{}
	for i, e := range e.value {
		key := PathJSON("[" + strconv.Itoa(i) + "]")
		arr = append(arr, key)
		for _, sub := range e.Paths() {
			wk := key + jsonPathSeparater + sub
			arr = append(arr, wk)
		}
	}
	return arr
}

// Type は要素の型を取得します。
func (e *ElemArray) Type() ElementType {
	return TypeArray
}

// Text は文字列表現を返します。
func (e *ElemArray) Text() string {
	arr := []string{}
	for _, value := range e.value {
		arr = append(arr, value.String())
	}
	ret := "["
	ret += strings.Join(arr, ", ")
	ret += "]"
	return ret
}

// AsObject は ElemObject にキャストします。
func (e *ElemArray) AsObject() (*ElemObject, bool) {
	return nil, false
}

// AsArray は ElemArray にキャストします。
func (e *ElemArray) AsArray() (*ElemArray, bool) {
	return e, true
}

// AsString は ElemString にキャストします。
func (e *ElemArray) AsString() (*ElemString, bool) {
	return nil, false
}

// AsFloat は ElemFloat にキャストします。
func (e *ElemArray) AsFloat() (*ElemFloat, bool) {
	return nil, false
}

// AsBool は ElemBool にキャストします。
func (e *ElemArray) AsBool() (*ElemBool, bool) {
	return nil, false
}

func (e ElemArray) String() string {
	return e.Text()
}
