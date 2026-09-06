---
title: "os.Path.mkDir"
source: "fgl-topics/c_fgl_ext_os_path_mkdir.html"
breadcrumb: "Library reference > Extension packages > The os package > The os.Path class > os.Path methods > os.Path.mkDir"
type: "concept"
---

# os.Path.mkDir

> Creates a new directory.

## Syntax

```
os.Path.mkDir(
   path STRING)
  RETURNS INTEGER
```

1. path is the path of the directory to create.

## Usage

The specified path must be an absolute or relative path with the new directory to be created, and
the parent directories must already exist.

The function returns `TRUE` if the directory has been successfully created,
`FALSE` otherwise.
