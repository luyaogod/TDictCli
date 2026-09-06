---
title: "os.Path.getModificationTime"
source: "fgl-topics/c_fgl_ext_os_path_getmodificationtime.html"
breadcrumb: "Library reference > Extension packages > The os package > The os.Path class > os.Path methods > os.Path.getModificationTime"
type: "concept"
---

# os.Path.getModificationTime

> Returns the last time the file was modified.

## Syntax

```
os.Path.getModificationTime(
     path STRING)
  RETURNS DATETIME YEAR TO FRACTION(5)
```

1. path is the name of the file.

## Usage

This method returns the last modification time of a file as a `DATETIME`
value.

Depending on the platform, the fraction of seconds may not be supported. Linux,
Microsoft®
Windows® and Apple®
macOS™ support fraction of seconds.

If the function fails, it returns `NULL`.

## Related links

**Related concepts**  

[os.Path.setModificationTime](3743-os-path-setmodificationtime.md "Sets the modification time of a file.")

[os.Path.getAccessTime](3718-os-path-getaccesstime.md "Returns the last time the file was accessed.")

[os.Path.mtime](3731-os-path-mtime.md "Returns the time of the last file modification.")
