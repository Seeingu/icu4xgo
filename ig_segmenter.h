#ifndef ICUGO_SEGMENTER_H
#define ICUGO_SEGMENTER_H

#include "./icu4x/WordSegmenter.h"
#include "./icu4x/SentenceSegmenter.h"
#include "./icu4x/DataProvider.h"
#include "./icu4x/WordBreakIteratorUtf8.h"
#include "./icu4x/WordBreakIteratorUtf16.h"
#include "./icu4x/WordBreakIteratorLatin1.h"
#include "./icu4x/SentenceBreakIteratorUtf8.h"
#include "./icu4x/SentenceBreakIteratorUtf16.h"
#include "./icu4x/SentenceBreakIteratorLatin1.h"

#include <stdlib.h>

typedef struct IGWordSegmenter
{
    WordSegmenter *segmenter;
    bool is_ok;
    union
    {
        WordBreakIteratorUtf8 *utf8;
        WordBreakIteratorUtf16 *utf16;
        WordBreakIteratorLatin1 *latin1;
    } iterator;
} IGWordSegmenter;

IGWordSegmenter *ig_init_word_segmenter();

void *ig_init_word_iterator_utf8(IGWordSegmenter *ws, const char *s);

int ig_word_iterator_next(IGWordSegmenter *ws);

bool ig_word_iterator_is_word_like(IGWordSegmenter *ws);

void ig_free_word_segmenter(IGWordSegmenter *ws);

typedef struct IGSentenceSegmenter
{
    SentenceSegmenter *segmenter;
    bool is_ok;
    union
    {
        SentenceBreakIteratorUtf8 *utf8;
        SentenceBreakIteratorUtf16 *utf16;
        SentenceBreakIteratorLatin1 *latin1;
    } iterator;
} IGSentenceSegmenter;

IGSentenceSegmenter *ig_init_sentence_segmenter();

void *ig_init_sentence_iterator_utf8(IGSentenceSegmenter *ss, const char *s);

int ig_sentence_iterator_next(IGSentenceSegmenter *ss);

void ig_free_sentence_segmenter(IGSentenceSegmenter *ws);

#endif // ICUGO_SEGMENTER_H