---
title: "util.Channels.readNetInt8"
source: "fgl-topics/c_fgl_ext_util_Channels_readNetInt8.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Channels class > util.Channels methods > util.Channels.readNetInt8"
type: "concept"
---

# util.Channels.readNetInt8

> Reads next byte from a base.Channel as a 8 bit integer.

## Syntax

> **Important:**
>
> This feature is experimental, the syntax/name and
> semantics/behavior may change in a future version.

```
util.Channels.readNetInt8(
     in base.Channel )
RETURNS SMALLINT
```

1. in is the channel to read from.
2. Returns the 8 bit integer.

## Usage

The `util.Channels.readNetInt16()` reads a single byte as an 8 bit integer, from a
[`base.Channel`](2980-the-channel-class.md "The base.Channel class is a built-in class providing basic input/output functions.") object.

## Related links

**Related concepts**  

[The Channel class](2980-the-channel-class.md "The base.Channel class is a built-in class providing basic input/output functions.")
