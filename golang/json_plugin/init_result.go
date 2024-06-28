package json_plugin

import (
	"github.com/tidwall/gjson"
	"plugin"
)

type GetJsonPathFunc func(string, string) gjson.Result

func GetJsonPath(input string, path string) gjson.Result {
	return gjson.Get(input, path)
}

func LoadGetJsonPathFunc(lib string) GetJsonPathFunc {
	p, err := plugin.Open(lib)
	if err != nil {
		panic(err)
	}
	foo, err := p.Lookup("GetJsonPath")
	if err != nil {
		panic(err)
	}
	return foo.(func(string, string) gjson.Result)
}
