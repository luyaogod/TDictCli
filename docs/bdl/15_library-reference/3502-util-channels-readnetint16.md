---
title: "util.Channels.readNetInt16"
source: "fgl-topics/c_fgl_ext_util_Channels_readNetInt16.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Channels class > util.Channels methods > util.Channels.readNetInt16"
type: "concept"
---

# util.Channels.readNetInt16

> Reads next two bytes from a base.Channel as a 16 bit integer, in network byte order.

## Syntax

> **Important:**
>
> This feature is experimental, the syntax/name and
> semantics/behavior may change in a future version.

```
util.Channels.readNetInt16(
     in base.Channel )
RETURNS SMALLINT
```

1. in is the channel to read from.
2. Returns the 16 bit integer.

## Usage

The `util.Channels.readNetInt16()` reads two bytes as a 16 bit integer, using the
network byte order convention, from a [`base.Channel`](2980-the-channel-class.md "The base.Channel class is a built-in class providing basic input/output functions.") object.

## Related links

**Related concepts**  

[The Channel class](2980-the-channel-class.md "The base.Channel class is a built-in class providing basic input/output functions.")
