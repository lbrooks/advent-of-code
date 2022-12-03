#include "utils/file_reader.h"

#include <fstream>
#include <iostream>
#include <string>
#include <vector>

namespace brooks
{
  namespace utils
  {

    std::vector<std::string> read(std::string filename)
    {
      std::vector<std::string> input;

      std::ifstream infile(filename);

      std::string line;
      while (std::getline(infile, line))
      {
        input.push_back(line);
      }

      return input;
    }

  }
}