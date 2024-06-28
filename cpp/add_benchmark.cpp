#include <benchmark/benchmark.h>

#include <stdexcept>

#include "plugin.h"

static void BM_CGO(benchmark::State& state) {
  // Perform setup here
  for (auto _ : state) {
    // This code gets timed
    auto x = int(Add(GoInt(1), GoInt(2)));
    if (x != 3) {
      throw std::runtime_error("异常");
    }
  }
}

static void BM_Native(benchmark::State& state) {
  // Perform setup here
  for (auto _ : state) {
    // This code gets timed
    auto x = 1 + 2;
    if (x != 3) {
      throw std::runtime_error("异常");
    }
  }
}

BENCHMARK(BM_CGO);
BENCHMARK(BM_Native);
// Run the benchmark
BENCHMARK_MAIN();