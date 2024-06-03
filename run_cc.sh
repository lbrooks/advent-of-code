#!/opt/homebrew/bin/fish

bazel run //$argv[1]/cc:day$argv[2] -- (pwd)/$argv[1]/input/day$argv[2].txt
