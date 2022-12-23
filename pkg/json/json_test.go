package json_test

import (
	"testing"

	"github.com/xorvercom/util/pkg/json"
	testutil "github.com/xorvercom/util/pkg/testing"
)

func TestSettingLoad(t *testing.T) {
	_, e := json.LoadFromJSONFile(testutil.GetFilepath("../../test/setting.json"))
	if e != nil {
		t.Fatalf("%+v", e)
	}
	// 存在しないjsonを開く
	_, e = json.LoadFromJSONFile(testutil.GetFilepath("../../test/setting_noexists.json"))
	if e == nil {
		t.Fatalf("%+v", e)
	}
	// 空のjsonを開く
	_, e = json.LoadFromJSONByte([]byte{})
	if e == nil {
		t.Fatalf("%+v", e)
	}
}
func TestSettingParse(t *testing.T) {
	elem, e := json.LoadFromJSONFile(testutil.GetFilepath("../../test/setting.json"))
	if e != nil {
		t.Fatalf("%+v", e)
	}
	t.Logf("%+v", elem)
	// json文字列
	jsonstr := json.ToJSON(elem, false)
	t.Log(jsonstr)
	elem, e = json.LoadFromJSONByte([]byte(jsonstr))
	if e != nil {
		t.Fatalf("%+v", e)
	}
	t.Logf("%+v", elem)
}
