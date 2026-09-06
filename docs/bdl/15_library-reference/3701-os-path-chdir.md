---
title: "os.Path.chDir"
source: "fgl-topics/c_fgl_ext_os_path_chdir.html"
breadcrumb: "Library reference > Extension packages > The os package > The os.Path class > os.Path methods > os.Path.chDir"
type: "concept"
---

# os.Path.chDir

> Changes the current working directory.

## Syntax

```
os.Path.chDir(
   path STRING)
  RETURNS INTEGER
```

1. path is the path to the directory to select.

## Usage

Use this function to change the current working directory.

The function returns `TRUE` if the current directory is successfully selected,
`FALSE` otherwise.
