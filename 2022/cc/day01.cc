#include "utils/file_reader.h"
#include <string>
#include <vector>
#include <iostream>

std::vector<int> largest(std::vector<std::string> input, int topN)
{
  std::vector<int> results;

  int working = 0;
  for (auto l : input)
  {
    if (l == "")
    {
      results.push_back(working);
      std::sort(results.begin(), results.end(), std::greater<int>());
      while (results.size() > topN)
      {
        results.pop_back();
      }
      working = 0;
      continue;
    }
    working += std::stoi(l);
  }
  return results;
}

int sum(std::vector<int> input)
{
  int total = 0;
  for (auto l : input)
  {
    total += l;
  }
  return total;
}

int main(int argc, char **argv)
{
  auto input = brooks::utils::read(argv[1]);

  auto p1 = largest(input, 1);
  std::cout << "Part 1: Max: " << p1[0] << "; Total: " << sum(p1) << std::endl;

  auto p2 = largest(input, 3);
  std::cout << "Part 2: Max: " << p2[0] << "; Total: " << sum(p2) << std::endl;

  return 0;
}
