---
title: "os.Path.setAccessTime"
source: "fgl-topics/c_fgl_ext_os_path_setaccesstime.html"
breadcrumb: "Library reference > Extension packages > The os package > The os.Path class > os.Path methods > os.Path.setAccessTime"
type: "concept"
---

# os.Path.setAccessTime

> Sets the access time of a file.

## Syntax

```
os.Path.setAccessTime(
     path STRING,
     time DATETIME YEAR TO FRACTION(5)
)
```

1. path is the name of the file.
2. time is the date/time to be set.

## Usage

This method sets the access time of a file as a `DATETIME` value.

Depending on the platform, the fraction of seconds may not be supported. Linux,
Microsoft®
Windows® and Apple®
macOS™ support fraction of seconds.

> **Important:**
>
> The access time of a file may not be totally reliable, depending on the platform OS type and
> version, and file system type. Check [Wikipedia article about criticism of the atime of a file](https://en.wikipedia.org/wiki/Stat_%28system_call%29#Criticism_of_atime).

## Related links

**Related concepts**  

[os.Path.getAccessTime](3718-os-path-getaccesstime.md "Returns the last time the file was accessed.")

[os.Path.setModificationTime](3743-os-path-setmodificationtime.md "Sets the modification time of a file.")
