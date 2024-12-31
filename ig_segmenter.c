#include "./ig_segmenter.h"
#include <string.h>

IGGraphemeClusterSegmenter *ig_init_grapheme_segmenter()
{
    DataProvider *dataProvider = icu4x_DataProvider_compiled_mv1();
    IGGraphemeClusterSegmenter *segmenter = malloc(sizeof(IGGraphemeClusterSegmenter));
    icu4x_GraphemeClusterSegmenter_create_mv1_result result = icu4x_GraphemeClusterSegmenter_create_mv1(dataProvider);
    segmenter->is_ok = result.is_ok;
    if (!segmenter->is_ok)
    {
        return segmenter;
    }
    segmenter->segmenter = result.ok;
    return segmenter;
}

void ig_init_grapheme_iterator_utf8(IGGraphemeClusterSegmenter *gs, const char *s)
{
    DiplomatStringView input = ig_init_string(s);
    gs->iterator.utf8 = icu4x_GraphemeClusterSegmenter_segment_utf8_mv1(gs->segmenter, input);
}

int ig_grapheme_iterator_next(IGGraphemeClusterSegmenter *gs)
{
    if (gs->iterator.utf8)
    {
        return icu4x_GraphemeClusterBreakIteratorUtf8_next_mv1(gs->iterator.utf8);
    }
    else if (gs->iterator.utf16)
    {
        return icu4x_GraphemeClusterBreakIteratorUtf16_next_mv1(gs->iterator.utf16);
    }
    else if (gs->iterator.latin1)
    {
        return icu4x_GraphemeClusterBreakIteratorLatin1_next_mv1(gs->iterator.latin1);
    }
    return -1;
}

void ig_free_grapheme_segmenter(IGGraphemeClusterSegmenter *gs)
{
    icu4x_GraphemeClusterSegmenter_destroy_mv1(gs->segmenter);
    icu4x_GraphemeClusterBreakIteratorUtf8_destroy_mv1(gs->iterator.utf8);
    icu4x_GraphemeClusterBreakIteratorUtf16_destroy_mv1(gs->iterator.utf16);
    icu4x_GraphemeClusterBreakIteratorLatin1_destroy_mv1(gs->iterator.latin1);
    free(gs);
}

IGWordSegmenter *ig_init_word_segmenter()
{
    DataProvider *dataProvider = icu4x_DataProvider_compiled_mv1();
    IGWordSegmenter *segmenter = malloc(sizeof(IGWordSegmenter));
    icu4x_WordSegmenter_create_auto_mv1_result result =
        icu4x_WordSegmenter_create_auto_mv1(dataProvider);
    segmenter->is_ok = result.is_ok;
    if (!segmenter->is_ok)
    {
        return segmenter;
    }
    segmenter->segmenter = result.ok;
    return segmenter;
}

void ig_init_word_iterator_utf8(IGWordSegmenter *ws, const char *s)
{
    DiplomatStringView input = ig_init_string(s);
    ws->iterator.utf8 = icu4x_WordSegmenter_segment_utf8_mv1(ws->segmenter, input);
}

int ig_word_iterator_next(IGWordSegmenter *ws)
{
    if (ws->iterator.utf8)
    {
        return icu4x_WordBreakIteratorUtf8_next_mv1(ws->iterator.utf8);
    }
    else if (ws->iterator.utf16)
    {
        return icu4x_WordBreakIteratorUtf16_next_mv1(ws->iterator.utf16);
    }
    else if (ws->iterator.latin1)
    {
        return icu4x_WordBreakIteratorLatin1_next_mv1(ws->iterator.latin1);
    }
    return -1;
}

bool ig_word_iterator_is_word_like(IGWordSegmenter *ws)
{
    if (ws->iterator.utf8 == NULL)
    {
        // TODO: unimplemented
        exit(1);
    }
    return icu4x_WordBreakIteratorUtf8_is_word_like_mv1(ws->iterator.utf8);
}

void ig_free_word_segmenter(IGWordSegmenter *ws)
{
    icu4x_WordSegmenter_destroy_mv1(ws->segmenter);
    icu4x_WordBreakIteratorUtf8_destroy_mv1(ws->iterator.utf8);
    icu4x_WordBreakIteratorUtf16_destroy_mv1(ws->iterator.utf16);
    icu4x_WordBreakIteratorLatin1_destroy_mv1(ws->iterator.latin1);
    free(ws);
}

IGSentenceSegmenter *ig_init_sentence_segmenter()
{
    DataProvider *dataProvider = icu4x_DataProvider_compiled_mv1();
    IGSentenceSegmenter *segmenter = malloc(sizeof(IGSentenceSegmenter));
    icu4x_SentenceSegmenter_create_mv1_result result =
        icu4x_SentenceSegmenter_create_mv1(dataProvider);
    segmenter->is_ok = result.is_ok;
    if (!segmenter->is_ok)
    {
        return segmenter;
    }
    segmenter->segmenter = result.ok;
    return segmenter;
}

void ig_init_sentence_iterator_utf8(IGSentenceSegmenter *ss, const char *s)
{
    DiplomatStringView input = ig_init_string(s);
    ss->iterator.utf8 = icu4x_SentenceSegmenter_segment_utf8_mv1(ss->segmenter, input);
}

int ig_sentence_iterator_next(IGSentenceSegmenter *ss)
{
    if (ss->iterator.utf8)
    {
        return icu4x_SentenceBreakIteratorUtf8_next_mv1(ss->iterator.utf8);
    }
    else if (ss->iterator.utf16)
    {
        return icu4x_SentenceBreakIteratorUtf16_next_mv1(ss->iterator.utf16);
    }
    else if (ss->iterator.latin1)
    {
        return icu4x_SentenceBreakIteratorLatin1_next_mv1(ss->iterator.latin1);
    }
    return -1;
}

void ig_free_sentence_segmenter(IGSentenceSegmenter *ss)
{
    icu4x_SentenceSegmenter_destroy_mv1(ss->segmenter);
    icu4x_SentenceBreakIteratorUtf8_destroy_mv1(ss->iterator.utf8);
    icu4x_SentenceBreakIteratorUtf16_destroy_mv1(ss->iterator.utf16);
    icu4x_SentenceBreakIteratorLatin1_destroy_mv1(ss->iterator.latin1);
    free(ss);
}