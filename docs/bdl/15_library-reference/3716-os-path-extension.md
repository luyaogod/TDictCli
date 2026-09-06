---
title: "os.Path.extension"
source: "fgl-topics/c_fgl_ext_os_path_extension.html"
breadcrumb: "Library reference > Extension packages > The os package > The os.Path class > os.Path methods > os.Path.extension"
type: "concept"
---

# os.Path.extension

> Returns the file extension.

## Syntax

```
os.Path.extension(
   path STRING)
  RETURNS  STRING
```

1. path is the path to a file.

## Usage

The function returns the string following the last dot found in path.

If path does not have an extension, the function returns
`NULL`.
