#include <iostream>
#include <string>
#include <vector>

void part1() {
  std::cout << "Not Yet Implemented" << std::endl;
}

void part2() {
  std::cout << "Not Yet Implemented" << std::endl;
}

int main(int argc, char** argv) {
  int buffer = 1;
	if (argc > 1) {
    buffer = std::atoi(argv[1]);
	}

  if (buffer == 1) {
    part1();
  } else if (buffer == 2) {
    part2();
  }

  return 0;
}

