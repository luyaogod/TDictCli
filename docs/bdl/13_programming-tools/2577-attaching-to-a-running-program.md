---
title: "Attaching to a running program"
source: "fgl-topics/c_fgl_Debugger_008.html"
breadcrumb: "Programming tools > Integrated debugger > Attaching to a running program"
type: "concept"
---

# Attaching to a running program

> It is possible to start the debugger for a program running on the same computer.

## Basics

Use the fgldb command with the `-p
pid` option, to switch an fglrun process
running on the same computer into debug mode.

The fgldb command must be executed on the machine where the
fglrun process executes.

> **Tip:**
>
> The graphical debugger of Genero Studio is able to start the
> fgldb tool on another computer through ssh, so that it is possible to
> debug a program running on a remote machine. See the Genero Studio User Guide for
> more details.

The fgldb command line tool takes the fglrun process id as value for the
`-p` argument.

Before starting a debug session, make sure that you fulfill the [prerequisites for debugging](2575-prerequisites-to-run-the-debugger.md "Some requirements are needed before running the integrated debugger.").

## Debug a program running on a UNIX® server

First, identify the process id of the fglrun program running on your
server.

For example, on a UNIX platform, use
the ps command:

```
$ ps a | grep fglrun
10646 pts/0    S+     0:00 /opt/myapp/fgl/lib/fglrun stockinfo.42m
```

When using the Genero Application Server, inspect the GAS log files to find the pid of a
fglrun process running behind a GAS application server. Enable full log reports in the GAS
to get detailed information about process execution.

You may want to debug processes that use a lot of machine resources (processor, memory or
open files). Use a system utility to find a process id by resources used (for example, the
top command on Linux®).

Execute the fgldb tool with the process id of the program you want to
attach to:

```
$ fgldb -p 10646
108	    DISPLAY ARRAY contlist TO sr.*
(fgldb)
```

The `(fgldb)` prompt indicates that you are now connected to the
fglrun process, and the program flow is suspended. To continue with the
program flow, enter the "`continue`" debugger
command:

```
(fgldb) continue
Continuing.
```

The application will then resume. To suspend the program again and enter debugger commands,
press CTRL-C in the debug console. fgldb will display the interrupt message and
return control to the debugger:

```
...
Continuing.
^CINTERRUPT
108	    DISPLAY ARRAY contlist TO sr.*
(fgldb)
```

At this point, you can enter debugger commands. For example, set a break point and continue
until the break point is reached:

```
(fgldb) b 427
Breakpoint 2 at 0x00000000: file contacts.4gl, line 427.
(fgldb) continue 
Continuing.
Breakpoint 2, edit_contact() at contacts.4gl:427
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
