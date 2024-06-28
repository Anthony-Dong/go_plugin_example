#include <benchmark/benchmark.h>

#include <stdexcept>


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

BENCHMARK(BM_Native);
// Run the benchmark
BENCHMARK_MAIN();