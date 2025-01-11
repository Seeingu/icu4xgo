package icu4xgo

import (
	"github.com/clipperhouse/uax29/graphemes"
	"github.com/clipperhouse/uax29/iterators"
)

type GraphemeSegmenter struct {
	g *iterators.Segmenter
}

func NewGraphemeSegmenter(source string) *GraphemeSegmenter {
	s := graphemes.NewSegmenter([]byte(source))
	gs := &GraphemeSegmenter{
		g: s,
	}
	return gs
}

func (s *GraphemeSegmenter) Next() SegmenterNextResult {
	if s.g.Next() {
		return SegmenterNextResult{
			Segment: s.g.Text(),
			Index:   s.g.End(),
		}
	}
	return SegmenterNextResult{}
}

type WordSegmenter interface {
	Next() SegmenterNextResult
}

type SentenceSegmenter interface {
	Next() SegmenterNextResult
}

type SegmenterNextResult struct {
	Segment string
	Index   int
	// IsWordLike is true only if the segment is a word-like segment.
	IsWordLike bool
}
