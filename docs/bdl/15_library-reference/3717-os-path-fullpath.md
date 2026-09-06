---
title: "os.Path.fullPath"
source: "fgl-topics/c_fgl_ext_os_path_fullpath.html"
breadcrumb: "Library reference > Extension packages > The os package > The os.Path class > os.Path methods > os.Path.fullPath"
type: "concept"
---

# os.Path.fullPath

> Returns the canonical equivalent of a path.

## Syntax

```
os.Path.fullPath(
   path STRING)
  RETURNS STRING
```

1. path is the path to complete.

## Usage

The `os.path.fullPath()` class method takes a path as parameter and resolves extra
path separator characters (`/` on UNIX™,
`\` on Windows®), as well as references to
current (`.`) and parent directory (`..`). The result is called a
canonical path.

On UNIX, symbolic links are not
followed. Use the [`os.Path.isLink()`](3726-os-path-islink.md "Checks if a file is an UNIX symbolic link.") method to identify symbolic links.

## Example

```
DISPLAY os.Path.fullPath("/home/usr//scott/tmp/../images")
```

Resolves to:

```
/home/usr/scott/images
```
