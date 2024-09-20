package databus

import (
	"encoding/json"
	"errors"
	"fmt"

	"google.golang.org/protobuf/proto"
)

var IsNilErr = errors.New("nil error")

var (
	ProtocMarshal = func(v interface{}) ([]byte, error) {
		message, isOK := v.(proto.Message)
		if !isOK {
			return nil, fmt.Errorf(`data [%T] is not proto message`, v)
		}
		return proto.Marshal(message)
	}
	ProtobufUnMarshal = func(data []byte, v interface{}) error {
		message, isOK := v.(proto.Message)
		if !isOK {
			return fmt.Errorf(`data [%T] is not proto message`, v)
		}
		return proto.Unmarshal(data, message)
	}
)

type HostDataBus struct {
	bus DataBus
}

func (h *HostDataBus) GetDataBus() DataBus {
	return h.bus
}

func NewHostDataBus() *HostDataBus {
	return &HostDataBus{bus: &databus{}}
}

func NewHostDataBusV2(bus interface{}) (*HostDataBus, error) {
	dataBus, isOk := bus.(DataBus)
	if !isOk {
		return nil, fmt.Errorf(`dataBus [%T] is not DataBus`, bus)
	}
	return &HostDataBus{bus: dataBus}, nil
}

// Get HostDataBus
// 如果在插件中那么可能 proto.Message 就不是 google.golang.org/protobuf/proto.Message 了，导致失败，所以这里需要使用代理
func (h *HostDataBus) Get(key string, dst interface{}) error {
	switch dst.(type) {
	case proto.Message:
		return h.bus.Get(key, dst, ProtobufUnMarshal)
	default:
		return h.bus.Get(key, dst, json.Unmarshal)
	}
}

// Set HostDataBus
// 原因也是同上
func (h *HostDataBus) Set(key string, dst interface{}) {
	switch dst.(type) {
	case proto.Message:
		h.bus.Set(key, dst, ProtocMarshal)
	default:
		h.bus.Set(key, dst, json.Marshal)
	}
}

func Ptr[T any](input T) *T {
	return &input
}
