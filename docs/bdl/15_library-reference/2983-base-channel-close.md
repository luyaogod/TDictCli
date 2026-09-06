---
title: "base.Channel.close"
source: "fgl-topics/c_fgl_ClassChannel_close.html"
breadcrumb: "Library reference > Built-in packages > The base package > The Channel class > base.Channel methods > base.Channel.close"
type: "concept"
---

# base.Channel.close

> Closes the channel.

## Syntax

```
close()
```

## Usage

Call the `close()` method when you are finished using the channel. The channel can
be re-opened after it has been closed.

A channel is automatically closed, when the channel object is destroyed.

## Example

```
CALL ch.close()
```

For a complete example, see [Example 1: Using record-formatted data file](3009-example-1-using-record-formatted-data-file.md).

## Related links

**Related concepts**  

[base.Channel.flush](2986-base-channel-flush.md "Flushes the channel.")
