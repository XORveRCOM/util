package json

import (
	"os"
	fpath "path/filepath"
	"strings"
	"testing"
)

func TestElement(t *testing.T) {
	temp := fpath.Join(os.TempDir(), "config_test")
	os.MkdirAll(temp, os.ModeDir) // nolint
	defer os.RemoveAll(temp)
	elem := NewElemObject()
	aelem := NewElemArray()
	aelem.Append(NewElemString("abc"), NewElemBool(true), NewElemFloat(3.14), NewElemNull())
	elem.Put("abc", aelem)
	t.Log(ToJSON(elem, true))
	t.Log(ToJSON(elem, false))

	// ファイルに保存
	json1 := ToJSON(elem.Clone(), true)
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

func TestPaths(t *testing.T) {
	var elem Element
	var paths []PathJSON

	elem = NewElemObject()
	paths = elem.Paths()
	if len(paths) != 0 {
		t.Failed()
	}

	elem = NewElemArray()
	paths = elem.Paths()
	if len(paths) != 0 {
		t.Failed()
	}

	elem = NewElemBool(true)
	paths = elem.Paths()
	if len(paths) != 0 {
		t.Failed()
	}

	elem = NewElemFloat(0)
	paths = elem.Paths()
	if len(paths) != 0 {
		t.Failed()
	}

	elem = NewElemNull()
	paths = elem.Paths()
	if len(paths) != 0 {
		t.Failed()
	}

	elem = NewElemString("")
	paths = elem.Paths()
	if len(paths) != 0 {
		t.Failed()
	}
}

func TestType(t *testing.T) {
	var elem Element

	elem = NewElemObject()
	if elem.Type() != TypeObject {
		t.Failed()
	}

	elem = NewElemArray()
	if elem.Type() != TypeArray {
		t.Failed()
	}

	elem = NewElemBool(true)
	if elem.Type() != TypeBool {
		t.Failed()
	}

	elem = NewElemFloat(0)
	if elem.Type() != TypeFloat {
		t.Failed()
	}

	elem = NewElemNull()
	if elem.Type() != TypeNull {
		t.Failed()
	}

	elem = NewElemString("")
	if elem.Type() != TypeString {
		t.Failed()
	}
}
func TestAs(t *testing.T) {
	var elem Element

	elem = NewElemObject()
	if e, ok := elem.AsObject(); false == ok {
		t.Failed()
	} else {
		if e.String() == "" {
			t.Failed()
		}
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
	if e, ok := elem.AsArray(); false == ok {
		t.Failed()
	} else {
		if e.String() == "" {
			t.Failed()
		}
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
	if e, ok := elem.AsString(); false == ok {
		t.Failed()
	} else {
		if e.String() == "" {
			t.Failed()
		}
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
	if e, ok := elem.AsFloat(); false == ok {
		t.Failed()
	} else {
		if e.String() == "" {
			t.Failed()
		}
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
	if e, ok := elem.AsBool(); false == ok {
		t.Failed()
	} else {
		if e.String() == "" {
			t.Failed()
		}
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
	if e, ok := elem.AsBool(); false == ok {
		t.Failed()
	} else {
		if e.String() == "" {
			t.Failed()
		}
	}
}

func TestIs(t *testing.T) {
	var elem Element

	elem = NewElemObject()
	if false == elem.IsObject() {
		t.Failed()
	}
	if elem.IsArray() {
		t.Failed()
	}
	if elem.IsString() {
		t.Failed()
	}
	if elem.IsString() {
		t.Failed()
	}
	if elem.IsFloat() {
		t.Failed()
	}
	if elem.IsBool() {
		t.Failed()
	}
	if elem.IsNull() {
		t.Failed()
	}

	elem = NewElemArray()
	if elem.IsObject() {
		t.Failed()
	}
	if false == elem.IsArray() {
		t.Failed()
	}
	if elem.IsString() {
		t.Failed()
	}
	if elem.IsString() {
		t.Failed()
	}
	if elem.IsFloat() {
		t.Failed()
	}
	if elem.IsBool() {
		t.Failed()
	}
	if elem.IsNull() {
		t.Failed()
	}

	elem = NewElemString("")
	if elem.IsObject() {
		t.Failed()
	}
	if elem.IsArray() {
		t.Failed()
	}
	if elem.IsString() {
		t.Failed()
	}
	if false == elem.IsString() {
		t.Failed()
	}
	if elem.IsFloat() {
		t.Failed()
	}
	if elem.IsBool() {
		t.Failed()
	}
	if elem.IsNull() {
		t.Failed()
	}

	elem = NewElemFloat(1)
	if elem.IsObject() {
		t.Failed()
	}
	if elem.IsArray() {
		t.Failed()
	}
	if elem.IsString() {
		t.Failed()
	}
	if elem.IsString() {
		t.Failed()
	}
	if false == elem.IsFloat() {
		t.Failed()
	}
	if elem.IsBool() {
		t.Failed()
	}
	if elem.IsNull() {
		t.Failed()
	}

	elem = NewElemBool(true)
	if elem.IsObject() {
		t.Failed()
	}
	if elem.IsArray() {
		t.Failed()
	}
	if elem.IsString() {
		t.Failed()
	}
	if elem.IsString() {
		t.Failed()
	}
	if elem.IsFloat() {
		t.Failed()
	}
	if false == elem.IsBool() {
		t.Failed()
	}
	if elem.IsNull() {
		t.Failed()
	}

	elem = NewElemNull()
	if elem.IsObject() {
		t.Failed()
	}
	if elem.IsArray() {
		t.Failed()
	}
	if elem.IsString() {
		t.Failed()
	}
	if elem.IsString() {
		t.Failed()
	}
	if elem.IsFloat() {
		t.Failed()
	}
	if elem.IsBool() {
		t.Failed()
	}
	if false == elem.IsNull() {
		t.Failed()
	}
}

func TestLength(t *testing.T) {
	var elem ElemString
	var str string

	str = ""
	elem = NewElemString(str)
	if elem.Length() != len(str) {
		t.Failed()
	}

	str = "1234"
	elem = NewElemString(str)
	if elem.Length() != len(str) {
		t.Failed()
	}
}
