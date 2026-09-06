---
title: "os.Path.dirClose"
source: "fgl-topics/c_fgl_ext_os_path_dirclose.html"
breadcrumb: "Library reference > Extension packages > The os package > The os.Path class > os.Path methods > os.Path.dirClose"
type: "concept"
---

# os.Path.dirClose

> Closes the directory referenced by the directory opened by os.Path.diropen().

## Syntax

```
os.Path.dirClose(
   dirHandle INTEGER)
```

1. dirHandle is the directory handle of the directory to close.

## Usage

This function closes the directory search handle opened with [`os.Path.dirOpen()`](3712-os-path-diropen.md "Opens a directory and returns an integer handle to this directory.").
