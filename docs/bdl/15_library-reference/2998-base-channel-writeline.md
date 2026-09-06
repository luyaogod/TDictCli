---
title: "base.Channel.writeLine"
source: "fgl-topics/c_fgl_ClassChannel_writeLine.html"
breadcrumb: "Library reference > Built-in packages > The base package > The Channel class > base.Channel methods > base.Channel.writeLine"
type: "concept"
---

# base.Channel.writeLine

> Write a complete line to the channel.

## Syntax

```
writeLine(
   value STRING )
```

1. value is the string expression to be written to the channel.

## Usage

After opening a channel, use the `writeLine()` method to write a
line of text to the channel.

The `writeLine()` method does not use the field delimiter, it writes the text data
to the stream, with an ending newline character.

To write a string with no ending newline character, use the `writeNoNL()`
method.

Errors [-6344](4483-genero-bdl-errors.md) or [-6345](4483-genero-bdl-errors.md) are
thrown, if the channel fails to write data.

## Example

```
CALL ch.writeLine("Customer number: "|| custno)
```

For a complete example, see [Example 5: Writing to STDERR](3013-example-5-writing-to-stderr.md).

## Related links

**Related concepts**  

[Read and write text lines](3002-read-and-write-text-lines.md "Read and write text lines")
