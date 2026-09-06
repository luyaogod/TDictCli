---
title: "standard.execute"
source: "fgl-topics/c_fgl_frontcall_standard_execute.html"
breadcrumb: "Library reference > Built-in front calls > Standard front calls > standard.execute"
type: "concept"
---

# standard.execute

> Executes a command on the front-end platform, with or without waiting.

## Syntax

```
ui.Interface.frontCall("standard", "execute",
 [cmd,wait], [result])
```

1. command - The command to be executed.
2. wait - The wait option (`TRUE`=wait, `FALSE`=do
   not wait).
3. result - The execution status (`TRUE`=success,
   `FALSE`=error).

## Usage

The "`execute`" front call runs a command on the front-end platform, with or
without waiting option.

This front call is only supported with the desktop front-end (GDC). When using a browser and the
GAS, due to browser security limitation, it is not possible to start a process on the end user
workstation.

If the second parameter is set to `TRUE`, the runtime system will wait until
the front-end gives the control back after the local command was executed.

With the front-end running on Microsoft™ Windows™ systems platforms, depending on the program to be
executed, you may need to add the "cmd /C" command before the executable name and
program arguments.

With the front-end running on a Unix-like platform, shell scripts need to be preceeded with the
shell program (sh or bash).

OS commands return an exit status where zero means a success, while a non-zero exit status
indicates a failure, that typically reflects the system error it encountered. When the command
produces a zero exit status, the `standard.execute` front call will return
`TRUE/1`. When the command returns a non-zero exit status, the front call will return
`FALSE/0`. This is only valid when the front call waits for the end of the command
execution.

When specifying a file path, pay attention to platform specific rules regarding
directory separators and space characters in filenames. When the front-end executes on a recent Microsoft Windows
system, you can use the `/` slash character as directory separator, like on Unix
systems. A directory or filename can contain spaces, and there is no need to surround the path with
double quotes in such case. When using backslash directory separators, make sure to escape backslash
characters in string literals with `\\`.

## Example

For Microsoft Windows system:

```
-- Using backslash as directory separator:
CALL ui.Interface.frontCall("standard", "execute",
     ["C:\\Program Files\\FourJS\\gdc\\bin\\gdc.exe --help",FALSE], [res] )

-- Using slash as directory separator:
CALL ui.Interface.frontCall("standard", "execute",
     ["C:/Program Files/FourJS/gdc/bin/gdc.exe --help",FALSE], [res] )
```

## Related links

**Related concepts**  

[android.startActivity](3467-android-startactivity.md "Starts an external Android application (activity), and returns to the GMA application immediately.")

[android.startActivityForResult](3468-android-startactivityforresult.md "Starts an external application (Android activity) and waits until the activity is closed.")
