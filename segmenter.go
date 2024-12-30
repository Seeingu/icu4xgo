package icu4xgo

//#include <icu4xgo.h>
//#include <stdlib.h>
//#include <string.h>
import "C"
import "runtime"

type GraphemeSegmenter struct {
	ptr    *C.IGGraphemeClusterSegmenter
	source string
	index  int
}

func NewGraphemeSegmenter(source string) *GraphemeSegmenter {
	s := &GraphemeSegmenter{
		ptr:    C.ig_init_grapheme_segmenter(),
		source: source,
	}
	C.ig_init_grapheme_iterator_utf8(s.ptr, C.CString(s.source))
	runtime.SetFinalizer(s, (*GraphemeSegmenter).free)
	return s
}

type GraphemeSegmenterNextResult struct {
	Segment string
	Index   int
}

func (s *GraphemeSegmenter) Next() GraphemeSegmenterNextResult {
	newIndex := C.ig_grapheme_iterator_next(s.ptr)
	grapheme := s.source[s.index:newIndex]
	s.index = int(newIndex)
	return GraphemeSegmenterNextResult{
		Segment: grapheme,
		Index:   s.index,
	}
}

func (s *GraphemeSegmenter) free() {
	C.ig_free_grapheme_segmenter(s.ptr)
}

type WordSegmenter struct {
	ptr    *C.IGWordSegmenter
	source string
	index  int
}

func NewWordSegmenter(source string) *WordSegmenter {
	s := &WordSegmenter{
		ptr:    C.ig_init_word_segmenter(),
		source: source,
	}
	C.ig_init_word_iterator_utf8(s.ptr, C.CString(s.source))
	runtime.SetFinalizer(s, (*WordSegmenter).free)
	return s
}

type WordSegmenterNextResult struct {
	Segment    string
	Index      int
	IsWordLike bool
}

func (s *WordSegmenter) Next() WordSegmenterNextResult {
	newIndex := C.ig_word_iterator_next(s.ptr)
	word := s.source[s.index:newIndex]
	s.index = int(newIndex)
	return WordSegmenterNextResult{
		Segment:    word,
		Index:      s.index,
		IsWordLike: bool(C.ig_word_iterator_is_word_like(s.ptr)),
	}
}

func (s *WordSegmenter) free() {
	C.ig_free_word_segmenter(s.ptr)
}

type SentenceSegmenter struct {
	ptr    *C.IGSentenceSegmenter
	source string
	index  int
}

func NewSentenceSegmenter(source string) *SentenceSegmenter {
	s := &SentenceSegmenter{
		ptr:    C.ig_init_sentence_segmenter(),
		source: source,
	}
	C.ig_init_sentence_iterator_utf8(s.ptr, C.CString(s.source))
	runtime.SetFinalizer(s, (*SentenceSegmenter).free)
	return s
}

type SentenceSegmenterNextResult struct {
	Segment string
	Index   int
}

func (s *SentenceSegmenter) Next() SentenceSegmenterNextResult {
	newIndex := C.ig_sentence_iterator_next(s.ptr)
	sentence := s.source[s.index:newIndex]
	s.index = int(newIndex)
	return SentenceSegmenterNextResult{
		Segment: sentence,
		Index:   s.index,
	}
}

func (s *SentenceSegmenter) free() {
	C.ig_free_sentence_segmenter(s.ptr)
}
