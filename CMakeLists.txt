cmake_minimum_required(VERSION 3.10)

add_subdirectory(c)

project(icu4xgo)

install(FILES ./target/release/libicu_capi.a DESTINATION lib)