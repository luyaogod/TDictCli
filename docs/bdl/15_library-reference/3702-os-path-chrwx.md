---
title: "os.Path.chRwx"
source: "fgl-topics/c_fgl_ext_os_path_chrwx.html"
breadcrumb: "Library reference > Extension packages > The os package > The os.Path class > os.Path methods > os.Path.chRwx"
type: "concept"
description: "Changes the UNIX permissions of a file."
---

# os.Path.chRwx

> Changes the UNIX™ permissions of a file.

## Syntax

```
os.Path.chRwx(
   path STRING,
   mode INTEGER )
  RETURNS INTEGER
```

1. path is the name of the file.
2. mode is the UNIX permission combination
   as integer (not octal!).

## Usage

This method can only be used on UNIX!

Function returns `TRUE` on success, `FALSE` otherwise.

The mode must be an integer value which is the combination of read, write and
execution bits for the user, group and other part of the UNIX
file permission. Pass the mode as the integer version of permissions, not as
octal (the chrwx UNIX command takes an octal value as
parameter). For example, to set `-rw-r--r--` permissions, use integer 420:
`(((4+2)*64)+(4*8)+4) = 420`.
