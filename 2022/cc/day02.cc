#include "utils/file_reader.h"
#include <string>
#include <vector>
#include <iostream>

enum Play
{
  kRock,
  kPaper,
  kScissors,
  kInvalid,
};

Play getWinner(Play pos)
{
  switch (pos)
  {
  case kRock:
    return kPaper;
  case kPaper:
    return kScissors;
  case kScissors:
    return kRock;
  default:
    return pos;
  }
}

Play getLoser(Play pos)
{
  switch (pos)
  {
  case kRock:
    return kScissors;
  case kPaper:
    return kRock;
  case kScissors:
    return kPaper;
  default:
    return pos;
  }
}

Play translateOpp(char play)
{
  if (play == 'A')
  {
    return kRock;
  }
  if (play == 'B')
  {
    return kPaper;
  }
  if (play == 'C')
  {
    return kScissors;
  }
  return kInvalid;
}

Play translateYou_part1(char play)
{
  if (play == 'X')
  {
    return kRock;
  }
  if (play == 'Y')
  {
    return kPaper;
  }
  if (play == 'Z')
  {
    return kScissors;
  }
  return kInvalid;
}

Play translateYou_part2(Play opp, char you)
{
  if (you == 'Y')
  {
    return opp;
  }
  if (you == 'X')
  {
    return getLoser(opp);
  }
  return getWinner(opp);
}

int pointsForGame(Play opp, Play you)
{
  if (opp == you)
  {
    return 3;
  }
  if (opp == kRock && you == kPaper)
  {
    return 6;
  }
  if (opp == kPaper && you == kScissors)
  {
    return 6;
  }
  if (opp == kScissors && you == kRock)
  {
    return 6;
  }
  return 0;
}

int pointsForPlay(Play p)
{
  switch (p)
  {
  case kRock:
    return 1;
  case kPaper:
    return 2;
  case kScissors:
    return 3;
  default:
    return 0;
  }
}

int part1_score(char opp, char you)
{
  auto opp_play = translateOpp(opp);
  auto you_play = translateYou_part1(you);

  return pointsForPlay(you_play) + pointsForGame(opp_play, you_play);
}

int part2_score(char opp, char you)
{
  auto opp_play = translateOpp(opp);
  auto you_play = translateYou_part2(opp_play, you);

  return pointsForPlay(you_play) + pointsForGame(opp_play, you_play);
}

int main(int argc, char **argv)
{
  auto input = brooks::utils::read(argv[1]);

  int p1 = 0;
  int p2 = 0;
  for (auto line : input)
  {
    p1 += part1_score(line[0], line[2]);
    p2 += part2_score(line[0], line[2]);
  }

  std::cout << "Part 1: " << p1 << std::endl;
  std::cout << "Part 2: " << p2 << std::endl;

  return 0;
}
