---
title: "Example 2: Executing UNIX commands"
source: "fgl-topics/c_fgl_ClassChannel_example_2.html"
breadcrumb: "Library reference > Built-in packages > The base package > The Channel class > Examples > Example 2: Executing UNIX™ commands"
type: "concept"
description: "Read the stdout of a command This program executes the ls command, read from the standard output stream of the command, to display the file names and extensions separately. The second parameter of the ..."
---

# Example 2: Executing UNIX commands

## Read the stdout of a command

This program executes the ls command, read from the standard output stream of
the command, to display the file names and extensions separately.

The second parameter of the [`openPipe()`](2991-base-channel-openpipe.md "Opens a pipe channel to a subprocess.") method is `"r"`, only to read from the standard
output of the command:

```
MAIN
  DEFINE fn, ex STRING
  DEFINE ch base.Channel
  LET ch = base.Channel.create()
  CALL ch.setDelimiter(".")
  CALL ch.openPipe("ls -1","r") -- Warning: ls option is -1 (digit one)
  WHILE ch.read([fn,ex])
    DISPLAY "Filename: ",fn, COLUMN 40, "extension: ",NVL(ex,"NONE")
  END WHILE
  CALL ch.close()
END MAIN
```

## Write into stdin of a command

The next example opens a pipe to write text lines into the standard input of the
sort command.

The second parameter of the [`openPipe()`](2991-base-channel-openpipe.md "Opens a pipe channel to a subprocess.") method is `"w"`, only to write into the standard
input of the command.

Note that the data will only be processed when closing the channel:

```
MAIN
  DEFINE ch base.Channel
  LET ch = base.Channel.create()
  CALL ch.openPipe("sort -u","w")
  CALL ch.writeLine("zzz")
  CALL ch.writeLine("aaa")
  CALL ch.writeLine("ooo")
  SLEEP 1
  CALL ch.writeLine("ooo")
  CALL ch.writeLine("bbb")
  CALL ch.writeLine("bbb")
  SLEEP 1
  CALL ch.close()
  DISPLAY "done."
END MAIN
```

## Write to stdin and read from stdout of a command

The next program executes the tr command, writing to its standard input stream
and reading the standard output stream.

The second parameter of the [`openPipe()`](2991-base-channel-openpipe.md "Opens a pipe channel to a subprocess.") command is `"u"`, to write to
stdin and read from stdout of the command.

Note that we must force the tr command to flush each line written to
stdout, with the stdbuf -oL command:

```
MAIN
  DEFINE ch base.Channel, x INTEGER
  LET ch = base.Channel.create()
  --CALL ch.openPipe("tee file.tmp","u") -- OK
  --CALL ch.openPipe("tr '[:lower:]' '[:upper:]'","u") -- Hangs in readLine()
  CALL ch.openPipe("stdbuf -oL tr '[:lower:]' '[:upper:]'","u") -- OK
  FOR x=1 TO 3
    CALL ch.writeLine(SFMT("line #%1",x))
    DISPLAY ch.readLine()
  END FOR
  CALL ch.close()
END MAIN
```

## Checking the execution status of the child process

The exit status of the shell started by the [`openPipe()`](2991-base-channel-openpipe.md "Opens a pipe channel to a subprocess.") method to execute a command can be checked after closing the
channel, by using the [`getExitStatus()`](2987-base-channel-getexitstatus.md "Returns the exit status of the child process of a pipe channel.") method.

The `getExitStatus()` method must be called after closing the channel with [`close()`](2983-base-channel-close.md "Closes the channel."):

```
MAIN
  DEFINE ch base.Channel

  LET ch = base.Channel.create()

  CALL ch.openPipe("cat","w")
  CALL ch.writeLine("xxxxxxxxx")
  CALL ch.close()
  DISPLAY "exit status 1: ", ch.getExitStatus()

  CALL ch.openPipe("zzz","w") -- invalid command!
  CALL ch.writeLine("xxxxxxxxx")
  CALL ch.close()
  DISPLAY "exit status 2: ", ch.getExitStatus()

END MAIN
```

Output:

```
xxxxxxxxx
exit status 1:           0
sh: 1: zzz: not found         <-- this is an error message of the shell
exit status 2:         127
```
