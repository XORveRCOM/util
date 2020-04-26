package json

import (
	"strconv"
)

// Query は要素を問い合わせます。
// オブジェクトのキーもしくは配列のインデックスを ':' で結合したパスにある要素を取得します。
// 要するに xpath もどきです。
// JSON ではキーに数字が使えるように読める(RFC8259)ので、配列インデックスは[]で囲って記述します。
func Query(root Element, query PathJSON) Element {
	path := SplitJSONPath(query)
	cur := root
	for _, elem := range path {
		if len(elem) > 2 && elem[0] == '[' && elem[len(elem)-1] == ']' {
			// 配列
			num, err := strconv.Atoi(elem[1 : len(elem)-1])
			if err != nil {
				// 配列インデックスが数字ではなかった
				return NewElemNull()
			}
			switch val := cur.(type) {
			case ElemArray:
				cur = val.Child(num)
			default:
				return NewElemNull()
			}
		} else {
			// オブジェクト
			switch val := cur.(type) {
			case ElemObject:
				cur = val.Child(elem)
			default:
				return NewElemNull()
			}
		}
	}
	return cur
}
