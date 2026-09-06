---
title: "os.Path.dirFMask"
source: "fgl-topics/c_fgl_ext_os_path_dirfmask.html"
breadcrumb: "Library reference > Extension packages > The os package > The os.Path class > os.Path methods > os.Path.dirFMask"
type: "concept"
---

# os.Path.dirFMask

> Defines a filter mask for os.Path.dirOpen().

## Syntax

```
os.Path.dirFMask(
   mask INTEGER)
```

1. mask defines the filter mask.

## Usage

When you call this function, you define the filter mask for any subsequent
[`os.Path.dirOpen()`](3712-os-path-diropen.md "Opens a directory and returns an integer handle to this directory.")
call.

By default, all kinds of directory entries are selected by the `dirOpen()`
function. You can restrict the number of entries by using a filter mask.

The parameter of the `os.Path.dirFMask()` function must be a combination
of the following bits:

- `0x01` = Exclude hidden files (.\*)
- `0x02` = Exclude directories
- `0x04` = Exclude symbolic links
- `0x08` = Exclude regular files

For example, to retrieve only regular files, you must
call:

```
CALL os.Path.dirFMask( 1 + 2 + 4 )
```
