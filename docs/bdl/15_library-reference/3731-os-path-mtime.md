---
title: "os.Path.mtime"
source: "fgl-topics/c_fgl_ext_os_path_mtime.html"
breadcrumb: "Library reference > Extension packages > The os package > The os.Path class > os.Path methods > os.Path.mtime"
type: "concept"
---

# os.Path.mtime

> Returns the time of the last file modification.

## Syntax

```
os.Path.mtime(
   path STRING)
  RETURNS STRING
```

1. path is the path to a file.

## Usage

The function returns a string containing the last modification time for the
specified file, in the standard format '`YYYY-MM-DD HH:MM:SS`'.

If the function fails, it returns `NULL`.

## Related links

**Related concepts**  

[os.Path.getModificationTime](3719-os-path-getmodificationtime.md "Returns the last time the file was modified.")
