package main

import "encoding/json"

func Add(x, y int) int {
	return x + y
}

func JsonDecode(input []byte, v interface{}) error {
	return json.Unmarshal(input, v)
}

func main() {}
