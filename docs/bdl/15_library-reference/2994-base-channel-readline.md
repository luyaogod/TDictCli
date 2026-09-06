---
title: "base.Channel.readLine"
source: "fgl-topics/c_fgl_ClassChannel_readLine.html"
breadcrumb: "Library reference > Built-in packages > The base package > The Channel class > base.Channel methods > base.Channel.readLine"
type: "concept"
---

# base.Channel.readLine

> Read a complete line from the channel.

## Syntax

```
readLine()
  RETURNS STRING
```

## Usage

After opening the channel object, use the `readLine()` method to read a complete
line from the channel.

The `readLine()` method returns an empty string if the line is empty.

A call to `readLine()` is blocking until the read operation is complete: The
source must write characters, terminate the line with a NL character, and if buffered, it must flush
the stream.

The `readLine()` function returns [`NULL`](../08_language-basics/0572-null.md "The NULL constant defines a non-value.")
if end of file is reached. To distinguish empty lines from `NULL`, you must use the
`STRING` data type. If you use a `CHAR` or `VARCHAR`, you will get
`NULL` for empty lines. To detect the end of file, use the
[`isEof()`](2988-base-channel-iseof.md "Detect the end of a file.") method.

Error [-6346](4483-genero-bdl-errors.md) is thrown, if the channel fails to read data.

## Example

```
WHILE TRUE
  LET s = ch.readLine()
  IF ch.isEof() THEN EXIT WHILE END IF
  ...
END WHILE
```

For a complete example, see [Example 3: Reading lines from a text file](3011-example-3-reading-lines-from-a-text-file.md).

## Related links

**Related concepts**  

[Read and write text lines](3002-read-and-write-text-lines.md "Read and write text lines")

[base.Channel.dataAvailable](2985-base-channel-dataavailable.md "Tests if some data can be read from the channel.")
