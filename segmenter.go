package icu4xgo

//#include <icu4xgo.h>
//#include <stdlib.h>
//#include <string.h>
import "C"
import "runtime"

type Segmenter struct {
	ptr    *C.IGWordSegmenter
	source string
	index  int
}

type SegmenterType int

const (
	WordSegmenter SegmenterType = iota
	SentenceSegmenter
)

func NewSegmenter(st SegmenterType, source string) *Segmenter {
	s := &Segmenter{
		ptr:    C.ig_init_word_segmenter(),
		source: source,
	}
	C.ig_init_word_iterator_utf8(s.ptr, C.CString(s.source))
	runtime.SetFinalizer(s, (*Segmenter).free)
	return s
}

type SegmenterNextResult struct {
	Segment    string
	Index      int
	IsWordLike bool
}

func (s *Segmenter) Next() SegmenterNextResult {
	newIndex := C.ig_word_iterator_next(s.ptr)
	word := s.source[s.index:newIndex]
	s.index = int(newIndex)
	return SegmenterNextResult{
		Segment:    word,
		Index:      s.index,
		IsWordLike: bool(C.ig_word_iterator_is_word_like(s.ptr)),
	}
}

func (s *Segmenter) free() {
	C.ig_free_word_segmenter(s.ptr)
}
