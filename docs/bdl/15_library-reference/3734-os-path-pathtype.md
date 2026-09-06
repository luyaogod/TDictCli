---
title: "os.Path.pathType"
source: "fgl-topics/c_fgl_ext_os_path_pathtype.html"
breadcrumb: "Library reference > Extension packages > The os package > The os.Path class > os.Path methods > os.Path.pathType"
type: "concept"
---

# os.Path.pathType

> Checks if a path is a relative path or an absolute path.

## Syntax

```
os.Path.pathType(
   path STRING)
  RETURNS STRING
```

1. path is the path to check.

## Usage

The function returns "`absolute`" if the path is an absolute path, or
"`relative`" if the path is a relative path.

> **Note:**
>
> On Unix-style platforms, an absolute path starts with a ( `/` ) slash. On Windows® platforms, an absolute path starts
> with a backslash ( `\` ), a slash ( `/` ), or starts with a drive
> letter followed by a colon ( `C:` )

If the path is `NULL`, the function returns `NULL`.
