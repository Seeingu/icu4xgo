# -----
# from https://github.com/unicode-org/icu4x/blob/main/tutorials/c/Makefile
ICU_CAPI := $(shell cargo metadata --format-version 1 | jq '.packages[] | select(.name == "icu_capi").manifest_path' | xargs dirname)
HEADERS := ${ICU_CAPI}/bindings/c
ALL_HEADERS := $(wildcard ${HEADERS}/*)
# -----
LIB_PATH := /usr/local/lib
INCLUDE_PATH := /usr/local/include

rustlib: Cargo.toml
	cargo rustc -p icu_capi --crate-type staticlib --release

header: 
	cp -r ${HEADERS} ./c/include

clib: c/* rustlib header 
	cmake -B build
	cd ./build && make
	cd ./build && sudo make install

build: *.go clib
	go build .

test: *.go
	go test ./...

clean:
	rm -rf build
	rm -rf c/include
	rm -rf target