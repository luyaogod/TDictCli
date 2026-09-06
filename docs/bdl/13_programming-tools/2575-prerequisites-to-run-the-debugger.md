---
title: "Prerequisites to run the debugger"
source: "fgl-topics/c_fgl_Debugger_004.html"
breadcrumb: "Programming tools > Integrated debugger > Prerequisites to run the debugger"
type: "concept"
---

# Prerequisites to run the debugger

> Some requirements are needed before running the integrated debugger.

## FGLPROFILE `fglrun.ignoreDebuggerEvent` option

By default, fglrun processes can be stopped with the `SIGTRAP`
signal on UNIX-like systems, or with the `CTRL_BREAK_EVENT` signal on Microsoft™ Windows® systems.
These signals are used by tools such as [fgldb](2577-attaching-to-a-running-program.md "It is possible to start the debugger for a program running on the same computer."), to attach to a running program for a debug session.

The [FGLPROFILE](../07_configuration/0483-the-fglprofile-file-s.md "FGLPROFILE environment variable defines Genero BDL configuration files") configuration
option `fglrun.ignoreDebuggerEvent` allows you to disable the detection of the
debugger attachment signal. It is good practice to disable this feature in production environments,
by setting that option to
`true`:

```
fglrun.ignoreDebuggerEvent = true
```

Disabling the detection of the debugger attachment signal does not prevent to start the runtime
system directly in debug mode, with [fglrun
-d](2576-starting-fglrun-in-debug-mode.md "The runtime system can be started in debug mode with the -d option.").

## Source file search path (FGLSOURCEPATH)

Before starting the debugger, make sure you have properly set the [FGLLDPATH](../07_configuration/0528-fglldpath.md "Defines a list of paths to find program modules.") and [FGLSOURCEPATH](../07_configuration/0533-fglsourcepath.md "Defines a list of paths to program source files.") environment variable to let the
debugger find the source files.

To find source files, the debugger will search in the directories defined by FGLLDPATH. If the
source file is not found, the search continues in the directories defined by FGLSOURCEPATH.

The FGLSOURCEPATH environment variable is provided to distinguish execution directories
(containing .42m files) from source directories (containing
.4gl files), when the sources are not located in the same directory as the
pcode files.

UNIX™ example:

```
$ FGLSOURCEPATH="/usr/app/source:/home/scott/sources"
$ export FGLSOURCEPATH
```

Windows
example:

```
C:\> set FGLSOURCEPATH=C:\app\sources;C:\scott\sources
```

## Starting the debug-server

If the fglrun process to debug runs on a different computer, or if it has been
started by a user which is not the current development user, you need to start the debug-server (as
a debug proxy), to let the debugger attach to the running fglrun process.

For more details, read [Using the debug-server](2585-using-the-debug-server.md "The Genero BDL debug-server is a proxy for fglrun processes which can not be accessed directly by the debugger.").
