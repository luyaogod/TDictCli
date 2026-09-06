---
title: "os.Path.dirOpen"
source: "fgl-topics/c_fgl_ext_os_path_diropen.html"
breadcrumb: "Library reference > Extension packages > The os package > The os.Path class > os.Path methods > os.Path.dirOpen"
type: "concept"
---

# os.Path.dirOpen

> Opens a directory and returns an integer handle to this directory.

## Syntax

```
os.Path.dirOpen(
   path STRING)
  RETURNS INTEGER
```

1. path is the name of the directory.
2. handle is the directory handle.

## Usage

This function creates a handle to scan the elements of a directory.

The function returns a value of 0 if it fails to open the directory.

Before calling the `dirOpen()` method, you can define a filter
with [`os.Path.dirFMask()`](3709-os-path-dirfmask.md "Defines a filter mask for os.Path.dirOpen()."),
and a sort order
with [`os.Path.dirSort()`](3713-os-path-dirsort.md "Defines the sort criteria and sort order for os.Path.dirOpen().").
