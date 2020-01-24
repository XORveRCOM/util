package json

// ElementType は要素の型を示す標識です。
type ElementType string

const (
	// TypeObject はオブジェクト型です。
	TypeObject = ElementType("object")
	// TypeArray は配列型です。
	TypeArray = ElementType("array")
	// TypeString は数値型です。
	TypeString = ElementType("string")
	// TypeFloat は数値型です。
	TypeFloat = ElementType("float")
	// TypeBool は論理値型です。
	TypeBool = ElementType("bool")
	// TypeNull はNULLです。
	TypeNull = ElementType("null")
)
