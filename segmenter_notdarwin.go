//go:build !darwin
// +build !darwin

package icu4xgo

import (
	"github.com/clipperhouse/uax29/iterators"
	"github.com/clipperhouse/uax29/iterators/filter"
	"github.com/clipperhouse/uax29/sentences"
	"github.com/clipperhouse/uax29/words"
)

type CWordSegmenter struct {
	g *iterators.Segmenter
}

func NewWordSegmenter(source string) *CWordSegmenter {
	s := &CWordSegmenter{
		g: words.NewSegmenter([]byte(source)),
	}
	return s
}

func (s *CWordSegmenter) Next() SegmenterNextResult {
	if s.g.Next() {
		return SegmenterNextResult{
			Segment:    s.g.Text(),
			Index:      s.g.End(),
			IsWordLike: filter.Wordlike(s.g.Bytes()),
		}
	}
	return SegmenterNextResult{}
}

type CSentenceSegmenter struct {
	g *iterators.Segmenter
}

func NewSentenceSegmenter(source string) *CSentenceSegmenter {
	s := &CSentenceSegmenter{
		g: sentences.NewSegmenter([]byte(source)),
	}
	return s
}

func (s *CSentenceSegmenter) Next() SegmenterNextResult {
	if s.g.Next() {
		return SegmenterNextResult{
			Segment: s.g.Text(),
			Index:   s.g.End(),
		}
	}
	return SegmenterNextResult{}
}
