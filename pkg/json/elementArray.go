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
	// Delete は要素を削除します。
	Delete(num int)
}

// ElemArray は配列型要素です
type elemArray struct {
	element
	value []Element
}

// NewElemArray は要素を作成します。
func NewElemArray() ElemArray {
	return &elemArray{value: []Element{}}
}

// JSON 表現として {...} や [...] や "string" というような文字列を返します。
func (e elemArray) String() string {
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

// AsArray は ElemArray にキャストします。
func (e *elemArray) AsArray() (ElemArray, bool) {
	return e, true
}

// IsArray は ElemArray であるか判定します。
func (e *elemArray) IsArray() bool {
	return true
}

// Clone はディープコピーした Element を返します。
func (e *elemArray) Clone() Element {
	arr := NewElemArray()
	for _, elem := range e.value {
		arr.Append(elem.Clone())
	}
	return arr
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

// Delete は要素を削除します。
func (e *elemArray) Delete(num int) {
	s := e.Size()
	switch {
	case num < 0 || s <= num:
		return
	case num == 0:
		e.value = e.value[1:]
	case s == num+1:
		e.value = e.value[:num+1]
	default:
		e.value = append(e.value[:num], e.value[num+1:]...)
	}
}
