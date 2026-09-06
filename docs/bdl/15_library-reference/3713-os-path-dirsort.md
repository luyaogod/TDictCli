---
title: "os.Path.dirSort"
source: "fgl-topics/c_fgl_ext_os_path_dirsort.html"
breadcrumb: "Library reference > Extension packages > The os package > The os.Path class > os.Path methods > os.Path.dirSort"
type: "concept"
---

# os.Path.dirSort

> Defines the sort criteria and sort order for os.Path.dirOpen().

## Syntax

```
os.Path.dirSort(
   criteria STRING,
   order INTEGER )
```

1. criteria is the sort criteria.
2. order defines ascending (1) or descending (-1) order.

## Usage

When you call this function, you define the sort criteria and sort order for any
subsequent [`os.Path.dirOpen()`](3712-os-path-diropen.md "Opens a directory and returns an integer handle to this directory.") call.

The criteria parameter must be one of the following strings:

- `"undefined"` = No sort. This is the default. Entries are read as returned by the OS functions.
- `"name"` = Sort by file name.
- `"size"` = Sort by file size.
- `"type"` = Sort by file type (directory, link, regular file).
- `"atime"` = Sort by access time.
- `"mtime"` = Sort by modification time.
- `"extension"` = Sort by file extension.

When sorting by name, directory entries will be ordered based on the [current locale](../09_advanced-features/0863-localization.md "Localization support allows you to implement programs that follow specific language and cultural rules.").

When sorting by any criteria other than the file name, entries having the
same value for the given criteria are ordered by name following the value
of the order parameter.
