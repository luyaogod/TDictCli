---
title: "base.Channel.writeNoNL"
source: "fgl-topics/c_fgl_ClassChannel_writenonl.html"
breadcrumb: "Library reference > Built-in packages > The base package > The Channel class > base.Channel methods > base.Channel.writeNoNL"
type: "concept"
---

# base.Channel.writeNoNL

> Writes a string to the channel (without newline character).

## Syntax

```
writeNoNL(
   value STRING )
```

1. value is the character string to be written to the channel.

## Usage

After opening a channel, use the `writeNoNL()` method to write a string
to the channel, without a trailing newline character.

Note that the `writeNoNL()` and `write()` methods have distinct
purpose. The first is provided to write raw character strings to the stream. The second is designed
to write records with formatted data and field delimiters. The Channel class also provides the
`writeLine()` method, to write a string with an ending newline character.

Errors [-6344](4483-genero-bdl-errors.md) or [-6345](4483-genero-bdl-errors.md) are
thrown, if the channel fails to write data.

## Example

```
CALL ch.writeNoNL("Some text ...")
```

## Related links

**Related concepts**  

[base.Channel.readOctets](2995-base-channel-readoctets.md "Read a given number of bytes and return as a character string.")

[base.Channel.writeLine](2998-base-channel-writeline.md "Write a complete line to the channel.")

[base.Channel.write](2997-base-channel-write.md "Writes a list of data delimited by a separator to the channel.")
