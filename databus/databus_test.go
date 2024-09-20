package databus

import (
	"testing"
)

func TestNewDataBus(t *testing.T) {
	bus := NewHostDataBus()
	{
		bus.Set("k1", "v1")
		value := ""
		if err := bus.Get("k1", &value); err != nil {
			t.Fatal(err)
		}
		t.Log(value)
	}

	{
		type Data struct {
			Name string
		}
		data1 := Data{Name: "tom"}
		data2 := &Data{}
		bus.Set("k1", &data1)
		if err := bus.Get("k1", &data2); err != nil {
			t.Fatal(err)
		}
		data2.Name = "xiaoming"
		t.Log(data1)
		t.Log(data2)
	}
}

func BenchmarkDataBus(b *testing.B) {
	type Data struct {
		Name string
	}
	bus := NewHostDataBus()
	bus.Set("k1", &Data{Name: "tom"})
	for i := 0; i < b.N; i++ {
		var data2 *Data
		if err := bus.Get("k1", &data2); err != nil {
			b.Fatal(err)
		}
		data2.Name = "xiaoming"
	}
}
