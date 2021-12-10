# [Advent of Code](https://adventofcode.com/)
* All solutions expect the input to be passed through via stdin.
* All solutions accept one flag to determine the question's part

> Commands listed below assume this is the current working directory

## To Run:
```cat <year>/day-<day>/input.txt | bazel run //<year>/day-<day>:go -- <part>```

The hashmarks are replaced with the day and question part respectivly.

Ex: `cat 2021/day-01/input.txt | bazel run //2021/day-01:go -- 1` will run the solution to part 1 of the first day's question of 2021
