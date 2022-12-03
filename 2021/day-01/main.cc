#include <iostream>
#include <string>

void process(int countBack) {
  int buffer[countBack];
  memset(buffer, 0, countBack * sizeof(int));

  int idxToReplace = 0;
  int rowsVisited = 0;
  int numOfIncreases = 0;
  int prev = 0;

  std::string line;
  while (std::cin >> line) {
    int current = std::stoi(line);
    rowsVisited++;

    if (rowsVisited > countBack) {
      if ((prev + current - buffer[idxToReplace]) > prev) {
        numOfIncreases++;
      }
    }

    prev = prev - buffer[idxToReplace] + current;
    buffer[idxToReplace] = current;
    idxToReplace = (idxToReplace + 1) % countBack;
  }

  std::cout << std::endl
            << "Total Increase Count: " << numOfIncreases << std::endl;
}

int main(int argc, char** argv) {
  int countBack = 1;
  if (argc > 1) {
    countBack = std::atoi(argv[1]);
  }

  if (countBack == 1) {
    process(1);
  } else if (countBack == 2) {
    process(3);
  }

  return 0;
}
