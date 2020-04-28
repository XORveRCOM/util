package json

// Element は設定要素です。
type Element interface {
	// JSON 表現として {...} や [...] や "string" というような文字列を返します。
	String() string
	// Text は文字列表現を返します。
	// Stringer を JSON 表現にしている代替です。ElemString では内容としての文字列です。
	Text() string
	// Type は要素の種類を取得します。
	Type() ElementType
	// Paths は子供のパス要素の一覧を取得します。
	Paths() []PathJSON
	// Value はインタフェースとしての内容を取得します。
	Value() interface{}
	// AsObject は ElemObject にキャストします。
	AsObject() (ElemObject, bool)
	// IsObject は ElemObject であるか判定します。
	IsObject() bool
	// AsArray は ElemArray にキャストします。
	AsArray() (ElemArray, bool)
	// IsArray は ElemArray であるか判定します。
	IsArray() bool
	// AsString は ElemString にキャストします。
	AsString() (ElemString, bool)
	// AsString は ElemString であるか判定します。
	IsString() bool
	// AsFloat は ElemFloat にキャストします。
	AsFloat() (ElemFloat, bool)
	// IsFloat は ElemFloat であるか判定します。
	IsFloat() bool
	// AsBool は ElemBool にキャストします。
	AsBool() (ElemBool, bool)
	// IsBool は ElemBool であるか判定します。
	IsBool() bool
	// IsNull は ElemNull であるか判定します。
	IsNull() bool
}
