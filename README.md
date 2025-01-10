# icu4xgo

Go Internalization library

When on macOS, it uses the Foundation framework. 
Otherwise, it uses the [ICU4X](https://github.com/unicode-org/icu4x) C FFI.

## Features

| Feature           | macOS | Others |
|-------------------|-------|--------|
| Locale            | [x]   | [x]    |
| NumberFormatter   | [x]   | [x]    |
| ListFormatter     | [ ]   | [x]    |
| PluralRules       | [ ]   | [x]    |
| DateTimeFormatter | [x]   | [ ]    |
| Collator          | [x]   | [x]    |
| Segmenter         | [x]   | [ ]    |

Issues:

macOS:

- [ ] join a list of items with ListInitType and ListLength

others:

- [ ] setTimezone not work
- [ ] Segmenter should not return empty at the start

## Installation

1. Get the package

```bash
# or use latest commit, e.g:
# go get github.com/Seeingu/icu4xgo@4ebc70a0d4977f56b26a2eb65d77267b2c57e8c4
go get github.com/Seeingu/icu4xgo@v0.1.3
```

2. Link the C static library(on non-macOS systems)

Prerequisites

- [Cargo](https://doc.rust-lang.org/cargo/getting-started/installation.html): We'll use Cargo to build C static library

```bash
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
