---
title: "os.Path.rename"
source: "fgl-topics/c_fgl_ext_os_path_rename.html"
breadcrumb: "Library reference > Extension packages > The os package > The os.Path class > os.Path methods > os.Path.rename"
type: "concept"
---

# os.Path.rename

> Renames a file or a directory.

## Syntax

```
os.Path.rename(
   oldPath STRING,
   newPath STRING )
  RETURNS INTEGER
```

1. oldPath is the current name of the file or
   directory to be renamed.
2. newPath is the new name to assign to the file
   or directory.

## Usage

The function returns `TRUE` if the file or directory has been
successfully renamed, `FALSE` otherwise.

On UNIX™ platforms, you can rename/move
files and directories located on the same file system.

On Microsoft™
Windows® platforms only files can be renamed/moved.
However, on Windows you can move files across disks and
directories.
