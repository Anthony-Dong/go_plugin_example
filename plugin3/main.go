package main

import (
	"context"
	"fmt"
)

// 全部都以接口的形式处理！！
// 这个接口也可以放到一个包里！！

type Request interface {
	GetHeader(key string) string
	GetBody() ([]byte, error)
}

type Response interface {
	SetHeader(key, value string)
	SetBody([]byte) error
}

type plugin struct {
}

func (*plugin) Init() error {
	return nil
}

func (*plugin) Close() error {
	return nil
}

func (*plugin) Handle(ctx context.Context, req interface{}, resp interface{}) error {
	request := req.(Request)
	response := resp.(Response)
	body, err := request.GetBody()
	if err != nil {
		return err
	}
	fmt.Printf("recevie http request body: %s\n", body)
	if err := response.SetBody(body); err != nil {
		return err
	}
	return nil
}

// NewPlugin 初始化插件，每个插件都需要定义此函数，且函数签名一致
func NewPlugin() interface{} {
	return &plugin{}
}

func main() {}
