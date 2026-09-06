---
title: "Compiling message files (.msg)"
source: "fgl-topics/c_fgl_message_files_004.html"
breadcrumb: "User interface > Form definitions > Message files > Compiling message files (.msg)"
type: "concept"
---

# Compiling message files (.msg)

> The .msg message files must be compiled to .iem binary files, in order to be loaded by the runtime system.

In order to use message files in a program, the message source files (with
.msg extension) must be compiled with the [fglmkmsg](../13_programming-tools/2518-fglmkmsg.md "The fglmkmsg tool compiles .msg message files into a binary version used by programs.") utility to produce compiled message files (with
.iem extension).

The following command line compiles the message source file
mess01.msg:

```
fglmkmsg mess01.msg
```

This creates the compiled message file mess01.iem.

For backward compatibility, you can specify the output file as second argument:

```
fglmkmsg mess01.msg mess01.iem
```

The .iem compiled version of the message file must be distributed on the
machine where the programs are executed.
