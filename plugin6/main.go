package main

import (
	"fmt"

	"github.com/anthony-dong/go_plugin_example/databus"
	"github.com/anthony-dong/go_plugin_example/databus/model"
)

func Handle(_bus interface{}) error {
	// 加载 databus (和宿主机是同一个 databus实例)
	bus, err := databus.NewHostDataBusV2(_bus)
	if err != nil {
		return err
	}
	// 加载数据
	data := model.Player{}
	if err := bus.Get("k1", &data); err != nil {
		return err
	}
	fmt.Println(data.String())

	data.Name = databus.Ptr("plugin6")

	// set 数据
	bus.Set("k1", &data)
	return nil
}

func main() {}
