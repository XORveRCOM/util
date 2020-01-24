package config_test

import (
	"testing"

	"github.com/xorvercom/util/pkg/config"
	testutil "github.com/xorvercom/util/pkg/testing"
)

func TestSetting(t *testing.T) {
	is, e := config.LoadFrom(testutil.GetFilepath("../../test/setting.json"))
	if e != nil {
		t.Fatalf("%+v", e)
	}
	for _, i := range is {
		t.Logf("%+v", i)
		//		s, ok := i.(map[string]interface{})
	}
}
