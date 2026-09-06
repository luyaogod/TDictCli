---
title: "os.Path.gid"
source: "fgl-topics/c_fgl_ext_os_path_gid.html"
breadcrumb: "Library reference > Extension packages > The os package > The os.Path class > os.Path methods > os.Path.gid"
type: "concept"
description: "Returns the UNIX group id of a file."
---

# os.Path.gid

> Returns the UNIX™ group id of a file.

## Syntax

```
os.Path.gid(
   path STRING)
  RETURNS INTEGER
```

1. path is the path to a file.

## Usage

This function returns the group id of the file.

This method can only be used on UNIX!

Function returns -1 if it fails to get the user id.
