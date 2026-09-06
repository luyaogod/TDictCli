---
title: "os.Path.rwx"
source: "fgl-topics/c_fgl_ext_os_path_rwx.html"
breadcrumb: "Library reference > Extension packages > The os package > The os.Path class > os.Path methods > os.Path.rwx"
type: "concept"
description: "Returns the UNIX file permissions of a file."
---

# os.Path.rwx

> Returns the UNIX™ file permissions of a file.

## Syntax

```
os.Path.rwx(
   path STRING)
  RETURNS INTEGER
```

1. path is the path to a file or directory.

## Usage

This function returns the combination of permissions for user, group and other.

The mode is returned as an integer value which is the combination of read,
write and execution bits for the user, group and other part of the UNIX file permission. For example, if a file has the `-rwxr-xr-x` permissions,
the method returns integer 493: `((4+2+1)*64+(4+1)*8)+(4+1)) = 493`.

In case of failure, the returned value is -1.

On Microsoft™ Windows®, this method always returns -1.
