// Package json は json ファイルを扱うパッケージです。
package json

import (
	"bytes"
	libjson "encoding/json"
	"fmt"
	"os"
)

// LoadFromJSONFile は設定ファイルから Element を作成します。
func LoadFromJSONFile(filename string) (Element, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("os.ReadFile(%s) return %v", filename, err)
	}
	return LoadFromJSONByte(data)
}

// SaveToJSONFile は Element を設定ファイルに出力します。
func SaveToJSONFile(filename string, elem Element, indentation bool) error {
	jsonstr := ToJSON(elem, indentation)
	return os.WriteFile(filename, []byte(jsonstr), 0755)
}

// ToJSON は Element を json 文字列に変換します。
func ToJSON(elem Element, indentation bool) string {
	if indentation {
		b, _ := libjson.Marshal(elem.Value())
		var out bytes.Buffer
		libjson.Indent(&out, b, "", "    ") // nolint
		return out.String()
	}
	return elem.String()
}

// LoadFromJSONByte はバイト列のjsonから Element を作成します。
func LoadFromJSONByte(data []byte) (Element, error) {
	var jsondata interface{}
	e := libjson.Unmarshal(data, &jsondata)
	if e != nil {
		return nil, fmt.Errorf("%v", e)
	}
	return Parse(jsondata), nil
}

// Parse は要素を返します。
func Parse(i interface{}) Element {
	switch val := i.(type) {
	case []interface{}:
		// 配列[]
		ret := NewElemArray()
		for _, value := range val {
			ret.Append(Parse(value))
		}
		return ret
	case map[string]interface{}:
		// オブジェクト{}
		ret := NewElemObject()
		for key, value := range val {
			ret.Put(key, Parse(value))
		}
		return ret
	case string:
		ret := NewElemString(val)
		return ret
	case float64:
		ret := NewElemFloat(float64(val))
		return ret
	case bool:
		ret := NewElemBool(val)
		return ret
	default:
		return NewElemNull()
	}
}
