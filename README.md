# icu4xgo

This is a Go port of the ICU4X project. It includes a C wrapper for ICU4X's CFFI binding, which provides internationalization capabilities for software applications.

## Features

## Install

Prerequisites

- [Cargo](https://doc.rust-lang.org/cargo/getting-started/installation.html): We'll use Cargo to build C static library

```bash
go get github.com/Seeingu/icu4xgo@0.1.1

# If you encounter permission issues, run with sudo
cd `go list -f "{{.Dir}}" github.com/Seeingu/icu4xgo` && make rustlib
```

### Build from source

2. Build and run tests

```sh
make all
```

## Usage

## Contributing

## License

MIT

## Contact
