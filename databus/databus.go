package databus

import (
	"errors"
	"reflect"
	"sync"
)

type DataBus interface {
	Set(key string, value interface{}, marshal func(v interface{}) ([]byte, error))
	Get(key string, dst interface{}, unmarshal func(data []byte, dst interface{}) error) error
}

type Value struct {
	Data    interface{}
	Marshal func(v interface{}) ([]byte, error)
}

type databus struct {
	data map[string]*Value

	lock *sync.RWMutex
}

func (d *databus) Set(key string, value interface{}, marshal func(v interface{}) ([]byte, error)) {
	data := Value{Data: value, Marshal: marshal}
	if d.lock != nil {
		d.lock.Lock()
		defer d.lock.Unlock()
	}
	if d.data == nil {
		d.data = make(map[string]*Value)
	}
	d.data[key] = &data
}

func (d *databus) get(key string) (*Value, bool) {
	if d.lock != nil {
		d.lock.RLock()
		defer d.lock.RUnlock()
	}
	srcValue, isExist := d.data[key]
	return srcValue, isExist
}

func (d *databus) Get(key string, dst interface{}, unmarshal func(data []byte, dst interface{}) error) error {
	src, isExist := d.get(key)
	if !isExist {
		return IsNilErr
	}
	if reflect.TypeOf(dst).Kind() != reflect.Ptr {
		return errors.New("dst value must be a pointer")
	}
	if d.copyValue(src.Data, dst) {
		return nil
	}
	return d.cloneValue(src, dst, unmarshal)
}

func (d *databus) cloneValue(src *Value, dst interface{}, unmarshal func(data []byte, dst interface{}) error) error {
	marshal, err := src.Marshal(src.Data)
	if err != nil {
		return err
	}
	if err := unmarshal(marshal, dst); err != nil {
		return err
	}
	return nil
}

func (*databus) copyValue(src interface{}, dst interface{}) bool {
	dstVal := reflect.ValueOf(dst)
	srcVal := reflect.ValueOf(src)
	srcType := srcVal.Type()

	// dst type contains src type
	// dst src
	// **a *a
	// *a  a
	for dstVal.Kind() == reflect.Ptr {
		dstVal = dstVal.Elem()
		if dstVal.Type() == srcType {
			dstVal.Set(srcVal)
			return true
		}
	}
	// dst src
	// *a *a
	for srcVal.Kind() == reflect.Ptr {
		srcVal = srcVal.Elem()
	}
	if dstVal.Type() == srcVal.Type() {
		dstVal.Set(srcVal)
		return true
	}
	return false
}
