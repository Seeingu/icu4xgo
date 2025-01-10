# icu4xgo

This is a Go binding to the [ICU4X](https://github.com/unicode-org/icu4x) C FFI

## Features

- [x] Get Locale properties
- [x] NumberFormatter
- [ ] ListFormatter: join a list of items with ListInitType and ListLength
- [ ] PluralRules
- [ ] DateTimeFormatter

## Installation

### macOS

```bash
# or use latest commit, e.g:
# go get github.com/Seeingu/icu4xgo@4ebc70a0d4977f56b26a2eb65d77267b2c57e8c4
go get github.com/Seeingu/icu4xgo@v0.1.3
```

### Other OS

Prerequisites

- [Cargo](https://doc.rust-lang.org/cargo/getting-started/installation.html): We'll use Cargo to build C static library

```bash
# or use latest commit, e.g:
# go get github.com/Seeingu/icu4xgo@4ebc70a0d4977f56b26a2eb65d77267b2c57e8c4
go get github.com/Seeingu/icu4xgo@v0.1.3

# If you encounter permission issues, run make with sudo
cd `go list -f "{{.Dir}}" github.com/Seeingu/icu4xgo` && make rustlib
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
	fmt.Println(locale.BaseName()) // zh-Hans-CN
	hc := locale.HourCycle()
	fmt.Println(hc) // h12
}
```

## Contributing

## License

MIT

## Contact
