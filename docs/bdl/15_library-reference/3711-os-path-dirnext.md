---
title: "os.Path.dirNext"
source: "fgl-topics/c_fgl_ext_os_path_dirnext.html"
breadcrumb: "Library reference > Extension packages > The os package > The os.Path class > os.Path methods > os.Path.dirNext"
type: "concept"
---

# os.Path.dirNext

> Reads the next entry in the directory opened with os.Path.dirOpen().

## Syntax

```
os.Path.dirNext(
   dirHandle INTEGER)
  RETURNS STRING
```

1. dirHandle is the directory handle of the directory to read.

## Usage

This function returns the next entry of the directory opened with [`os.Path.dirOpen()`](3712-os-path-diropen.md "Opens a directory and returns an integer handle to this directory.").

Returns `NULL` if all entries have been read.
