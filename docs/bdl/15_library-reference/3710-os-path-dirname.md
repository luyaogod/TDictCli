---
title: "os.Path.dirName"
source: "fgl-topics/c_fgl_ext_os_path_dirname.html"
breadcrumb: "Library reference > Extension packages > The os package > The os.Path class > os.Path methods > os.Path.dirName"
type: "concept"
---

# os.Path.dirName

> Returns all components of a path excluding the last one.

## Syntax

```
os.Path.dirName(
   path STRING)
  RETURNS STRING
```

1. path is the path to a file or directory.

## Usage

This method removes the last component of a path provided as argument.

For example, if you pass "/root/dir1/file.ext" as the parameter, it will
return "/root/dir1".

See [Example 1: Filename parts](3750-example-1-filename-parts.md) for more examples.
