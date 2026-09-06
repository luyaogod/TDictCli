---
title: "fgl_system()"
source: "fgl-topics/c_fgl_BuiltInFunctions_FGL_SYSTEM.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > fgl_system()"
type: "concept"
---

# fgl_system()

> Runs a command on the application server.

## Syntax

```
FUNCTION fgl_system(
   program STRING )
```

1. program is the command line to be executed on the server.

## Usage

The `fgl_system()` function suspends the execution of the program and
executes the command passed as parameter on the application server where fglrun is executed.
The command is executed in a new shell and the program is suspended until the command
terminates.

When running the program in TUI mode, the terminal is switched to
[line mode](../09_advanced-features/0830-run.md "The RUN instruction executes the command passed as argument.") before executing the command passed
to the `fgl_system()` function.

This function is provided for backward compatibility. In older versions, the function
could raise a terminal emulator on the front-end to show the command output on the
workstation. This feature is no longer supported.
