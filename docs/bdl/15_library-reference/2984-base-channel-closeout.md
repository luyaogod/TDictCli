---
title: "base.Channel.closeOut"
source: "fgl-topics/c_fgl_ClassChannel_closeOut.html"
breadcrumb: "Library reference > Built-in packages > The base package > The Channel class > base.Channel methods > base.Channel.closeOut"
type: "concept"
---

# base.Channel.closeOut

> Closes the writing stream of the channel.

## Syntax

```
closeOut()
```

## Usage

When the channel object writes to and reads from a child process (using [`openPipe()`](2991-base-channel-openpipe.md "Opens a pipe channel to a subprocess.") with `"u"`
or `"ub"` mode), the output stream of the channel is connected to the
stdin stream of the child process, and the input stream of the
channel is connected to the stdout stream of the child process.

Calling the `closeOut()` method will close the output stream writing
to the child process, which is equivalent to an EOF on the stdin of the child
process.

The `closeOut()` function must be called when the child process waits for an EOF
to process the data that was written to its stdin stream.

If the channel is unidirectional and only writing to the child process
(`openPipe()` was called with the `"w"` or `"wb"`
mode), `closeOut()` behaves the same as [`close()`](2983-base-channel-close.md "Closes the channel.").

## Example

```
DEFINE x INTEGER
FOR x=1 TO 10
    CALL ch.writeLine( SFMT("line %1",x) )
END FOR
CALL ch.closeOut() 
DISPLAY ch.readLine()
```

For a complete example, see [Example 6: Closing the output stream](3014-example-6-closing-the-output-stream.md).

## Related links

**Related concepts**  

[base.Channel.close](2983-base-channel-close.md "Closes the channel.")

[base.Channel.flush](2986-base-channel-flush.md "Flushes the channel.")
