//go:build !darwin
// +build !darwin

package icu4xgo

//#include <icu4xgo.h>
//#include <stdlib.h>
//#include <string.h>
import "C"
import "runtime"

type CGraphemeSegmenter struct {
	ptr    *C.IGGraphemeClusterSegmenter
	source string
	index  int
}

func NewGraphemeSegmenter(source string) *CGraphemeSegmenter {
	s := &CGraphemeSegmenter{
		ptr:    C.ig_init_grapheme_segmenter(),
		source: source,
	}
	C.ig_init_grapheme_iterator_utf8(s.ptr, C.CString(s.source))
	runtime.SetFinalizer(s, (*CGraphemeSegmenter).free)
	return s
}

func (s *CGraphemeSegmenter) Next() SegmenterNextResult {
	newIndex := C.ig_grapheme_iterator_next(s.ptr)
	grapheme := s.source[s.index:newIndex]
	s.index = int(newIndex)
	return SegmenterNextResult{
		Segment: grapheme,
		Index:   s.index,
	}
}

func (s *CGraphemeSegmenter) free() {
	C.ig_free_grapheme_segmenter(s.ptr)
}

type CWordSegmenter struct {
	ptr    *C.IGWordSegmenter
	source string
	index  int
}

func NewWordSegmenter(source string) *CWordSegmenter {
	s := &CWordSegmenter{
		ptr:    C.ig_init_word_segmenter(),
		source: source,
	}
	C.ig_init_word_iterator_utf8(s.ptr, C.CString(s.source))
	runtime.SetFinalizer(s, (*CWordSegmenter).free)
	return s
}

func (s *CWordSegmenter) Next() SegmenterNextResult {
	newIndex := C.ig_word_iterator_next(s.ptr)
	word := s.source[s.index:newIndex]
	s.index = int(newIndex)
	return SegmenterNextResult{
		Segment:    word,
		Index:      s.index,
		IsWordLike: bool(C.ig_word_iterator_is_word_like(s.ptr)),
	}
}

func (s *CWordSegmenter) free() {
	C.ig_free_word_segmenter(s.ptr)
}

type CSentenceSegmenter struct {
	ptr    *C.IGSentenceSegmenter
	source string
	index  int
}

func NewSentenceSegmenter(source string) *CSentenceSegmenter {
	s := &CSentenceSegmenter{
		ptr:    C.ig_init_sentence_segmenter(),
		source: source,
	}
	C.ig_init_sentence_iterator_utf8(s.ptr, C.CString(s.source))
	runtime.SetFinalizer(s, (*CSentenceSegmenter).free)
	return s
}

func (s *CSentenceSegmenter) Next() SegmenterNextResult {
	newIndex := C.ig_sentence_iterator_next(s.ptr)
	sentence := s.source[s.index:newIndex]
	s.index = int(newIndex)
	return SegmenterNextResult{
		Segment: sentence,
		Index:   s.index,
	}
}

func (s *CSentenceSegmenter) free() {
	C.ig_free_sentence_segmenter(s.ptr)
}
