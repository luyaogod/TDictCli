---
title: "os.Path.chOwn"
source: "fgl-topics/c_fgl_ext_os_path_chown.html"
breadcrumb: "Library reference > Extension packages > The os package > The os.Path class > os.Path methods > os.Path.chOwn"
type: "concept"
description: "Changes the UNIX owner and group of a file."
---

# os.Path.chOwn

> Changes the UNIX™ owner and group of a file.

## Syntax

```
os.Path.chOwn(
   path STRING,
   uid INTEGER,
   gui INTEGER )
  RETURNS INTEGER
```

1. path is the name of the file.
2. uid is the user id.
3. gui is the group id.

## Usage

This method can only be used on UNIX!

Function returns `TRUE` on success, `FALSE` otherwise.
