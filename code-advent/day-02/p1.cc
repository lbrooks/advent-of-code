#include <iostream>
#include <string>

int main(int argc, char** argv) {
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

  std::cout <<
    "Depth: " << depth <<
    "\tDistance: " << horizontal <<
    "\tProduct: " << (depth * horizontal) <<
    std::endl;
  return 0;
}

