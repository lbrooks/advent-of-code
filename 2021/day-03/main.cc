#include <iostream>
#include <string>
#include <vector>

void part1() {
  int counts[12] = {};
  int rowsVisited = 0;

  std::string line;
  while (std::cin >> line) {
    rowsVisited++;

    for (int i = 0; i < line.length(); i++) {
      if (line[i] == '1') {
        counts[i] += 1;
      }
    }
  }

  int gamma = 0;
  int epsilon = 0;
  for (int i = 0; i < 12; i++) {
    gamma = gamma << 1;
    epsilon = epsilon << 1;
    if (counts[i] > (rowsVisited - counts[i])) {
      gamma += 1;
    } else {
      epsilon += 1;
    }
  }

  std::cout << std::endl
            << "Gamma: " << gamma << "\tEpsilon: " << epsilon
            << "\tProduct: " << (gamma * epsilon) << std::endl;
}

void part2() { std::cout << "Not Yet Implemented" << std::endl; }

int main(int argc, char** argv) {
  int buffer = 1;
  if (argc > 1) {
    buffer = std::atoi(argv[1]);
  }

  if (buffer == 1) {
    part1();
  } else if (buffer == 2) {
    part2();
  } else {
    std::cout << "No valid part number selected." << std::endl;
  }

  return 0;
}
