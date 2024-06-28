package benchmark

import "plugin"

func LoadAddFunc(lib string) func(x, y int) int {
	p, err := plugin.Open(lib)
	if err != nil {
		panic(err)
	}
	foo, err := p.Lookup("Add")
	if err != nil {
		panic(err)
	}
	return foo.(func(x, y int) int)
}
