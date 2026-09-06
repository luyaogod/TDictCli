---
title: "os.Path.isRoot"
source: "fgl-topics/c_fgl_ext_os_path_isroot.html"
breadcrumb: "Library reference > Extension packages > The os package > The os.Path class > os.Path methods > os.Path.isRoot"
type: "concept"
---

# os.Path.isRoot

> Checks if a file path is a root path.

## Syntax

```
os.Path.isRoot(
   path STRING)
  RETURNS BOOLEAN
```

1. path is the path to check.

## Usage

The function returns `TRUE` if the path is a root path,
`FALSE` otherwise.

On UNIX™ the root path is '`/`'.

On Windows® the root path matches
"`[a-zA-Z]:\`".
