package json

import (
	"sort"
	"strings"
)

// ElemObject はオブジェクト要素です
type ElemObject struct {
	value map[string]Element
}

// NewElemObject は要素を作成します。
func NewElemObject() *ElemObject {
	return &ElemObject{value: map[string]Element{}}
}

// Value はインタフェースとしての内容を取得します。
func (e *ElemObject) Value() interface{} {
	ret := make(map[string]interface{})
	for key, val := range e.value {
		ret[key] = val.Value()
	}
	return interface{}(ret)
}

// Keys はキーの一覧をソートして返します。
func (e *ElemObject) Keys() []string {
	ret := []string{}
	for key := range e.value {
		ret = append(ret, key)
	}
	// ソート
	sort.Strings(ret)
	return ret
}

// Child は子供の要素を返します。
func (e *ElemObject) Child(key string) Element {
	eval, ok := e.value[key]
	if !ok {
		return NewElemNull()
	}
	return eval
}

// Put は要素を追加します。
func (e *ElemObject) Put(key string, child Element) *ElemObject {
	e.value[key] = child
	return e
}

// Paths は要素の一覧をソートして取得します。
func (e *ElemObject) Paths() []PathJSON {
	arr := []PathJSON{}
	for key, value := range e.value {
		key := PathJSON(key)
		arr = append(arr, key)
		for _, sub := range value.Paths() {
			wk := key + jsonPathSeparater + sub
			arr = append(arr, wk)
		}
	}
	// ソート
	sort.Strings(arr)
	return arr
}

// Type は要素の型を取得します。
func (e *ElemObject) Type() ElementType {
	return TypeObject
}

// Text は文字列表現を返します。
func (e *ElemObject) Text() string {
	arr := []string{}
	for key, value := range e.value {
		wk := "\"" + key + "\":" + value.String()
		arr = append(arr, wk)
	}
	// ソート
	sort.Strings(arr)

	ret := "{"
	ret += strings.Join(arr, ", ")
	ret += "}"
	return ret
}

// AsObject は ElemObject にキャストします。
func (e *ElemObject) AsObject() (*ElemObject, bool) {
	return e, true
}

// AsArray は ElemArray にキャストします。
func (e *ElemObject) AsArray() (*ElemArray, bool) {
	return nil, false
}

// AsString は ElemString にキャストします。
func (e *ElemObject) AsString() (*ElemString, bool) {
	return nil, false
}

// AsFloat は ElemFloat にキャストします。
func (e *ElemObject) AsFloat() (*ElemFloat, bool) {
	return nil, false
}

// AsBool は ElemBool にキャストします。
func (e *ElemObject) AsBool() (*ElemBool, bool) {
	return nil, false
}

func (e ElemObject) String() string {
	return e.Text()
}
