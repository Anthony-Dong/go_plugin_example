package http_plugin

import (
	"plugin"
)

func Load(lib string) Plugin {
	p, err := plugin.Open(lib)
	if err != nil {
		panic(err)
	}
	foo, err := p.Lookup("NewPlugin")
	if err != nil {
		panic(err)
	}
	return foo.(func() interface{})().(Plugin)
}
