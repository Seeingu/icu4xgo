# icu4xgo

This is a Go binding to the [ICU4X](https://github.com/unicode-org/icu4x) C FFI

## Features

## Install

Prerequisites

- [Cargo](https://doc.rust-lang.org/cargo/getting-started/installation.html): We'll use Cargo to build C static library

```bash
# or use latest commit, e.g:
# go get github.com/Seeingu/icu4xgo@2188b6255e13106e0ee1725170037a0cf1777e8f
go get github.com/Seeingu/icu4xgo@v0.1.1

# If you encounter permission issues, run make with sudo
cd `go list -f "{{.Dir}}" github.com/Seeingu/icu4xgo` && make rustlib
```

### Optional: Build from source

2. Build and run tests

```sh
make all
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/Seeingu/icu4xgo"
)

func main() {
	locale := icu4xgo.NewLocale("zh-Hans-CN-u-ca-chinese-hc-h12")
	defer locale.Free()
	fmt.Println(locale.BaseName()) // zh-Hans-CN
	hc, _ := locale.HourCycle()
	fmt.Println(hc) // h12
}
```

## Contributing

## License

MIT

## Contact
