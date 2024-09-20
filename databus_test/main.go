package main

import (
	"fmt"
	"github.com/anthony-dong/go_plugin_example/databus"
	"github.com/anthony-dong/go_plugin_example/databus/model"
	"os"
	"plugin"
)

func LoadHandle(lib string) func(interface{}) error {
	p, err := plugin.Open(lib)
	if err != nil {
		panic(err)
	}
	foo, err := p.Lookup("Handle")
	if err != nil {
		panic(err)
	}
	return foo.(func(interface{}) error)
}

func main() {
	fmt.Println(os.Getpid())
	lib := os.Args[1]
	fmt.Println(lib)
	// 创建 databus
	bus := databus.NewHostDataBus()

	data := model.Player{Id: databus.Ptr(int32(1)), Name: databus.Ptr("111")}
	bus.Set("k1", &data)

	// 加载插件并且调用
	if err := LoadHandle(lib)(bus.GetDataBus()); err != nil {
		panic(err)
	}

	// 处理完成数据加载数据
	data2 := model.Player{}
	if err := bus.Get("k1", &data2); err != nil {
		panic(err)
	}
	fmt.Println(data2.String())
}
