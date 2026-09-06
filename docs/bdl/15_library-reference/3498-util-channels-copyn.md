---
title: "util.Channels.copyN"
source: "fgl-topics/c_fgl_ext_util_Channels_copyN.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Channels class > util.Channels methods > util.Channels.copyN"
type: "concept"
---

# util.Channels.copyN

> Copies a given number of bytes from one base.Channel to another.

## Syntax

```
util.Channels.copyN(
     in base.Channel,
     out base.Channel,
     n INTEGER
 )
RETURNS INTEGER
```

1. in is the channel to read from.
2. out is the channel to write to.
3. n is the number of bytes to be copied.
4. Returns the number of bytes copied.

## Usage

The `util.Channels.copyN()` copies data between two [`base.Channel`](2980-the-channel-class.md "The base.Channel class is a built-in class providing basic input/output functions.") objects.

If the returned value is less than n, then the source channel has reached
EOF.

## Related links

**Related concepts**  

[The Channel class](2980-the-channel-class.md "The base.Channel class is a built-in class providing basic input/output functions.")
