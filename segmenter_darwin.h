#import <stdlib.h>
#import <stdint.h>
#import <stdbool.h>
#import <CoreFoundation/CoreFoundation.h>
void *C_CFStringTokenizer_CreateSentence(void *s);
void *C_CFStringTokenizer_CreateWord(void *s);

typedef struct
{
    bool hasNonLetters;
    long position;
} TokenizerResult;

TokenizerResult CFStringTokenizerRef_Next(void *tokenizer);
void *C_CFString_Create(const char *s);
void C_CFString_Release(void *s);
void C_CFStringTokenizer_Release(void *t);
