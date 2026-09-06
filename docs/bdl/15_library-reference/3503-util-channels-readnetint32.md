---
title: "util.Channels.readNetInt32"
source: "fgl-topics/c_fgl_ext_util_Channels_readNetInt32.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Channels class > util.Channels methods > util.Channels.readNetInt32"
type: "concept"
---

# util.Channels.readNetInt32

> Reads next four bytes from a base.Channel as a 32 bit integer, in network byte order.

## Syntax

> **Important:**
>
> This feature is experimental, the syntax/name and
> semantics/behavior may change in a future version.

```
util.Channels.readNetInt32(
     in base.Channel )
RETURNS INTEGER
```

1. in is the channel to read from.
2. Returns the 32 bit integer.

## Usage

The `util.Channels.readNetInt16()` reads four bytes as a 32 bit integer, using the
network byte order convention, from a [`base.Channel`](2980-the-channel-class.md "The base.Channel class is a built-in class providing basic input/output functions.") object.

## Related links

**Related concepts**  

[The Channel class](2980-the-channel-class.md "The base.Channel class is a built-in class providing basic input/output functions.")
