#include "./ig_segmenter.h"
#include <string.h>

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

void *ig_init_word_iterator_utf8(IGWordSegmenter *ws, const char *s)
{
    DiplomatStringView input = {
        s,
        strlen(s)};
    ws->iterator.utf8 = icu4x_WordSegmenter_segment_utf8_mv1(ws->segmenter, input);
}

int ig_word_iterator_next(IGWordSegmenter *ws)
{
    if (ws->iterator.utf8 == NULL)
    {
        // TODO: unimplemented
        exit(1);
        return -1;
    }
    return icu4x_WordBreakIteratorUtf8_next_mv1(ws->iterator.utf8);
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