package icu4xgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSegmenter(t *testing.T) {
	s := NewSegmenter(WordSegmenter, "Hello, world!")
	segments := []SegmenterNextResult{
		{},
		{
			Segment:    "Hello",
			IsWordLike: true,
			Index:      5,
		},
		{
			Segment:    ",",
			IsWordLike: false,
			Index:      6,
		},
		{
			Segment:    " ",
			IsWordLike: false,
			Index:      7,
		},
		{
			Segment:    "world",
			IsWordLike: true,
			Index:      12,
		},
	}
	for _, seg := range segments {
		assert.Equal(t, seg, s.Next())
	}
}
