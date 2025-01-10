package icu4xgo

//#import "segmenter_darwin.h"
import "C"

import (
	"runtime"
	"unsafe"
)

type WordSegmenterDarwin struct {
	WordSegmenter
	ptr     unsafe.Pointer
	source  string
	cSource unsafe.Pointer
	index   int
}

type SentenceSegmenterDarwin struct {
	SentenceSegmenter
	ptr     unsafe.Pointer
	source  string
	cSource unsafe.Pointer
	index   int
}

func NewWordSegmenter(source string) WordSegmenter {
	cSource := unsafe.Pointer(
		C.C_CFString_Create(C.CString(source)),
	)
	tokenizer := C.C_CFStringTokenizer_CreateWord(cSource)
	s := &WordSegmenterDarwin{
		ptr:    tokenizer,
		source: source,
	}
	s.cSource = cSource
	runtime.SetFinalizer(s, (*WordSegmenterDarwin).free)
	return s
}

func (s *WordSegmenterDarwin) Next() SegmenterNextResult {
	result := C.CFStringTokenizerRef_Next(s.ptr)
	if result.position == -1 {
		return SegmenterNextResult{}
	}
	newIndex := result.position
	word := s.source[s.index:newIndex]
	s.index = int(newIndex)
	return SegmenterNextResult{
		Segment:    word,
		Index:      s.index,
		IsWordLike: bool(!result.hasNonLetters),
	}
}

func (s *WordSegmenterDarwin) free() {
	C.C_CFString_Release(s.cSource)
	C.C_CFStringTokenizer_Release(s.ptr)
}

func NewSentenceSegmenter(source string) SentenceSegmenter {
	cSource := unsafe.Pointer(
		C.C_CFString_Create(C.CString(source)),
	)
	tokenizer := C.C_CFStringTokenizer_CreateSentence(cSource)
	s := &SentenceSegmenterDarwin{
		ptr:    tokenizer,
		source: source,
	}
	s.cSource = cSource
	runtime.SetFinalizer(s, (*SentenceSegmenterDarwin).free)
	return s
}

func (s *SentenceSegmenterDarwin) Next() SegmenterNextResult {
	result := C.CFStringTokenizerRef_Next(s.ptr)
	if result.position == -1 {
		return SegmenterNextResult{}
	}
	newIndex := result.position
	sentence := s.source[s.index:newIndex]
	s.index = int(newIndex)
	return SegmenterNextResult{
		Segment: sentence,
		Index:   s.index,
	}
}

func (s *SentenceSegmenterDarwin) free() {
	C.C_CFString_Release(s.cSource)
	C.C_CFStringTokenizer_Release(s.ptr)
}
