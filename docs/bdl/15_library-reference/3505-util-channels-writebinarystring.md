---
title: "util.Channels.writeBinaryString"
source: "fgl-topics/c_fgl_ext_util_Channels_writeBinaryString.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Channels class > util.Channels methods > util.Channels.writeBinaryString"
type: "concept"
---

# util.Channels.writeBinaryString

> Writes a character string to a base.Channel, without the trailing zero.

## Syntax

> **Important:**
>
> This feature is experimental, the syntax/name and
> semantics/behavior may change in a future version.

```
util.Channels.writeBinaryString(
     out base.Channel,
     s STRING
)
```

1. out is the channel to write to.
2. s is the character string to be written.

## Usage

The `util.Channels.writeBinaryString()` writes the specified string into a [`base.Channel`](2980-the-channel-class.md "The base.Channel class is a built-in class providing basic input/output functions.") object.

The bytes written to the channel represent characters in the encoding defined by the current
[application locale](../09_advanced-features/0864-application-locale.md "The application locale defines the language and codeset for your application.").

No string terminator (trailing zero) is written to the output channel.

## Related links

**Related concepts**  

[The Channel class](2980-the-channel-class.md "The base.Channel class is a built-in class providing basic input/output functions.")
