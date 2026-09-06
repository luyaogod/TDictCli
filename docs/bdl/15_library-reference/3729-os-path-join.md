---
title: "os.Path.join"
source: "fgl-topics/c_fgl_ext_os_path_join.html"
breadcrumb: "Library reference > Extension packages > The os package > The os.Path class > os.Path methods > os.Path.join"
type: "concept"
---

# os.Path.join

> Joins two path segments adding the platform-dependent separator.

## Syntax

```
os.Path.join(
   begin STRING,
   end STRING )
  RETURNS STRING
```

1. begin is the beginning path segment.
2. end is the ending path segment.

## Usage

Use this method to construct a path with no system-specific code to use the
correct path separator:

```
LET path = os.Path.join(os.Path.homedir(), name)
```

This method returns the ending path segment if it is an absolute path.

If one of the arguments is `NULL`, the function returns
`NULL`.
