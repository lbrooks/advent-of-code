#include <iostream>
#include <string>
#include <vector>

#include "utils/file_reader.h"

int main(int argc, char **argv) {
  auto input = brooks::utils::read(argv[1]);

  for (auto line : input) {
    std::cout << "Lin4: " << line << std::endl;
  }

  return 0;
}
