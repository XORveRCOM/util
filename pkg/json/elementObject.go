package json

import (
	"sort"
	"strings"
)

type ElemObject interface {
	Element
	// Keys はキーの一覧をソートして返します。
	Keys() []string
	// Child は子供の要素を返します。
	Child(key string) Element
	// Put は要素を追加します。
	Put(key string, child Element) ElemObject
}

// ElemObject はオブジェクト要素です
type elemObject struct {
	element
	value map[string]Element
}

// NewElemObject は要素を作成します。
func NewElemObject() ElemObject {
	return &elemObject{value: map[string]Element{}}
}

// JSON 表現として {...} や [...] や "string" というような文字列を返します。
func (e *elemObject) String() string {
	return e.Text()
}

// Text は文字列表現を返します。
func (e *elemObject) Text() string {
	arr := []string{}
	for key, item := range e.value {
		wk := "\"" + key + "\":" + item.String()
		arr = append(arr, wk)
	}
	// ソート
	sort.Strings(arr)

	ret := "{"
	ret += strings.Join(arr, ", ")
	ret += "}"
	return ret
}

// Type は要素の型を取得します。
func (e *elemObject) Type() ElementType {
	return TypeObject
}

// Paths は子供のパス要素の一覧を取得します。
func (e *elemObject) Paths() []PathJSON {
	arr := []PathJSON{}
	for key, item := range e.value {
		key := PathJSON(key)
		arr = append(arr, key)
		for _, path := range item.Paths() {
			wk := key + jsonPathSeparater + path
			arr = append(arr, wk)
		}
	}
	// ソート
	sort.Strings(arr)
	return arr
}

// Value はインタフェースとしての内容を取得します。
func (e *elemObject) Value() interface{} {
	ret := make(map[string]interface{})
	for key, item := range e.value {
		ret[key] = item.Value()
	}
	return interface{}(ret)
}

// AsObject は ElemObject にキャストします。
func (e *elemObject) AsObject() (ElemObject, bool) {
	return e, true
}

// IsObject は ElemObject であるか判定します。
func (e *elemObject) IsObject() bool {
	return true
}

// Clone はディープコピーした Element を返します。
func (e *elemObject) Clone() Element {
	arr := NewElemObject()
	for key, elem := range e.value {
		arr.Put(key, elem.Clone())
	}
	return arr
}

// Keys はキーの一覧をソートして返します。
func (e *elemObject) Keys() []string {
	ret := []string{}
	for key := range e.value {
		ret = append(ret, key)
	}
	// ソート
	sort.Strings(ret)
	return ret
}

// Child は子供の要素を返します。
func (e *elemObject) Child(key string) Element {
	eval, ok := e.value[key]
	if !ok {
		return NewElemNull()
	}
	return eval
}

// Put は要素を追加します。
func (e *elemObject) Put(key string, child Element) ElemObject {
	e.value[key] = child
	return e
}
