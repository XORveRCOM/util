package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Setting は設定ファイルの中身です。
type Setting interface {
	// Keys はキーの一覧を返します。
	Keys() []string
	// Item はキーに対する値を返します。
	Item(string) string
}

type keyvalue struct {
	key   string
	value interface{}
}

func (kv *keyvalue) IsSlice() bool {
	_, ok := kv.value.([]keyvalue)
	return ok
}

type setting struct {
	data keyvalue
}

func (s *setting) Keys() []string {
	return nil
}
func (s *setting) Item(key string) string {
	return ""
}

// LoadFrom は設定ファイルから Setting を作成します。
func LoadFrom(filename string) ([]interface{}, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadFile(%s) return %v", filename, err)
	}

	var jsondata []interface{}
	e := json.Unmarshal(data, &jsondata)
	if e != nil {
		return nil, fmt.Errorf("%v", e)
	}
	return jsondata, nil
}
