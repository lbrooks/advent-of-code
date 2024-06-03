#include "utils/stream_reader.h"

#include <iostream>
#include <string>
#include <vector>

namespace brooks {
namespace utils {

std::vector<std::string> readLines() {
  std::vector<std::string> input;

  std::string line;
  while (std::cin >> line) {
    input.push_back(line);
  }

  return input;
}

}  // namespace utils
}  // namespace brooks