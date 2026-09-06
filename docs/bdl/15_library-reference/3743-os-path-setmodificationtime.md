---
title: "os.Path.setModificationTime"
source: "fgl-topics/c_fgl_ext_os_path_setmodificationtime.html"
breadcrumb: "Library reference > Extension packages > The os package > The os.Path class > os.Path methods > os.Path.setModificationTime"
type: "concept"
---

# os.Path.setModificationTime

> Sets the modification time of a file.

## Syntax

```
os.Path.setModificationTime(
     path STRING,
     time DATETIME YEAR TO FRACTION(5)
)
```

1. path is the name of the file.
2. time is the date/time to be set.

## Usage

This method sets the modification time of a file as a `DATETIME` value.

Depending on the platform, the fraction of seconds may not be supported. Linux,
Microsoft®
Windows® and Apple®
macOS™ support fraction of seconds.

## Related links

**Related concepts**  

[os.Path.getModificationTime](3719-os-path-getmodificationtime.md "Returns the last time the file was modified.")

[os.Path.setAccessTime](3742-os-path-setaccesstime.md "Sets the access time of a file.")
