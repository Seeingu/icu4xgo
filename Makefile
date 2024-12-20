# -----
# from https://github.com/unicode-org/icu4x/blob/main/tutorials/c/Makefile
ICU_CAPI := $(shell cargo metadata --format-version 1 | jq '.packages[] | select(.name == "icu_capi").manifest_path' | xargs dirname)
HEADERS := ${ICU_CAPI}/bindings/c
ALL_HEADERS := $(wildcard ${HEADERS}/*)
# -----

rustlib: Cargo.toml
	cargo rustc -p icu_capi --crate-type staticlib --release

header: 
	cp -r ${HEADERS} ./c/include

clib: c/* rustlib header 
	cmake -B build
	cd ./build && make

build: *.go clib
	go build .

run: *.go
	go run .

clean:
	rm $(OUT_FILES)