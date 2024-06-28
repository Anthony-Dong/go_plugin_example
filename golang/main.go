package main

import (
	"context"
	"fmt"
	"github.com/anthony-dong/cgo_demo/golang/http_plugin"
	"github.com/anthony-dong/cgo_demo/golang/json_plugin"
	"net/http"
)

func testHTTPPlugin() {
	lib := http_plugin.Load("plugin3/output/plugin.so") // 动态加载
	if err := lib.Init(); err != nil {
		panic(err)
	}
	defer lib.Close()
	if err := http.ListenAndServe(":8080", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ctx := context.Background()
		if err := lib.Handle(ctx, http_plugin.NewHTTPRequest(request), http_plugin.NewHTTPResponse(writer)); err != nil {
			fmt.Printf("[ERROR] %s\n", err.Error())
		}
	})); err != nil {
		panic(err)
	}
}

func testJsonPlugin() {
	v := json_plugin.LoadGetJsonRowFunc("plugin2/output/plugin.so")(`{"k1":"v1"}`, "k1")
	fmt.Println(v)
	r := json_plugin.LoadGetJsonPathFunc("plugin2/output/plugin.so")(`{"k1":"v1"}`, "k1")
	fmt.Println(r)
}

func main() {
	testHTTPPlugin()
}
