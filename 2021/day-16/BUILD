load("@rules_cc//cc:defs.bzl", "cc_binary")
load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

cc_binary(
    name = "cpp",
    srcs = ["main.cc"],
)

go_library(
    name = "source",
    srcs = [
        "main.go",
        "bits.go",
        "bits_reader.go",
    ],
    importpath = "github.com/lbrooks/adventofcode/2021/day16"
)

go_binary(
    name = "go",
    embed = [
        ":source",
    ],
)

go_test(
    name = "go_test",
    srcs = [
      "main_test.go",
    ],
    embed = [
        ":source",
    ],
)