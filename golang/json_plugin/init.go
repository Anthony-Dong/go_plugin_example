package json_plugin

import (
	"github.com/tidwall/gjson"
	"plugin"
)

type GetJsonRowFunc func(string, string) string

func LoadGetJsonRowFunc(lib string) GetJsonRowFunc {
	p, err := plugin.Open(lib)
	if err != nil {
		panic(err)
	}
	foo, err := p.Lookup("GetJsonRow")
	if err != nil {
		panic(err)
	}
	return foo.(func(string, string) string)
}

func GetJsonRow(input string, path string) string {
	return gjson.Get(input, path).Raw
}
