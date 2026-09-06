---
title: "base.Channel.getExitStatus"
source: "fgl-topics/c_fgl_ClassChannel_getExitStatus.html"
breadcrumb: "Library reference > Built-in packages > The base package > The Channel class > base.Channel methods > base.Channel.getExitStatus"
type: "concept"
---

# base.Channel.getExitStatus

> Returns the exit status of the child process of a pipe channel.

## Syntax

```
getExitStatus()
  RETURNS INTEGER
```

## Usage

The `getExitStatus()` method returns the exit status of the shell created to
execute the command passed as parameter to an [`openPipe()`](2991-base-channel-openpipe.md "Opens a pipe channel to a subprocess.") method.

The `getExitStatus()` method must be called after closing the channel with [`close()`](2983-base-channel-close.md "Closes the channel."): The method returns
`NULL`, when it is called before `close()`, or when the channel is not
associated with a process created by `openPipe()`.

In case of successful command execution and termination, the `getExitStatus()`
returns zero.

On UNIX™-like systems, if the child process is terminated by
the signal `signum`, `getExitStatus()` always
returns `(128 + signum)`. For example, `134 (128 +
6)` is returned, if the child process is terminated by `SIGABRT`
(`6`). Note that the process associated with a channel pipe is always a shell and the
returned status is the status of this shell. The shell may propagate the status returned by the
executed command.

Best practice to check the execution status of a child process of a pipe channel is to compare
`getExitStatus()` to zero: Zero means success, and any other value indicates a
failure.

## Example

The following code example checks the exit status of a pipe channel:

```
MAIN
  DEFINE ch base.Channel
  LET ch = base.Channel.create()
  CALL ch.openPipe("cat","w")
  CALL ch.writeLine("xxxxxxxxx")
  CALL ch.close()
  IF ch.getExitStatus() <> 0 THEN
     DISPLAY "pipe command failed!"
  END IF
END MAIN
```

## Related links

**Related concepts**  

[base.Channel.openPipe](2991-base-channel-openpipe.md "Opens a pipe channel to a subprocess.")

[base.Channel.close](2983-base-channel-close.md "Closes the channel.")
