---
title: "util.Channels.writeNetInt32"
source: "fgl-topics/c_fgl_ext_util_Channels_writeNetInt32.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Channels class > util.Channels methods > util.Channels.writeNetInt32"
type: "concept"
---

# util.Channels.writeNetInt32

> Writes four bytes of a 32 bit integer to a base.Channel, in network byte order.

## Syntax

> **Important:**
>
> This feature is experimental, the syntax/name and
> semantics/behavior may change in a future version.

```
util.Channels.writeNetInt32(
     out base.Channel,
     v INTEGER
)
```

1. out is the channel to write to.
2. v is the 32 bit integer to be written.

## Usage

The `util.Channels.writeNetInt16()` writes the specified 32 bit integer in network
by order, into a [`base.Channel`](2980-the-channel-class.md "The base.Channel class is a built-in class providing basic input/output functions.")
object.

## Related links

**Related concepts**  

[The Channel class](2980-the-channel-class.md "The base.Channel class is a built-in class providing basic input/output functions.")
