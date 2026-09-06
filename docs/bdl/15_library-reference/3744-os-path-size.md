---
title: "os.Path.size"
source: "fgl-topics/c_fgl_ext_os_path_size.html"
breadcrumb: "Library reference > Extension packages > The os package > The os.Path class > os.Path methods > os.Path.size"
type: "concept"
---

# os.Path.size

> Returns the size of a file.

## Syntax

```
os.Path.size(
   path STRING)
  RETURNS BIGINT
```

1. path is the path to a file.

## Usage

The function returns the size in bytes for the specified file.

Large files sizes ( > 2GB ) are supported.
