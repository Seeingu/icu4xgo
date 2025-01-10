#import <Foundation/Foundation.h>
#import <CoreFoundation/CoreFoundation.h>
#import "segmenter_darwin.h"

void* C_CFString_Create(const char *s) {
  return CFStringCreateWithCString(kCFAllocatorDefault, s, kCFStringEncodingUTF8);
}

void C_CFString_Release(void *s) {
  CFRelease((CFStringRef)s);
}

void* C_CFStringTokenizer_Create(void* s, CFOptionFlags flag) {
  CFStringRef string = (CFStringRef)s;
  return CFStringTokenizerCreate(kCFAllocatorDefault,
                                 string,
                                 CFRangeMake(0, CFStringGetLength(string)),
                                 flag,
                                 CFLocaleCopyCurrent());
}

void* C_CFStringTokenizer_CreateSentence(void* s) {
  return C_CFStringTokenizer_Create(s, kCFStringTokenizerUnitSentence);
}

// C_CFStringTokenizer_CreateWord include word boundaries
void* C_CFStringTokenizer_CreateWord(void* s) {
  return C_CFStringTokenizer_Create(s, kCFStringTokenizerUnitWord|kCFStringTokenizerUnitWordBoundary);
}


void C_CFStringTokenizer_Release(void* t) {
  CFStringTokenizerRef tokenizer = (CFStringTokenizerRef)t;
  CFRelease(tokenizer);
}


TokenizerResult CFStringTokenizerRef_Next(void* tokenizer) {
  CFOptionFlags tokenType = CFStringTokenizerAdvanceToNextToken(tokenizer);
  TokenizerResult result;
  if (tokenType == kCFStringTokenizerTokenNone) {
    result.hasNonLetters = false;
    result.position = -1;
    return result;
  }
  result.hasNonLetters = (tokenType & kCFStringTokenizerTokenHasNonLettersMask) != 0;
  CFRange tokenRange = CFStringTokenizerGetCurrentTokenRange(tokenizer);
  result.position = tokenRange.location + tokenRange.length;
  return result;
}

