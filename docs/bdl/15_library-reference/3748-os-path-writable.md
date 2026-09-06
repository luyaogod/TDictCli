---
title: "os.Path.writable"
source: "fgl-topics/c_fgl_ext_os_path_writable.html"
breadcrumb: "Library reference > Extension packages > The os package > The os.Path class > os.Path methods > os.Path.writable"
type: "concept"
---

# os.Path.writable

> Checks if a file is writable.

## Syntax

```
os.Path.writable(
   path STRING)
  RETURNS INTEGER
```

1. path is the path to a file or directory.

## Usage

The function returns `TRUE` if the file is writable, `FALSE`
otherwise.

If the file does not exist, or when the specified path is `NULL`, the method
returns `FALSE`.
