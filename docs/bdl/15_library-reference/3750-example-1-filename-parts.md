---
title: "Example 1: Filename parts"
source: "fgl-topics/c_fgl_ext_os_path_006.html"
breadcrumb: "Library reference > Extension packages > The os package > The os.Path class > Examples > Example 1: Filename parts"
type: "concept"
description: "This program uses the file functions to extract the directory name, the base name, the root name, and the file extension: IMPORT os MAIN DISPLAY \"Dir name = \", os.Path.dirName(arg_val(1)) DISPLAY ..."
---

# Example 1: Filename parts

This program uses the file functions to extract the directory name,
the base name, the root name, and the file extension:

```
IMPORT os
MAIN
  DISPLAY "Dir name   = ", os.Path.dirName(arg_val(1))
  DISPLAY "Base name  = ", os.Path.baseName(arg_val(1))
  DISPLAY "Root name  = ", os.Path.rootName(arg_val(1))
  DISPLAY "Extension  = ", os.Path.extension(arg_val(1))
END MAIN
```

Example results:

```
Path                 dirName        baseName       rootName          extension
-------------------------------------------------------------------------------
.                    .              .              .                 NULL
..                   .              ..             .                 NULL
/                    /              /              /                 NULL
/usr/lib             /usr           lib            /usr/lib          NULL
/usr/                /              usr            /usr/             NULL
usr                  .              usr            usr               NULL
file.xx              .              file.xx        file              xx
/tmp.yy/file.xx      /tmp.yy        file.xx        /tmp.yy/file      xx
/tmp.yy/file.xx.yy   /tmp.yy        file.xx.yy     /tmp.yy/file.xx   yy
/tmp.yy/             /              tmp.yy         /tmp.yy/          NULL
/tmp.yy/.            /tmp.yy        .              /tmp.yy/          NULL
```

These examples use UNIX™ filenames. On Windows® the result would be different, as the filename separator is a
backslash (`\`).
