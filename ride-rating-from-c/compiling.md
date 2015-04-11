This will document my attempts to compile the OpenRCT2 project from outside of
the project, so I can call functions in it from the outside.

## Compiling with GCC

When attempting to link `OpenRCT2/src/ride/ride_ratings.c`, you get two errors:

- unknown use of instruction mnemonic without a size suffix
- error: brackets expression not supported on this target

Both of these sound like either the source needs to be changed or GCC does not
support compiling inline ASM yet.

### Update

Surprise! GCC is not actually GCC, it's a linked version of Clang. Downloading
GCC now.

The actual GCC 4.9 fails because of a `error: -masm=intel not supported in this
configuration` error. Apparently the Mac assembler can't support Intel syntax.

## Compiling with Clang

I tried

```
clang -m32 -mllvm  --x86-asm-syntax=intel -std=gnu99 get_rating.c
```

This failed with the same errors as above.

## Rewriting the syntax

The next suggestion was to add a GAS-compatible version of the assembler in
addresses.h. I might do this, and/or maintain a fork that doesn't ever call
RCT_CALLPROC/CALLFUNC.
