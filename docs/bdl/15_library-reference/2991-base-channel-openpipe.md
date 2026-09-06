---
title: "base.Channel.openPipe"
source: "fgl-topics/c_fgl_ClassChannel_openPipe.html"
breadcrumb: "Library reference > Built-in packages > The base package > The Channel class > base.Channel methods > base.Channel.openPipe"
type: "concept"
---

# base.Channel.openPipe

> Opens a pipe channel to a subprocess.

## Syntax

```
openPipe(
   command STRING,
   mode STRING )
```

1. command is the system command to be executed.
2. mode is the open mode. Can be `"r"`, `"w"` or
   `"u"` (combined with `"b"` if needed).

## Usage

With the `openPipe()` method, you can read from the standard output of a subprocess,
write to the standard input, or both.

> **Important:**
>
> This feature is not supported on mobile platforms.

The opening mode can be one of the following:

- `r`: For read only from standard output of the command.
- `w`: For write only to standard input of the command.
- `u`: For read from standard output and write to standard input of the command.

Any of these modes can be followed by `b`, to use binary mode and
avoid CR/LF translation on Windows®
platforms. The binary mode is only required in specific cases, and will only take effect when
writing data.

If the opening mode is not one of the above letters, the method will raise error
[-8085](4483-genero-bdl-errors.md).

Some commands like the Unix tr command buffer the stdout stream. This may
block the [`readLine()`](2994-base-channel-readline.md "Read a complete line from the channel.") calls,
when using the `openPipe()` method with `"u"` option, and writing to
stdin / reading from stdout line by line. To avoid
`readLine()` to block in such case, you must force the command to flush each line
written to stdout, for example with the stdbuf -oL
command.

When using the `"u"` mode to write to and read from a child process, it is
possible to close only the writing stream with the [`closeOut()`](2984-base-channel-closeout.md "Closes the writing stream of the channel.") method. Using this method is equivalent to an EOF in a
command-line pipe.

To get the exit status of the command executed by the pipe channel, use the [`getExitStatus()`](2987-base-channel-getexitstatus.md "Returns the exit status of the child process of a pipe channel.") method after
closing the channel with [`close()`](2983-base-channel-close.md "Closes the channel.").

> **Important:**
>
> On Windows: In unidirectional modes
> (`"r"` and `"w"`), `openPipe()` creates a command
> processor (cmd.exe) and uses the command parameter as command
> line. In contrast, when opening a bidirectional pipe (mode `"u"`), the
> command parameter is used as is when spawning the process. This means that
> environment variables are not replaced and special characters such as
> `<>|()^&` have no special meaning.
>
> The current behavior cannot be changed without potentially breaking existing user code.
>
> If needed, prefix the command to be executed with "cmd /c", but take care with
> cmd command line interpretation for example when using double quotes.

## Example

```
CALL ch.openPipe( "ls", "r" )
```

For a complete example, see [Example 2: Executing UNIX commands](3010-example-2-executing-unix-commands.md).
