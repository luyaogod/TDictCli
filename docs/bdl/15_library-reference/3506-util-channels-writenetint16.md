---
title: "util.Channels.writeNetInt16"
source: "fgl-topics/c_fgl_ext_util_Channels_writeNetInt16.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Channels class > util.Channels methods > util.Channels.writeNetInt16"
type: "concept"
---

# util.Channels.writeNetInt16

> Writes two bytes of a 16 bit integer to a base.Channel, in network byte order.

## Syntax

> **Important:**
>
> This feature is experimental, the syntax/name and
> semantics/behavior may change in a future version.

```
util.Channels.writeNetInt16(
     out base.Channel,
     v SMALLINT
)
```

1. out is the channel to write to.
2. v is the 16 bit integer to be written.

## Usage

The `util.Channels.writeNetInt16()` writes the specified 16 bit integer in network
by order, into a [`base.Channel`](2980-the-channel-class.md "The base.Channel class is a built-in class providing basic input/output functions.")
object.

## Related links

**Related concepts**  

[The Channel class](2980-the-channel-class.md "The base.Channel class is a built-in class providing basic input/output functions.")
