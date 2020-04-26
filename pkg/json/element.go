package json

// Element は設定要素です。
type Element interface {
	String() string
	// Text は文字列表現を返します。
	// Stringer を JSON 表現にしている代替です。
	Text() string
	// Type は要素の種類を取得します。
	Type() ElementType
	// Paths は子供のパス要素の一覧を取得します。
	Paths() []PathJSON
	// Value はインタフェースとしての内容を取得します。
	Value() interface{}
	// AsObject は ElemObject にキャストします。
	AsObject() (ElemObject, bool)
	// AsArray は ElemArray にキャストします。
	AsArray() (ElemArray, bool)
	// AsString は ElemString にキャストします。
	AsString() (ElemString, bool)
	// AsFloat は ElemFloat にキャストします。
	AsFloat() (ElemFloat, bool)
	// AsBool は ElemBool にキャストします。
	AsBool() (ElemBool, bool)
}
