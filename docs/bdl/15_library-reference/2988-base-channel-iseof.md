---
title: "base.Channel.isEof"
source: "fgl-topics/c_fgl_ClassChannel_isEof.html"
breadcrumb: "Library reference > Built-in packages > The base package > The Channel class > base.Channel methods > base.Channel.isEof"
type: "concept"
---

# base.Channel.isEof

> Detect the end of a file.

## Syntax

```
isEof()
  RETURNS BOOLEAN
```

## Usage

Use the `isEof()` method
to detect the end of a file while reading from a channel.

The
end of file is only detected after the last read. In other words,
you first read, then check for the end of file and process if not
end of file.

## Example

```
DEFINE s STRING
WHILE TRUE
  LET s = ch.readLine()
  IF ch.isEof() THEN 
     EXIT WHILE
  END IF
  DISPLAY s 
END WHILE
```

For a complete example, see [Example 3: Reading lines from a text file](3011-example-3-reading-lines-from-a-text-file.md).
