# -----
# from https://github.com/unicode-org/icu4x/blob/main/tutorials/c/Makefile
ICU_CAPI := $(shell cargo metadata --format-version 1 | jq '.packages[] | select(.name == "icu_capi").manifest_path' | xargs dirname)
HEADERS := ${ICU_CAPI}/bindings/c
ALL_HEADERS := $(wildcard ${HEADERS}/*)
# -----
LIB_PATH := ./lib
INCLUDE_PATH := icu4x

rustlib: Cargo.toml header
	cargo rustc -p icu_capi --crate-type staticlib --release
	cp ./target/release/libicu_capi.a ${LIB_PATH}

header: 
	cp -r ${HEADERS} ${INCLUDE_PATH}

build: *.go
	go build .

test: *.go
	go test ./...

all: rustlib build test

clean:
	rm -rf build
	rm -rf target