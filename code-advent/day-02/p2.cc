#include <iostream>
#include <string>

int main(int argc, char** argv) {
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

  std::cout <<
    "Depth: " << depth <<
    "\tDistance: " << horizontal <<
    "\tProduct: " << (depth * horizontal) <<
    std::endl;
  return 0;
}

