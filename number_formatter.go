package icu4xgo

import "C"
import (
	"errors"
)

var NumberFormatFailedError = errors.New("number format failed")

type NumberFormatter interface {
	Format(value string) (string, error)
}
