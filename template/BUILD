load("@io_bazel_rules_go//go:def.bzl", "go_binary")
load("@rules_cc//cc:defs.bzl", "cc_binary")

go_binary(
    name = "dayXX_go",
    srcs = ["dayXX.go"],
    deps = ["//utils:stream_reader_go"],
)

cc_binary(
    name = "dayXX_cc",
    srcs = ["dayXX.cc"],
    deps = ["//utils:file_reader_cc"],
)
