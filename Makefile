.PHONY: all golang_benchmark golang_build golang_run cpp_lib cpp cpp_benchmark_run cpp_benchmark_build cpp_benchmark

DIR=$(shell pwd)

golang_benchmark:
	cd plugin && bash -e build.sh
	cd golang && CGO_ENABLED=1 go test -v -run=none -bench=Benchmark -count=2 -benchmem ./benchmark/...

golang_build: # 编译
	cd plugin1 && bash -e build.sh
	cd plugin2 && bash -e build.sh
	cd plugin3 && bash -e build.sh
	cd golang && CGO_ENABLED=1 go build -o output/main

golang_run:
	golang/output/main

cpp_lib:
	cd plugin4 && bash -e build.sh
	cd plugin5 && bash -e build.sh

cpp: cpp_lib
	mkdir -p output
	clang++ -std=c++20 -Wall -I$(DIR)/plugin4/output \
	-I$(DIR)/plugin4/output  -L$(DIR)/plugin4/output -lplugin \
	-o output/main cpp/main.cpp
	LD_LIBRARY_PATH=$(DIR)/plugin4/output DYLD_LIBRARY_PATH=$(DIR)/plugin4/output output/main

cpp_benchmark_build: cpp_lib ## https://github.com/google/benchmark
	clang++ -Wall -O2 -std=c++20 \
	-I$(DIR)/plugin5/output -I/usr/local/include \
	-c ./cpp/add_benchmark.cpp \
	-o output/add_benchmark.o
	clang++ -o output/add_benchmark output/add_benchmark.o \
	-L/usr/local/lib -lpthread -lbenchmark \
	-L$(DIR)/plugin5/output -lplugin

cpp_benchmark_run:
	DYLD_LIBRARY_PATH=$(DIR)/plugin4/output output/add_benchmark

cpp_benchmark: cpp_benchmark_build cpp_benchmark_run

clean:
	rm -rf */output/*