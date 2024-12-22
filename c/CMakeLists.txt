cmake_minimum_required(VERSION 3.10)

project(icu4xgo)

set(SRC_FILES
    locale.c
)

add_library(icu4xgo STATIC ${SRC_FILES})
target_include_directories(icu4xgo PUBLIC include)

find_library(ICU_CAPI_LIBRARY NAMES icu_capi PATHS ../target/release)

if(ICU_CAPI_LIBRARY)
    target_link_libraries(icu4xgo PUBLIC ${ICU_CAPI_LIBRARY})
else()
    message(FATAL_ERROR "icu_capi library not found in ../target/release")
endif()

install(TARGETS icu4xgo DESTINATION lib)
install(DIRECTORY include/ DESTINATION include)
file(GLOB HEADER_FILES "*.h")
install(FILES ${HEADER_FILES} DESTINATION include)
