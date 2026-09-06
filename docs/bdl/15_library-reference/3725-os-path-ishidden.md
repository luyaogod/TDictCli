---
title: "os.Path.isHidden"
source: "fgl-topics/c_fgl_ext_os_path_ishidden.html"
breadcrumb: "Library reference > Extension packages > The os package > The os.Path class > os.Path methods > os.Path.isHidden"
type: "concept"
---

# os.Path.isHidden

> Checks if a file is hidden.

## Syntax

```
os.Path.isHidden(
   path STRING)
  RETURNS BOOLEAN
```

1. path is the file or directory path.

## Usage

The function returns `TRUE` if the file is hidden, `FALSE`
otherwise.

For example, on UNIX™, files starting with a dot in
the filename are considered as hidden when using the ls command.
