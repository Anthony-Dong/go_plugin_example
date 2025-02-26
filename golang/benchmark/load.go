package benchmark

import (
	"fmt"
	"os"
	"path/filepath"
	"plugin"
)

func LoadAddFunc(lib string) func(x, y int) int {
	pwd, _ := os.Getwd()
	fmt.Println(filepath.Join(pwd, lib))
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

func LoadJsonDecodeFunc(lib string) func(input []byte, v interface{}) error {
	pwd, _ := os.Getwd()
	fmt.Println(filepath.Join(pwd, lib))
	p, err := plugin.Open(lib)
	if err != nil {
		panic(err)
	}
	foo, err := p.Lookup("JsonDecode")
	if err != nil {
		panic(err)
	}
	return foo.(func(input []byte, v interface{}) error)
}
