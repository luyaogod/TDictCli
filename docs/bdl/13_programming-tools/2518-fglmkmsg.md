---
title: "fglmkmsg"
source: "fgl-topics/c_fgl_tools_fglmkmsg.html"
breadcrumb: "Programming tools > Command reference > fglmkmsg"
type: "concept"
---

# fglmkmsg

> The fglmkmsg tool compiles .msg message files into a binary version used by programs.

## Syntax

```
fglmkmsg [options] srcfile.msg [outfile.iem]
```

1. options are described in Table 1.
2. srcfile.msg is the source message file.
3. outfile.iem is the destination file.

## Options

| Option | Description |
| --- | --- |
| `-V` | Displays version information. |
| `-h` | Displays options for the tool. |
| `-r msgfile` | De-compiles a binary message file. |

## Usage

The fglmkmsg command line tool compiles a .msg message
file into a .iem compiled
version:

```
fglmkmsg mess01.msg
```

For
backward compatibility, you can specify the output file as second
argument:

```
fglmkmsg mess01.msg mess01.iem
```

The .iem compiled version can be used by BDL programs, for example, when
the `HELP` clause is used in a [`MENU`](../11_user-interface/1904-ring-menus-menu.md "The MENU instruction implements a list of options the end user can choose from.") or [`INPUT`](../11_user-interface/1930-record-input-input.md "The INPUT instruction provides single record input control in an application form.") instruction.

If case of error, the fglmkmsg command execution status is different from
zero. Consequently, it is possible to detect compilation errors in scripts and makefiles.

## Related links

**Related concepts**  

[Compiling message files (.msg)](../11_user-interface/1596-compiling-message-files-msg.md "The .msg message files must be compiled to .iem binary files, in order to be loaded by the runtime system.")

[Message files](../11_user-interface/1593-message-files.md "Message files centralize strings and larger texts identified by a number, that can be used in programs.")
