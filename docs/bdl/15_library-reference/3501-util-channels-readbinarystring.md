---
title: "util.Channels.readBinaryString"
source: "fgl-topics/c_fgl_ext_util_Channels_readBinaryString.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Channels class > util.Channels methods > util.Channels.readBinaryString"
type: "concept"
---

# util.Channels.readBinaryString

> Reads a given number of bytes from a base.Channel as a character string.

## Syntax

> **Important:**
>
> This feature is experimental, the syntax/name and
> semantics/behavior may change in a future version.

```
util.Channels.readBinaryString(
     in base.Channel,
     n INTEGER
 )
RETURNS STRING
```

1. in is the channel to read from.
2. n is the number of bytes to be read.
3. Returns the character string corresponding to the number of bytes read.

## Usage

The `util.Channels.readBinaryString()` reads a number of bytes from a [`base.Channel`](2980-the-channel-class.md "The base.Channel class is a built-in class providing basic input/output functions.") object, and returns the
character string built from the byte sequence.

The bytes read from the channel must represent characters in the encoding defined by the current
[application locale](../09_advanced-features/0864-application-locale.md "The application locale defines the language and codeset for your application.").

## Related links

**Related concepts**  

[The Channel class](2980-the-channel-class.md "The base.Channel class is a built-in class providing basic input/output functions.")
