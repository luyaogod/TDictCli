---
title: "Understanding the debugger"
source: "fgl-topics/c_fgl_Debugger_002.html"
breadcrumb: "Programming tools > Integrated debugger > Understanding the debugger"
type: "concept"
---

# Understanding the debugger

> This is an introduction to the integrated debugger.

The debugger is a feature built in to the runtime system (fglrun) that allows
you to control the execution of a program step by step, so that you can find logical and runtime
errors.

There are three debug modes possible with the Genero runtime system:

1. Start the fglrun program from the command line with the
   `-d` option. For more details, see [Starting fglrun in debug mode](2576-starting-fglrun-in-debug-mode.md "The runtime system can be started in debug mode with the -d option.").
2. Attaching with the fgldb tool, to an fglrun process running on the same
   machine, by using the process id. Note that with the graphical debugger of Genero Studio, a debug
   session with fgldb can be started on a remote server through ssh. For more
   details, see [Attaching to a running program](2577-attaching-to-a-running-program.md "It is possible to start the debugger for a program running on the same computer.").
3. Connect directly with the fgldb tool, to the debug TCP port of a runtime
   system running on a mobile device in standalone mode. For more details, see [Debugging on a mobile device](2579-debugging-on-a-mobile-device.md "It is possible to remotely start the debugger for an app running on a mobile device.").
4. Use Microsoft® Visual Studio Code
   editor and the VS Code extension for Genero BDL. For more details, see [Visual Studio Code extension](2545-visual-studio-code-extension.md).

The debugger supports a subset of the standard GNU C/C++ debugger called
gdb.

In command line mode, the debugger shows the following prompt

```
(fgldb)
```

A command is a single line of input. It starts with a command name, which may be followed by
arguments whose meaning depends on the command name. For example, the command `step`
accepts as an argument the number of times to step:

```
(fgldb) step 5
```

You can use command abbreviations. For example, the 'step' command abbreviation is 's':

```
(fgldb) s 5
```

Possible command abbreviations are shown in the command's syntax.

The debugger prompt provides command history and in-line editing capabilities, such as other
command-line tools using the GNU readline library, as well as command and expression completion with
TAB.

A blank line as input to the debugger (pressing just the RETURN or ENTER keys) usually causes
the previous command to repeat. However, commands whose unintentional repetition might cause
problems will not repeat in this way.
