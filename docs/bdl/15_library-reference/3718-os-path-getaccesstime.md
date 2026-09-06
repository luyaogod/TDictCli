---
title: "os.Path.getAccessTime"
source: "fgl-topics/c_fgl_ext_os_path_getaccesstime.html"
breadcrumb: "Library reference > Extension packages > The os package > The os.Path class > os.Path methods > os.Path.getAccessTime"
type: "concept"
---

# os.Path.getAccessTime

> Returns the last time the file was accessed.

## Syntax

```
os.Path.getAccessTime(
     path STRING)
  RETURNS DATETIME YEAR TO FRACTION(5)
```

1. path is the name of the file.

## Usage

This method returns the last access time of a file as a `DATETIME` value.

Depending on the platform, the fraction of seconds may not be supported. Linux,
Microsoft®
Windows® and Apple®
macOS™ support fraction of seconds.

If the function fails, it returns `NULL`.

> **Important:**
>
> The access time of a file may not be totally reliable, depending on the platform OS type and
> version, and file system type. Check [Wikipedia article about criticism of the atime of a file](https://en.wikipedia.org/wiki/Stat_%28system_call%29#Criticism_of_atime).

## Related links

**Related concepts**  

[os.Path.setAccessTime](3742-os-path-setaccesstime.md "Sets the access time of a file.")

[os.Path.getModificationTime](3719-os-path-getmodificationtime.md "Returns the last time the file was modified.")

[os.Path.atime](3699-os-path-atime.md "Returns the time of the last file access.")
