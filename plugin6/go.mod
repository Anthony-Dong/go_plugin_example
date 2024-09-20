module github.com/anthony-dong/go_plugin_example/plugin6

go 1.18

require github.com/anthony-dong/go_plugin_example/databus v0.0.0

require google.golang.org/protobuf v1.34.2 // indirect

replace github.com/anthony-dong/go_plugin_example/databus => ../databus
