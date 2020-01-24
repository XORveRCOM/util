package json

import (
	"os"
	fpath "path/filepath"
	"strings"
	"testing"
)

func TestElement(t *testing.T) {
	temp := fpath.Join(os.TempDir(), "config_test")
	os.MkdirAll(temp, os.ModeDir)
	defer os.RemoveAll(temp)
	elem := NewElemObject()
	aelem := NewElemArray()
	aelem.Append(NewElemString("abc"), NewElemBool(true), NewElemFloat(3.14), NewElemNull())
	elem.Put("abc", aelem)
	t.Log(ToJSON(elem, true))
	t.Log(ToJSON(elem, false))

	// ファイルに保存
	json1 := ToJSON(elem, true)
	filename := fpath.Join(temp, "test.json")
	if e := SaveToJSONFile(filename, elem, false); e != nil {
		t.Fatal(e)
	}
	elem2, e := LoadFromJSONFile(filename)
	if e != nil {
		t.Fatal(e)
	}
	json2 := ToJSON(elem2, true)
	if json1 != json2 {
		t.Fatalf("%s != %s", json1, json2)
	}
}

func TestElementObject(t *testing.T) {
	e := NewElemObject()
	e.Put("b", NewElemString("1"))
	e.Put("c", NewElemString("2"))
	e.Put("a", NewElemString("3"))
	keys := e.Keys()
	if strings.Join(keys, ",") != "a,b,c" {
		t.Fatal(keys)
	}
}

func TestElementFloat(t *testing.T) {
	f := 3.14
	e := NewElemFloat(f)
	if e.Float() != f {
		t.Fatal()
	}
}

func TestElementBool(t *testing.T) {
	f := true
	e := NewElemBool(f)
	if e.Bool() != f {
		t.Fatal()
	}
}

func TestAs(t *testing.T) {
	var elem Element
	elem = NewElemObject()
	if _, ok := elem.AsObject(); false == ok {
		t.Failed()
	}
	if _, ok := elem.AsArray(); ok {
		t.Failed()
	}
	if _, ok := elem.AsString(); ok {
		t.Failed()
	}
	if _, ok := elem.AsFloat(); ok {
		t.Failed()
	}
	if _, ok := elem.AsBool(); ok {
		t.Failed()
	}

	elem = NewElemArray()
	if _, ok := elem.AsObject(); ok {
		t.Failed()
	}
	if _, ok := elem.AsArray(); false == ok {
		t.Failed()
	}
	if _, ok := elem.AsString(); ok {
		t.Failed()
	}
	if _, ok := elem.AsFloat(); ok {
		t.Failed()
	}
	if _, ok := elem.AsBool(); ok {
		t.Failed()
	}

	elem = NewElemString("")
	if _, ok := elem.AsObject(); ok {
		t.Failed()
	}
	if _, ok := elem.AsArray(); ok {
		t.Failed()
	}
	if _, ok := elem.AsString(); false == ok {
		t.Failed()
	}
	if _, ok := elem.AsFloat(); ok {
		t.Failed()
	}
	if _, ok := elem.AsBool(); ok {
		t.Failed()
	}

	elem = NewElemFloat(1)
	if _, ok := elem.AsObject(); ok {
		t.Failed()
	}
	if _, ok := elem.AsArray(); ok {
		t.Failed()
	}
	if _, ok := elem.AsString(); ok {
		t.Failed()
	}
	if _, ok := elem.AsFloat(); false == ok {
		t.Failed()
	}
	if _, ok := elem.AsBool(); ok {
		t.Failed()
	}

	elem = NewElemBool(true)
	if _, ok := elem.AsObject(); ok {
		t.Failed()
	}
	if _, ok := elem.AsArray(); ok {
		t.Failed()
	}
	if _, ok := elem.AsString(); ok {
		t.Failed()
	}
	if _, ok := elem.AsFloat(); ok {
		t.Failed()
	}
	if _, ok := elem.AsBool(); false == ok {
		t.Failed()
	}

	elem = NewElemNull()
	if _, ok := elem.AsObject(); ok {
		t.Failed()
	}
	if _, ok := elem.AsArray(); ok {
		t.Failed()
	}
	if _, ok := elem.AsString(); ok {
		t.Failed()
	}
	if _, ok := elem.AsFloat(); ok {
		t.Failed()
	}
	if _, ok := elem.AsBool(); false == ok {
		t.Failed()
	}

}
