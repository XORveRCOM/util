package json

import "strings"

// PathJSON は json の唯一の要素を示す識別です。
// XPATH の簡略版で、':' をセパレータとしてキーとインデックスを並べて記述します。
type PathJSON = string

const (
	jsonPathSeparater = ":"
)

// JoinJSONPath はパス配列を結合してパス文字列を生成します。
func JoinJSONPath(paths []string) PathJSON {
	return PathJSON(strings.Join(paths, jsonPathSeparater))
}

// SplitJSONPath はパス文字列を分解してパス配列を生成します。
func SplitJSONPath(path PathJSON) []string {
	return strings.Split(string(path), jsonPathSeparater)
}

// QueryElemObject は path から ElemObject を問い合わせます。
func QueryElemObject(e Element, path PathJSON) (*ElemObject, bool) {
	el, ok := Query(e, path).(*ElemObject)
	return el, ok
}

// QueryElemArray は path から ElemArray を問い合わせます。
func QueryElemArray(e Element, path PathJSON) (*ElemArray, bool) {
	el, ok := Query(e, path).(*ElemArray)
	return el, ok
}

// QueryElemString は path から ElemString を問い合わせます。
func QueryElemString(e Element, path PathJSON) (*ElemString, bool) {
	el, ok := Query(e, path).(*ElemString)
	return el, ok
}

// QueryElemFloat は path から ElemFloat を問い合わせます。
func QueryElemFloat(e Element, path PathJSON) (*ElemFloat, bool) {
	el, ok := Query(e, path).(*ElemFloat)
	return el, ok
}

// QueryElemBool は path から ElemBool を問い合わせます。
func QueryElemBool(e Element, path PathJSON) (*ElemBool, bool) {
	el, ok := Query(e, path).(*ElemBool)
	return el, ok
}
