---
title: "Line terminators on Windows and UNIX"
source: "fgl-topics/c_fgl_ClassChannel_line_terminators_windows.html"
breadcrumb: "Library reference > Built-in packages > The base package > The Channel class > Usage > Line terminators on Windows and UNIX"
type: "concept"
description: "On Windows® platforms, DOS formatted text files use CR/LF as line terminators. You can manage these type of files with the base.Channel class. By default, on both Windows and UNIX™ platforms, when ..."
---

# Line terminators on Windows and UNIX

On Windows® platforms, DOS formatted text files use
CR/LF as line terminators. You can manage these type of files with the `base.Channel`
class.

By default, on both Windows and UNIX™ platforms, when records are read from a DOS file with the
`base.Channel` class, the CR/LF line terminator is removed. When a record is written
to a file on Windows, the lines are terminated with CR/LF
in the file; on UNIX, the lines are terminated with LF
only.

To avoid the automatic translation of CR/LF on Windows,
you can use the `b` option of the `openFile()` and
`openPipe()` methods. You can combine the `b` option with
`r` or `w`, based on the read or write operations that you want to
perform.

```
CALL ch.openFile( "mytext.txt", "rb" )
```

On Windows, when lines are read with the
`b` option, only LF is removed from CR/LF line terminators; CR will be copied as a
character part of the last field. In contrast, when lines are written with the `b`
option, LF characters will not be converted to CR/LF.

On UNIX, writing lines with
or without the binary mode option does not matter.
