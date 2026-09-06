---
title: "os.Path.baseName"
source: "fgl-topics/c_fgl_ext_os_path_basename.html"
breadcrumb: "Library reference > Extension packages > The os package > The os.Path class > os.Path methods > os.Path.baseName"
type: "concept"
---

# os.Path.baseName

> Returns the last element of a path.

## Syntax

```
os.Path.baseName(
     path STRING)
  RETURNS STRING
```

1. path is the name of the file.

## Usage

This method extracts the last component of a path provided as argument.

For example, if you pass `"/root/dir1/file.ext"` as the parameter,
it will return `"file.ext"`.

See [Example 1: Filename parts](3750-example-1-filename-parts.md) for more examples.
