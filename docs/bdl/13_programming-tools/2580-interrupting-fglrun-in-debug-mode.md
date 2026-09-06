---
title: "Interrupting fglrun in debug mode"
source: "fgl-topics/c_fgl_Debugger_007.html"
breadcrumb: "Programming tools > Integrated debugger > Interrupting fglrun in debug mode"
type: "concept"
---

# Interrupting fglrun in debug mode

> The SIGINT signal will suspend the execution of the runtime system executing in debug mode.

When the fglrun process is executing in debug mode, the program flow can be
stopped by sending the SIGINT signal to the process.

This solution can be used on programs using the TUI mode, or the GUI mode.

If the fglrun command was executed from a terminal, use the CTRL-C key
combination, to send a SIGINT signal to the current process:

```
$ fglrun -d myprog
(fgldb) run
... program starts and displays in TUI or GUI mode ...

<CTRL>-<C>

   1     MAIN
-> 2         MENU "Test"
   3             COMMAND "Hello"
   4                 MESSAGE "Hello! "||CURRENT
   5             COMMAND "Quit"
(fgldb)
```

Once the program execution is stopped, the (fgldb) prompt appears, and you can enter a
debugger command.

Another way to stop the program execution on Unix-like systems is to use the kill
command:

```
$ fglrun -d myprog
(fgldb) run
... program starts and displays in TUI or GUI mode ...
```

In another terminal, find the fglrun process id (pid) and send the SIGINT
signal:

```
$ ps -aef | grep fglrun
mike      198695  195976  0 15:23 pts/2    00:00:00 fglrun -d myprog
...

$ kill -SIGINT 198695
```

In the other terminal executing fglrun, you should see that the runtime system
has switched to the debug command prompt:

```
   1     MAIN
-> 2         MENU "Test"
   3             COMMAND "Hello"
   4                 MESSAGE "Hello! "||CURRENT
   5             COMMAND "Quit"
(fgldb)
```
