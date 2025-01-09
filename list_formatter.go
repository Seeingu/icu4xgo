package icu4xgo

type ListInitType int

const (
	ListAnd ListInitType = iota
	ListOr
	ListUnit
)

type ListLength int

const (
	ListLengthWide ListLength = iota
	ListLengthShort
	ListLengthNarrow
)

type ListFormatter interface {
	Format(items []string) string
}
