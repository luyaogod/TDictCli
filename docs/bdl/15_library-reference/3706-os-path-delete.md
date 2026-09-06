---
title: "os.Path.delete"
source: "fgl-topics/c_fgl_ext_os_path_delete.html"
breadcrumb: "Library reference > Extension packages > The os package > The os.Path class > os.Path methods > os.Path.delete"
type: "concept"
---

# os.Path.delete

> Deletes a file or a directory.

## Syntax

```
os.Path.delete(
   path STRING)
  RETURNS INTEGER
```

1. path is the path to the file or directory to delete.

## Usage

A directory can only be deleted if it is empty.

The function returns `TRUE` if the file has been successfully deleted,
`FALSE` otherwise.
