package main

import "github.com/tidwall/gjson"

func GetJsonRow(input string, path string) string {
	return gjson.Get(input, path).Raw
}

func GetJsonPath(input string, path string) gjson.Result {
	return gjson.Get(input, path)
}
func main() {}
