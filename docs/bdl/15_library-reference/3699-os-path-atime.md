---
title: "os.Path.atime"
source: "fgl-topics/c_fgl_ext_os_path_atime.html"
breadcrumb: "Library reference > Extension packages > The os package > The os.Path class > os.Path methods > os.Path.atime"
type: "concept"
---

# os.Path.atime

> Returns the time of the last file access.

## Syntax

```
os.Path.atime(
   path STRING)
  RETURNS STRING
```

1. path is the path to a file.

## Usage

The function returns a string containing the last access time for the specified
file, in the standard format '`YYYY-MM-DD HH:MM:SS`'.

If the function fails, it returns `NULL`.

> **Important:**
>
> The access time of a file may not be totally reliable, depending on the platform OS type and
> version, and file system type. Check [Wikipedia article about criticism of the atime of a file](https://en.wikipedia.org/wiki/Stat_%28system_call%29#Criticism_of_atime).

## Related links

**Related concepts**  

[os.Path.getAccessTime](3718-os-path-getaccesstime.md "Returns the last time the file was accessed.")

[os.Path.getModificationTime](3719-os-path-getmodificationtime.md "Returns the last time the file was modified.")
