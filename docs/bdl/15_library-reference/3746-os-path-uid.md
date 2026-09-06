---
title: "os.Path.uid"
source: "fgl-topics/c_fgl_ext_os_path_uid.html"
breadcrumb: "Library reference > Extension packages > The os package > The os.Path class > os.Path methods > os.Path.uid"
type: "concept"
description: "Returns the UNIX user id of a file."
---

# os.Path.uid

> Returns the UNIX™ user id of a file.

## Syntax

```
os.Path.uid(
   path STRING)
  RETURNS INTEGER
```

1. path is the path to a file or directory.

## Usage

This method can only be used on UNIX!

Function returns -1 if it fails to get the user id.
