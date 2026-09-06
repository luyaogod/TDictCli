---
title: "base.Channel methods"
source: "fgl-topics/c_fgl_ClassChannel_methods.html"
breadcrumb: "Library reference > Built-in packages > The base package > The Channel class > base.Channel methods"
type: "concept"
description: "Table 1. Class methods Name Description base.Channel.create () RETURNS base.Channel Create a new channel object. Table 2. Object methods Name Description dataAvailable () RETURNS BOOLEAN Tests if some ..."
---

# base.Channel methods

| Name | Description |
| --- | --- |
| base.Channel.create() RETURNS base.Channel | Create a new channel object. |

| Name | Description |
| --- | --- |
| dataAvailable() RETURNS BOOLEAN | Tests if some data can be read from the channel. |
| close() | Closes the channel. |
| closeOut() | Closes the writing stream of the channel. |
| flush() | Flushes the channel. |
| getExitStatus() RETURNS INTEGER | Returns the exit status of the child process of a pipe channel. |
| isEof() RETURNS BOOLEAN | Detect the end of a file. |
| openFile( path STRING, mode STRING ) | Opens a file channel. |
| openPipe( command STRING, mode STRING ) | Opens a pipe channel to a subprocess. |
| openClientSocket( host STRING, port INTEGER, mode STRING, timeout INTEGER ) | Open a TCP client socket channel. |
| openServerSocket( interface STRING, port INTEGER, mode STRING ) | Open a TCP server socket channel. |
| read( variableList ) RETURNS INTEGER | Reads a list of data delimited by a separator from the channel. |
| readLine() RETURNS STRING | Read a complete line from the channel. |
| readOctets( length INTEGER ) RETURNS STRING | Read a given number of bytes and return as a character string. |
| setDelimiter( delimiter STRING ) | Define the value delimiter for a channel. |
| write( valueList ) | Writes a list of data delimited by a separator to the channel. |
| writeLine( value STRING ) | Write a complete line to the channel. |
| writeNoNL( value STRING ) | Writes a string to the channel (without newline character). |
