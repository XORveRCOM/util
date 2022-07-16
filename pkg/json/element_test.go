package json_test

import (
	"os"
	fpath "path/filepath"
	"strings"
	"testing"

	"github.com/xorvercom/util/pkg/json"
)

func TestElement(t *testing.T) {
	temp := fpath.Join(os.TempDir(), "config_test")
	os.MkdirAll(temp, 0755) // nolint
	defer os.RemoveAll(temp)
	elem := json.NewElemObject()
	aelem := json.NewElemArray()
	aelem.Append(json.NewElemString("abc"), json.NewElemBool(true), json.NewElemFloat(3.14), json.NewElemNull())
	elem.Put("abc", aelem)
	t.Log(json.ToJSON(elem, true))
	t.Log(json.ToJSON(elem, false))

	// ファイルに保存
	json1 := json.ToJSON(elem.Clone(), true)
	filename := fpath.Join(temp, "test.json")
	if e := json.SaveToJSONFile(filename, elem, false); e != nil {
		t.Fatal(e)
	}
	elem2, e := json.LoadFromJSONFile(filename)
	if e != nil {
		t.Fatal(e)
	}
	json2 := json.ToJSON(elem2, true)
	if json1 != json2 {
		t.Fatalf("%s != %s", json1, json2)
	}
}

func TestElementObject(t *testing.T) {
	e := json.NewElemObject()
	e.Put("b", json.NewElemString("1"))
	e.Put("c", json.NewElemString("2"))
	e.Put("a", json.NewElemString("3"))
	keys := e.Keys()
	if strings.Join(keys, ",") != "a,b,c" {
		t.Fatal(keys)
	}
}

func TestElementFloat(t *testing.T) {
	f := 3.14
	e := json.NewElemFloat(f)
	if e.Float() != f {
		t.Fatal()
	}
}

func TestElementBool(t *testing.T) {
	f := true
	e := json.NewElemBool(f)
	if e.Bool() != f {
		t.Fatal()
	}
}

func TestPaths(t *testing.T) {
	var elem json.Element
	var paths []json.PathJSON

	elem = json.NewElemObject()
	paths = elem.Paths()
	if len(paths) != 0 {
		t.Failed()
	}

	elem = json.NewElemArray()
	paths = elem.Paths()
	if len(paths) != 0 {
		t.Failed()
	}

	elem = json.NewElemBool(true)
	paths = elem.Paths()
	if len(paths) != 0 {
		t.Failed()
	}

	elem = json.NewElemFloat(0)
	paths = elem.Paths()
	if len(paths) != 0 {
		t.Failed()
	}

	elem = json.NewElemNull()
	paths = elem.Paths()
	if len(paths) != 0 {
		t.Failed()
	}

	elem = json.NewElemString("")
	paths = elem.Paths()
	if len(paths) != 0 {
		t.Failed()
	}
}

func TestType(t *testing.T) {
	var elem json.Element

	elem = json.NewElemObject()
	if elem.Type() != json.TypeObject {
		t.Failed()
	}

	elem = json.NewElemArray()
	if elem.Type() != json.TypeArray {
		t.Failed()
	}

	elem = json.NewElemBool(true)
	if elem.Type() != json.TypeBool {
		t.Failed()
	}

	elem = json.NewElemFloat(0)
	if elem.Type() != json.TypeFloat {
		t.Failed()
	}

	elem = json.NewElemNull()
	if elem.Type() != json.TypeNull {
		t.Failed()
	}

	elem = json.NewElemString("")
	if elem.Type() != json.TypeString {
		t.Failed()
	}
}
func TestAs(t *testing.T) {
	var elem json.Element

	elem = json.NewElemObject()
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

	elem = json.NewElemArray()
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

	elem = json.NewElemString("")
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

	elem = json.NewElemFloat(1)
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

	elem = json.NewElemBool(true)
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

	elem = json.NewElemNull()
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
	var elem json.Element

	elem = json.NewElemObject()
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

	elem = json.NewElemArray()
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

	elem = json.NewElemString("")
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

	elem = json.NewElemFloat(1)
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

	elem = json.NewElemBool(true)
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

	elem = json.NewElemNull()
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
	var elem json.ElemString
	var str string

	str = ""
	elem = json.NewElemString(str)
	if elem.Length() != len(str) {
		t.Failed()
	}

	str = "1234"
	elem = json.NewElemString(str)
	if elem.Length() != len(str) {
		t.Failed()
	}
}
