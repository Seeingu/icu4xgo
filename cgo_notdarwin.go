//go:build !darwin

package icu4xgo

//#cgo LDFLAGS: -L./lib -licu_capi -lm -ldl
import "C"
