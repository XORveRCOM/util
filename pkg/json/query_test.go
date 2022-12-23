package json_test

import (
	"testing"

	"github.com/xorvercom/util/pkg/json"
	testutil "github.com/xorvercom/util/pkg/testing"
)

func TestPath(t *testing.T) {
	elem, e := json.LoadFromJSONFile(testutil.GetFilepath("../../test/setting.json"))
	if e != nil {
		t.Fatalf("%+v", e)
	}
	t.Logf("%+v", elem)

	// パスアクセス
	for i, p := range elem.Paths() {
		target := json.Query(elem, p)
		t.Logf("[%d] path:\"%s\", type:%s, value:%+v", i, p, target.Type(), target)
	}

	var target json.Element
	// パスアクセス異常
	s := json.JoinJSONPath([]string{})
	target = json.Query(elem, s)
	switch v := target.(type) {
	case json.ElemNull:
		null := v.Value()
		if null != nil {
			t.Fatal(v.Type())
		}
	default:
		t.Fatal(v.Type())
	}
	if _, ok := json.QueryElemObject(elem, s); ok {
		t.Fatal()
	}
	if _, ok := json.QueryElemArray(elem, s); ok {
		t.Fatal()
	}
	if _, ok := json.QueryElemString(elem, s); ok {
		t.Fatal()
	}
	if _, ok := json.QueryElemFloat(elem, s); ok {
		t.Fatal()
	}
	if _, ok := json.QueryElemBool(elem, s); ok {
		t.Fatal()
	}

	s = json.JoinJSONPath([]string{""})
	target = json.Query(elem, s)
	switch v := target.(type) {
	case json.ElemNull:
		if target.Type() != json.TypeNull {
			t.Fatal(target.Type())
		}
	default:
		t.Fatal(v.Type())
	}

	s = json.JoinJSONPath([]string{"notexist"})
	target = json.Query(elem, s)
	switch v := target.(type) {
	case json.ElemNull:
		if target.Type() != json.TypeNull {
			t.Fatal(target.Type())
		}
	default:
		t.Fatal(v.Type())
	}

	s = json.JoinJSONPath([]string{"[0]"})
	target = json.Query(elem, s)
	switch v := target.(type) {
	case json.ElemNull:
		if target.Type() != json.TypeNull {
			t.Fatal(target.Type())
		}
	default:
		t.Fatal(v.Type())
	}

	s = json.JoinJSONPath([]string{"arr", "object"})
	target = json.Query(elem, s)
	switch v := target.(type) {
	case json.ElemNull:
		if target.Type() != json.TypeNull {
			t.Fatal(target.Type())
		}
	default:
		t.Fatal(v.Type())
	}

	s = json.JoinJSONPath([]string{"arr", "[2]"})
	target = json.Query(elem, s)
	switch v := target.(type) {
	case json.ElemNull:
		if target.Type() != json.TypeNull {
			t.Fatal(target.Type())
		}
	default:
		t.Fatal(v.Type())
	}

	s = json.JoinJSONPath([]string{"arr", "[-1]"})
	target = json.Query(elem, s)
	switch v := target.(type) {
	case json.ElemNull:
		if target.Type() != json.TypeNull {
			t.Fatal(target.Type())
		}
	default:
		t.Fatal(v.Type())
	}

	s = json.JoinJSONPath([]string{"arr", "[string]"})
	target = json.Query(elem, s)
	switch v := target.(type) {
	case json.ElemNull:
		if target.Type() != json.TypeNull {
			t.Fatal(target.Type())
		}
	default:
		t.Fatal(v.Type())
	}
}
