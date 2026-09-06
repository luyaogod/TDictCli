---
title: "base.Channel.dataAvailable"
source: "fgl-topics/c_fgl_ClassChannel_dataAvailable.html"
breadcrumb: "Library reference > Built-in packages > The base package > The Channel class > base.Channel methods > base.Channel.dataAvailable"
type: "concept"
---

# base.Channel.dataAvailable

> Tests if some data can be read from the channel.

## Syntax

```
dataAvailable()
  RETURNS BOOLEAN
```

## Usage

The `dataAvailable()` method returns [`TRUE`](../08_language-basics/0573-true.md "TRUE is a predefined constant to be used in boolean expressions.") if some data can be read from the channel.

This method is only to be used in some rare cases. Use `dataAvailable()` if
the protocol allows asynchronous messages from the peer. An example is an asynchronous error
message from the peer, to stop sending more data.

`dataAvailable()` checks if at least one byte is available on the stream. A
subsequent read will block, if the read operation can not be completed. This should not
happen: the methods `read()` and `readLine()` and their
counterparts `write()` and `writeLine()` read and write
complete lines (a line is a sequence of characters terminated by the line separator).

The `dataAvailable()` method opens up the possibility to read data
asynchronously. One possible use for this method is to stop a data transfer from a local site, after
receiving an error message from the remote site.

## Example

The local site (parent.4gl) sends a huge amount of data to the remote site
(child.4gl) using `base.Channel.writeLine()`. If an error occurs
on the remote side during the processing of data, the remote site writes an error message to the
channel, causing the local site to stop the data transmission.

On the local site, the file is parent.4gl.

```
MAIN
  DEFINE n INT
  DEFINE c base.Channel
  LET c = base.Channel.create()
  CALL c.openPipe("fglrun child", "u")
  LET n = 0
  WHILE TRUE
    IF c.dataAvailable() THEN
      DISPLAY "message from child: ", c.readLine()
      EXIT WHILE
    END IF
    LET n = n + 1
    DISPLAY "parent: write line ", n
    CALL c.writeLine("line " || n)
  END WHILE
END MAIN
```

On the remote site, the file is child.4gl (do not add
`DISPLAY` instructions in the child program, as this would corrupt the pipe
communication: The parent program opens a pipe channel read from child's stdout and write to child's
stdin):

```
MAIN
  DEFINE c base.Channel
  DEFINE s STRING
  DEFINE n INT
  LET c = base.Channel.create()
  CALL c.openFile("", "u")
  LET n = 0
  WHILE TRUE
    LET s = c.readLine()
    IF c.isEof() THEN EXIT WHILE END IF
    LET n = n + 1
    IF n == 3 THEN
      CALL c.writeLine("error: something happens")
      CALL readRemainingData(c)
      EXIT WHILE
    END IF
  END WHILE
END MAIN

FUNCTION readRemainingData(c)
  DEFINE c base.Channel
  DEFINE s STRING
  WHILE TRUE
    LET s = c.readLine()
    IF c.isEof() THEN EXIT WHILE END IF
  END WHILE
END FUNCTION
```

## Related links

**Related concepts**  

[base.Channel.read](2993-base-channel-read.md "Reads a list of data delimited by a separator from the channel.")

[base.Channel.readLine](2994-base-channel-readline.md "Read a complete line from the channel.")
