---
title: "Example 2: Browsing directories"
source: "fgl-topics/c_fgl_ext_os_path_007.html"
breadcrumb: "Library reference > Extension packages > The os package > The os.Path class > Examples > Example 2: Browsing directories"
type: "concept"
description: "This program takes a directory path as an argument and scans the content recursively: IMPORT os MAIN CALL showDir(arg_val(1)) END MAIN FUNCTION showDir(path) DEFINE path STRING DEFINE child STRING ..."
---

# Example 2: Browsing directories

This program takes a directory path as an argument and scans the
content recursively:

```
IMPORT os

MAIN
  CALL showDir(arg_val(1))
END MAIN

FUNCTION showDir(path)
  DEFINE path STRING
  DEFINE child STRING
  DEFINE h INTEGER

  IF NOT os.Path.exists(path) THEN
     RETURN
  END IF

  IF NOT os.Path.isDirectory(path) THEN
     DISPLAY " ", os.Path.baseName(path)
     RETURN
  END IF

  DISPLAY "[", path, "]"
  CALL os.Path.dirSort("name", 1)
  LET h = os.Path.dirOpen(path)
  WHILE h > 0
     LET child = os.Path.dirNext(h)
     IF child IS NULL THEN EXIT WHILE END IF
     IF child == "." OR child == ".." THEN CONTINUE WHILE END IF
     CALL showDir( os.Path.join( path, child ) )
  END WHILE

  CALL os.Path.dirClose(h)

END FUNCTION
```
