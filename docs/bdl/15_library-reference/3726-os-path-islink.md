---
title: "os.Path.isLink"
source: "fgl-topics/c_fgl_ext_os_path_islink.html"
breadcrumb: "Library reference > Extension packages > The os package > The os.Path class > os.Path methods > os.Path.isLink"
type: "concept"
description: "Checks if a file is an UNIX symbolic link."
---

# os.Path.isLink

> Checks if a file is an UNIX™ symbolic link.

## Syntax

```
os.Path.isLink(
   path STRING)
  RETURNS BOOLEAN
```

1. path is the file or directory path.

## Usage

The function returns `TRUE` if the files is a symbolic link,
`FALSE` otherwise.

This method can only be used on UNIX!
