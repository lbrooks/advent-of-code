#include <iostream>
#include <string>

void part1() {
  int depth = 0;
  int horizontal = 0;

  std::string line;
  while (std::cin >> line) {
    std::string direction = line;

    std::cin >> line;
    int distance = std::stoi(line);

    if (direction == "up") {
      depth = depth - distance;
    } else if (direction == "down") {
      depth = depth + distance;
    } else if (direction == "forward") {
      horizontal = horizontal + distance;
    }
  }

  std::cout << "Depth: " << depth << "\tDistance: " << horizontal
            << "\tProduct: " << (depth * horizontal) << std::endl;
}

void part2() {
  int depth = 0;
  int horizontal = 0;
  int aim = 0;

  std::string line;
  while (std::cin >> line) {
    std::string direction = line;

    std::cin >> line;
    int distance = std::stoi(line);

    if (direction == "up") {
      aim -= distance;
    } else if (direction == "down") {
      aim += distance;
    } else if (direction == "forward") {
      horizontal += distance;
      depth += (aim * distance);
    }
  }

  std::cout << "Depth: " << depth << "\tDistance: " << horizontal
            << "\tProduct: " << (depth * horizontal) << std::endl;
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
  } else {
    std::cout << "No valid part number selected." << std::endl;
  }

  return 0;
}
