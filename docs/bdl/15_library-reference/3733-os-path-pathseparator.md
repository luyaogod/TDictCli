---
title: "os.Path.pathSeparator"
source: "fgl-topics/c_fgl_ext_os_path_pathseparator.html"
breadcrumb: "Library reference > Extension packages > The os package > The os.Path class > os.Path methods > os.Path.pathSeparator"
type: "concept"
---

# os.Path.pathSeparator

> Returns the character used in environment variables to separate path elements.

## Syntax

```
os.Path.pathSeparator()
  RETURNS STRING
```

## Usage

You typically use this method to build a path from two components.

On UNIX™, the path separator is
'`:`'.

On Windows®, the path separator is
'`;`'.
