---
title: "os.Path.type"
source: "fgl-topics/c_fgl_ext_os_path_type.html"
breadcrumb: "Library reference > Extension packages > The os package > The os.Path class > os.Path methods > os.Path.type"
type: "concept"
---

# os.Path.type

> Returns the file type as a string.

## Syntax

```
os.Path.type(
   path STRING)
  RETURNS STRING
```

1. path is the path to file.

## Usage

On UNIX™, this method follows symbolic links.
Use the [`os.Path.islink()`](3726-os-path-islink.md "Checks if a file is an UNIX symbolic link.")
method to identify symbolic links.

The possible values returned by this method are:

1. `file`: the file is a regular file
2. `directory`: the file is a directory
3. `socket`: the file is a socket
4. `fifo`: the file is a fifo
5. `block`: the file is a block device
6. `char`: the file is a character device
