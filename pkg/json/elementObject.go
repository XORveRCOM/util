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
	dic map[string]Element
}

// NewElemObject は要素を作成します。
func NewElemObject() ElemObject {
	return &elemObject{dic: map[string]Element{}}
}

func (e *elemObject) String() string {
	return e.Text()
}

// Text は文字列表現を返します。
func (e *elemObject) Text() string {
	arr := []string{}
	for key, item := range e.dic {
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
	for key, item := range e.dic {
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
	for key, item := range e.dic {
		ret[key] = item.Value()
	}
	return interface{}(ret)
}

// AsObject は ElemObject にキャストします。
func (e *elemObject) AsObject() (ElemObject, bool) {
	return e, true
}

// AsArray は ElemArray にキャストします。
func (e *elemObject) AsArray() (ElemArray, bool) {
	return nil, false
}

// AsString は ElemString にキャストします。
func (e *elemObject) AsString() (ElemString, bool) {
	return nil, false
}

// AsFloat は ElemFloat にキャストします。
func (e *elemObject) AsFloat() (ElemFloat, bool) {
	return nil, false
}

// AsBool は ElemBool にキャストします。
func (e *elemObject) AsBool() (ElemBool, bool) {
	return nil, false
}

// IsObject は ElemObject であるか判定します。
func (e *elemObject) IsObject() bool {
	return true
}

// IsArray は ElemArray であるか判定します。
func (e *elemObject) IsArray() bool {
	return false
}

// AsString は ElemString であるか判定します。
func (e *elemObject) IsString() bool {
	return false
}

// IsFloat は ElemFloat であるか判定します。
func (e *elemObject) IsFloat() bool {
	return false
}

// IsBool は ElemBool であるか判定します。
func (e *elemObject) IsBool() bool {
	return false
}

// IsNull は ElemNull であるか判定します。
func (e *elemObject) IsNull() bool {
	return false
}

// Keys はキーの一覧をソートして返します。
func (e *elemObject) Keys() []string {
	ret := []string{}
	for key := range e.dic {
		ret = append(ret, key)
	}
	// ソート
	sort.Strings(ret)
	return ret
}

// Child は子供の要素を返します。
func (e *elemObject) Child(key string) Element {
	eval, ok := e.dic[key]
	if !ok {
		return NewElemNull()
	}
	return eval
}

// Put は要素を追加します。
func (e *elemObject) Put(key string, child Element) ElemObject {
	e.dic[key] = child
	return e
}
