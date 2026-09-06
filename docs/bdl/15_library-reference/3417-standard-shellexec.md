---
title: "standard.shellExec"
source: "fgl-topics/c_fgl_frontcall_standard_shellexec.html"
breadcrumb: "Library reference > Built-in front calls > Standard front calls > standard.shellExec"
type: "concept"
---

# standard.shellExec

> Opens a file on the front-end platform with the program associated to the file extension.

## Syntax

> **Important:**
>
> This feature is deprecated, its use is
> discouraged although not prohibited.

```
ui.Interface.frontCall("standard", "shellExec",
  [document, action], [result])
```

1. document - The document file to be opened.
2. action - (optional, Windows® Only!)
   The action to perform, related to the way the file type is registered in Windows Registry.
3. result - The execution status (`TRUE`=success,
   `FALSE`=error).

## Usage

The "`shellExec`" front call opens a file on the front-end platform with the
program associated to the file extension.

This front call is mainly designed for the Genero Desktop Client on Windows platforms.

> **Important:**
>
> The `standard.shellExec` front call is deprecated. In order to
> view a document (like a PDF for example), if the document can be displayed by web browsers, use the
> [standard.launchURL](3406-standard-launchurl.md "Opens a URL with the default URL handler of the front-end.") front call instead.

When specifying a file path, pay attention to platform specific rules regarding
directory separators and space characters in filenames. When the front-end executes on a recent Microsoft™ Windows
system, you can use the `/` slash character as directory separator, like on Unix
systems. A directory or filename can contain spaces, and there is no need to surround the path with
double quotes in such case. When using backslash directory separators, make sure to escape backslash
characters in string literals with `\\`.

Under X11 Systems, this uses xdg-open, which needs to be installed and
configured on your system. Kfmclient will be used as a workaround when
xdg-open is not available.
