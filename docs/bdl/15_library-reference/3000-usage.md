---
title: "Usage"
source: "fgl-topics/c_fgl_ClassChannel_usage.html"
breadcrumb: "Library reference > Built-in packages > The base package > The Channel class > Usage"
type: "concept"
description: "The base.Channel class is a built-in class providing basic input/output functionality for: text file reading/writing subprocess communication (through pipes) basic network communication (through TCP ..."
---

# Usage

The `base.Channel` class is a built-in class providing
basic input/output functionality for:

- text file reading/writing
- subprocess communication (through pipes)
- basic network communication (through TCP sockets)

No character set conversion is done when reading or writing data with channel objects. The
character set used by the other end of the channel must correspond to the locale of the runtime
system, for both input and output. For more details, see [Character string encoding](3004-character-string-encoding.md).

Steps to use a channel object:

- Define a variable with the `base.Channel` type.
- Create a channel object with `base.Channel.create()` and assign it to
  the variable.
- Open the channel for a file, piped process or socket (as a client).
- Read and/or write data on the channel.
- Close the channel.

When reading or writing character strings, the escape character is the backslash
(`\`).

The are three modes to read and write data with Channels:

1. Reading/writing formatted data as a set of fields in a line (records), with the [`read()` and
   `write()` methods](3001-read-and-write-record-data.md).
2. Reading/writing complete lines with the [`readLine()` and `writeLine()` methods](3002-read-and-write-text-lines.md).
3. Handling raw character string data by reading/writing pieces of strings, with the [`readOctets()`](2995-base-channel-readoctets.md "Read a given number of bytes and return as a character string.") and [`writeNoNL()`](2999-base-channel-writenonl.md "Writes a string to the channel (without newline character).") methods.

Channels may raise exceptions that can be trapped with [`WHENEVER ERROR` or
`TRY/CATCH` blocks](3006-handle-channel-exceptions.md).

An extension class it provided to get better control over server TCP socket management with the
[`util.Channels`](3495-the-util-channels-class.md "The util.Channels class provides utility functions for base.Channel objects.") class.

## Child topics

- [Read and write record data](3001-read-and-write-record-data.md)
- [Read and write text lines](3002-read-and-write-text-lines.md)
- [Line terminators on Windows and UNIX](3003-line-terminators-on-windows-and-unix.md)
- [Character string encoding](3004-character-string-encoding.md)
- [BYTE data serialization](3005-byte-data-serialization.md)
- [Handle channel exceptions](3006-handle-channel-exceptions.md)
- [Setup a TCP socket channel](3007-setup-a-tcp-socket-channel.md)
