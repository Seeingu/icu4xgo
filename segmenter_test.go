package icu4xgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSegmenter(t *testing.T) {
	t.Run("GraphemeSegmenter", func(t *testing.T) {
		s := NewGraphemeSegmenter("Hä!Wo")
		segments := []SegmenterNextResult{
			{
				Segment: "H",
				Index:   1,
			},
			{
				Segment: "ä",
				Index:   3,
			},
			{
				Segment: "!",
				Index:   4,
			},
			{
				Segment: "W",
				Index:   5,
			},
			{
				Segment: "o",
				Index:   6,
			},
		}
		for _, seg := range segments {
			assert.Equal(t, seg, s.Next())
		}
	})

	t.Run("WordSegmenter", func(t *testing.T) {
		s := NewWordSegmenter("Hello, world!")
		segments := []SegmenterNextResult{
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
	})
	t.Run("CSentenceSegmenter", func(t *testing.T) {
		s := NewSentenceSegmenter("Hello, world! How are you?")
		segments := []SegmenterNextResult{
			{
				Segment: "Hello, world! ",
				Index:   14,
			},
			{
				Segment: "How are you?",
				Index:   26,
			},
		}
		for _, seg := range segments {
			assert.Equal(t, seg, s.Next())
		}
	})
}
