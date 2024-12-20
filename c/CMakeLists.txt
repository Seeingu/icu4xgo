cmake_minimum_required(VERSION 3.10)

project(icu4xgo)

set(SRC_FILES
    locale.c
)

add_library(icu4xgo STATIC ${SRC_FILES})
target_include_directories(icu4xgo PUBLIC include)
