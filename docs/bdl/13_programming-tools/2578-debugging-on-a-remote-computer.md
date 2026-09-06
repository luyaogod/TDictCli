---
title: "Debugging on a remote computer"
source: "fgl-topics/c_fgl_Debugger_013.html"
breadcrumb: "Programming tools > Integrated debugger > Debugging on a remote computer"
type: "concept"
---

# Debugging on a remote computer

> It is possible to remotely start the debugger for an app running on another computer.

## Basics

An fglrun process running on a remote computer can be switched to debug mode,
to listen to the TCP port 4711 for debug instructions. It is then possible to attach to the remote
runtime system from another computer, by using the `fgldb -m
host[:port]` command.

The remote computer can be a production site, where only .42m PCode program
files have been deployed. The .4gl source files must be available on the
computer where the fgldb tool is used. This is typically the development
platform.

Before starting a debug session, make sure that you fulfill the [prerequisites for debugging](2575-prerequisites-to-run-the-debugger.md "Some requirements are needed before running the integrated debugger."). The remote and local computer
must use the same FGL environment.

The fgldb command line tool takes two arguments: The host (or IP address) of
the remote computer, and the TCP port number to connect to. For remote fglrun
processes, the debug TCP port is `4711`.

> **Important:**
>
> For security reasons, best practice is to connect to the remote computer with SSH, and use port
> forwarding for the TCP port 4711. With an SSH connection, communication between the
> fgldb tool and the fglrun debuggee will be encrypted.

## Debugging an app running on a remote computer

The first step is to turn the remote fglrun process into debug mode, and make
it listen to the TCP port 4711 for debug instructions. On a Unix-like system, this is done by
sending the SIGTRAP signal to the fglrun process; on a Microsoft™ Windows® system, this
is done by sending the CTRL\_BREAK\_EVENT console event to the process.

In order to send a signal on a Unix-like system, you must open a terminal session on the remote
computer and use the kill command.

Best practice is to open an SSH connection to the remote computer (`prodserver`),
using TCP 4711 port forwarding, so that the fgldb tool can open a TCP connection
to the port 4711 on the local computer, and communicate with the remote fglrun
debuggee:

```
$ ssh tom@prodserver -L 127.0.0.1:4711:127.0.0.1:4711
```

Once connected to the remote computer, find the process ID of the fglrun to
debug:

```
[tom@prodserver ~]$ ps -aef | grep "fglrun account"
ted     1025064 1024899  0 12:43 pts/0    00:00:00 fglrun account.42m
```

Send the SIGTRAP signal to switch the fglrun process into debug
mode:

```
[tom@prodserver ~]$ kill -SIGTRAP 1025064
```

The fglrun process is not listening on the TCP port 4711, and the
fgldb tool on the local development computer (`devhost`) can
attach to the remote runtime system through the SSH tunel (port forwarding),
:

```
[tom@devhost ~]$ fgldb -m 127.0.0.1:4711
... (current
```

The remote fglrun program execution is now suspended.

Use the debugger, for example, by defining a breakpoint:

```
(fgldb) b 356
Breakpoint 1 at 0x00000000: file account.4gl, line 356.
```

And let the program continue:

```
(fgldb) continue
Continuing.
```

The application will then resume on the remote computer. To suspend the program again and enter
debugger commands, send a SIGINT signal from the terminal connected to the remote computer via
SSH:

```
[tom@prodserver ~]$ kill -SIGINT 1025064
```

On Microsoft Windows, you need to
send the CTRL\_C\_EVENT console event to the process.

Then on the local compuer, fgldb will show the interrupt message and give you
the control back:

```
...
Continuing.
INTERRUPT
108	    DISPLAY ARRAY acclist TO sr.*
(fgldb)
```

At this point, you can for example set a break point and continue until the break point is
reached:

```
(fgldb) b 427
Breakpoint 2 at 0x00000000: file account.4gl, line 427.
(fgldb) continue 
Continuing.
Breakpoint 2, edit_account() at account.4gl:427
427	    IF new THEN
(fgldb)
```

To finish the debug session, close the connection with the "`detach`"
debugger command:

```
(fgldb) detach
Connection closed by foreign host.
```

## Related links

**Related concepts**  

[fgldb](2520-fgldb.md "The fgldb tool is an interface program for remote debugging.")
