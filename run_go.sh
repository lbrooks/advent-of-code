#!/opt/homebrew/bin/fish

cat $argv[1]/day$argv[2].txt | bazel run //$argv[1]/go:day$argv[2]
