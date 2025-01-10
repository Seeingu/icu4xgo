package icu4xgo

type GraphemeSegmenter interface {
	Next() SegmenterNextResult
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
