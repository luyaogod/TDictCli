---
title: "os.Path.rootName"
source: "fgl-topics/c_fgl_ext_os_path_rootname.html"
breadcrumb: "Library reference > Extension packages > The os package > The os.Path class > os.Path methods > os.Path.rootName"
type: "concept"
---

# os.Path.rootName

> Returns the file path without the file extension of the last element of the file path.

## Syntax

```
os.Path.rootName(
   path STRING)
  RETURNS STRING
```

1. path is the path to a file or directory.

## Usage

This method removes the file extension from the path provided as parameter.

For example, if you pass "/root/dir1/file.ext" as the parameter it will
return "/root/dir1/file".

See [Example 1: Filename parts](3750-example-1-filename-parts.md) for more examples.
