package json

import (
	"testing"

	testutil "github.com/xorvercom/util/pkg/testing"
)

func TestPath(t *testing.T) {
	elem, e := LoadFromJSONFile(testutil.GetFilepath("../../test/setting.json"))
	if e != nil {
		t.Fatalf("%+v", e)
	}
	t.Logf("%+v", elem)

	// パスアクセス
	for i, p := range elem.Paths() {
		target := Query(elem, p)
		t.Logf("[%d] path:\"%s\", type:%s, value:%+v", i, p, target.Type(), target)
	}

	var target Element
	// パスアクセス異常
	s := JoinJSONPath([]string{})
	target = Query(elem, s)
	switch v := target.(type) {
	case ElemNull:
		null := v.Value()
		if null != nil {
			t.Fatal(v.Type())
		}
	default:
		t.Fatal(v.Type())
	}
	if _, ok := QueryElemObject(elem, s); ok {
		t.Fatal()
	}
	if _, ok := QueryElemArray(elem, s); ok {
		t.Fatal()
	}
	if _, ok := QueryElemString(elem, s); ok {
		t.Fatal()
	}
	if _, ok := QueryElemFloat(elem, s); ok {
		t.Fatal()
	}
	if _, ok := QueryElemBool(elem, s); ok {
		t.Fatal()
	}

	s = JoinJSONPath([]string{""})
	target = Query(elem, s)
	switch v := target.(type) {
	case ElemNull:
		if target.Type() != TypeNull {
			t.Fatal(target.Type())
		}
	default:
		t.Fatal(v.Type())
	}

	s = JoinJSONPath([]string{"notexist"})
	target = Query(elem, s)
	switch v := target.(type) {
	case ElemNull:
		if target.Type() != TypeNull {
			t.Fatal(target.Type())
		}
	default:
		t.Fatal(v.Type())
	}

	s = JoinJSONPath([]string{"[0]"})
	target = Query(elem, s)
	switch v := target.(type) {
	case ElemNull:
		if target.Type() != TypeNull {
			t.Fatal(target.Type())
		}
	default:
		t.Fatal(v.Type())
	}

	s = JoinJSONPath([]string{"arr", "object"})
	target = Query(elem, s)
	switch v := target.(type) {
	case ElemNull:
		if target.Type() != TypeNull {
			t.Fatal(target.Type())
		}
	default:
		t.Fatal(v.Type())
	}

	s = JoinJSONPath([]string{"arr", "[2]"})
	target = Query(elem, s)
	switch v := target.(type) {
	case ElemNull:
		if target.Type() != TypeNull {
			t.Fatal(target.Type())
		}
	default:
		t.Fatal(v.Type())
	}

	s = JoinJSONPath([]string{"arr", "[-1]"})
	target = Query(elem, s)
	switch v := target.(type) {
	case ElemNull:
		if target.Type() != TypeNull {
			t.Fatal(target.Type())
		}
	default:
		t.Fatal(v.Type())
	}

	s = JoinJSONPath([]string{"arr", "[string]"})
	target = Query(elem, s)
	switch v := target.(type) {
	case ElemNull:
		if target.Type() != TypeNull {
			t.Fatal(target.Type())
		}
	default:
		t.Fatal(v.Type())
	}
}
