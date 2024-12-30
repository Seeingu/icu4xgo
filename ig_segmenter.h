#ifndef ICUGO_SEGMENTER_H
#define ICUGO_SEGMENTER_H

#include "./icu4x/WordSegmenter.h"
#include "./icu4x/SentenceSegmenter.h"
#include "./icu4x/DataProvider.h"
#include "./icu4x/WordBreakIteratorUtf8.h"
#include "./icu4x/WordBreakIteratorUtf16.h"
#include "./icu4x/WordBreakIteratorLatin1.h"

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

typedef struct IGSentenceSegmenter
{
    SentenceSegmenter *segmenter;
} IGSentenceSegmenter;

IGWordSegmenter *ig_init_word_segmenter();

void *ig_init_word_iterator_utf8(IGWordSegmenter *ws, const char *s);

int ig_word_iterator_next(IGWordSegmenter *ws);

bool ig_word_iterator_is_word_like(IGWordSegmenter *ws);

void ig_free_word_segmenter(IGWordSegmenter *ws);

#endif // ICUGO_SEGMENTER_H