---
title: "Read and write text lines"
source: "fgl-topics/c_fgl_ClassChannel_readwrite_lines.html"
breadcrumb: "Library reference > Built-in packages > The base package > The Channel class > Usage > Read and write text lines"
type: "concept"
description: "When the channel is open, use the readLine() / writeLine() methods to read and write simple text lines of data terminated by a line terminator . When using the readLine() and writeLine() functions, a ..."
---

# Read and write text lines

When the channel is open, use the [`readLine()`](2994-base-channel-readline.md "Read a complete line from the channel.")/[`writeLine()`](2998-base-channel-writeline.md "Write a complete line to the channel.") methods to read and write simple text lines of data terminated
by a [line terminator](3003-line-terminators-on-windows-and-unix.md).

When using the `readLine()` and `writeLine()` functions,
a LF character represents the end of a line.

For example, a simple text file can look like this:

```
first line
second line
third line
```

Sample code to read the above text file is as follows:

```
MAIN
  DEFINE i INTEGER
  DEFINE s STRING
  DEFINE ch base.Channel 
  LET ch = base.Channel.create()
  CALL ch.openFile("file.txt","r")
  LET i = 1
  WHILE TRUE
    LET s = ch.readLine()
    IF ch.isEof() THEN EXIT WHILE END IF
    DISPLAY i, " ", s 
    LET i = i + 1
  END WHILE
  CALL ch.close()
END MAIN
```

LF characters escaped by a backslash are **not** interpreted
as part of the line during a `readLine()` call.

When a line is written, any LF characters in the string will be
written as is to the output. When a line is read, the
LF escaped by a backslash is **not** interpreted as
part of the line.

For example, this code:

```
CALL ch.writeLine("aaa\\\nbbb")  -- [aaa<bs><lf>bbb]
CALL ch.writeLine("ccc\nddd")    -- [ccc<lf>ddd]
```

would generate this output:

```
aaa\
bbb 
ccc 
ddd
```

and the subsequent `readLine()` will read four different
lines, where the first line is ended by a backslash:

```
Read 1 aaa<bs>
Read 2 bbb 
Read 3 ccc 
Read 4 ddd
```
