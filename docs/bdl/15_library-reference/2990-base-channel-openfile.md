---
title: "base.Channel.openFile"
source: "fgl-topics/c_fgl_ClassChannel_openFile.html"
breadcrumb: "Library reference > Built-in packages > The base package > The Channel class > base.Channel methods > base.Channel.openFile"
type: "concept"
---

# base.Channel.openFile

> Opens a file channel.

## Syntax

```
openFile(
   path STRING,
   mode STRING )
```

1. path is the path to the file to open, can be `NULL`
   for stdin/stdout.
2. mode is the open mode. Can be `"r"`, `"w"`,
   `"a"` or `"u"` (combined with `"b"` if needed).

## Usage

The `openFile()` method can be used to open a file for reading, writing, or both.

When passing `NULL` as filename, the channel can be used to read and/or write to
stdout or stdin, depending on the mode value.

When passing `"<stderr>"` as filename the standard error stream will be used.
The application can then print messages to stderr (typically for batch programs). See [Example 5: Writing to STDERR](3013-example-5-writing-to-stderr.md)

The opening mode can be one of the following:

- `r`: For read mode: reads from a file (standard input if path is
  `NULL`).
- `w`: For write mode: starts with an empty file (standard output if the path is
  `NULL`).
- `a`: For append mode: writes at the end of a file (standard output if the path is
  `NULL`).
- `u`: For read from standard input and write to standard output (path must be
  `NULL`).

Any of these modes can be followed by `b`, to use binary mode and
avoid CR/LF translation on Windows®
platforms. The binary mode is only required in specific cases, and will only take effect when
writing data.

If the opening mode is not one of the above letters, the method will raise
error [-8085](4483-genero-bdl-errors.md)

When you use the `w` or `a` modes, the file is created if it does not
exist.

The method raises error [-6340](4483-genero-bdl-errors.md) sif the file cannot be opened.

## Example

```
CALL ch.openFile( "file.txt", "w" )
```

For a complete example, see [Example 1: Using record-formatted data file](3009-example-1-using-record-formatted-data-file.md).
