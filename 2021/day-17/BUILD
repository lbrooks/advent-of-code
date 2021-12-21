load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "source",
    srcs = [
        "main.go",
    ],
    importpath = "github.com/lbrooks/adventofcode/2021/day17"
)

go_binary(
    name = "go",
    embed = [
        ":source",
    ],
)
