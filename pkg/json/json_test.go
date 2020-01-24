package json

import (
	"testing"

	testutil "github.com/xorvercom/util/pkg/testing"
)

func TestSettingLoad(t *testing.T) {
	_, e := LoadFromJSONFile(testutil.GetFilepath("../../test/setting.json"))
	if e != nil {
		t.Fatalf("%+v", e)
	}
	// 存在しないjsonを開く
	_, e = LoadFromJSONFile(testutil.GetFilepath("../../test/setting_noexists.json"))
	if e == nil {
		t.Fatalf("%+v", e)
	}
	// 空のjsonを開く
	_, e = LoadFromJSONByte([]byte{})
	if e == nil {
		t.Fatalf("%+v", e)
	}
}
func TestSettingParse(t *testing.T) {
	elem, e := LoadFromJSONFile(testutil.GetFilepath("../../test/setting.json"))
	if e != nil {
		t.Fatalf("%+v", e)
	}
	t.Logf("%+v", elem)
	// json文字列
	json := ToJSON(elem, false)
	t.Log(json)
	elem, e = LoadFromJSONByte([]byte(json))
	if e != nil {
		t.Fatalf("%+v", e)
	}
	t.Logf("%+v", elem)
}
