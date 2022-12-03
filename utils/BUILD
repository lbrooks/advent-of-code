load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@rules_cc//cc:defs.bzl", "cc_library")

go_library(
    name = "stream_reader_go",
    srcs = ["stream_reader.go"],
    importpath = "aoc/utils",
    visibility = ["//visibility:public"],
)

cc_library(
    name = "stream_reader_cc",
    srcs = ["stream_reader.cc"],
    hdrs = ["stream_reader.h"],
    visibility = ["//visibility:public"],
)

cc_library(
    name = "file_reader_cc",
    srcs = ["file_reader.cc"],
    hdrs = ["file_reader.h"],
    visibility = ["//visibility:public"],
)
