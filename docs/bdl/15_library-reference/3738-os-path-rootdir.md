---
title: "os.Path.rootDir"
source: "fgl-topics/c_fgl_ext_os_path_rootdir.html"
breadcrumb: "Library reference > Extension packages > The os package > The os.Path class > os.Path methods > os.Path.rootDir"
type: "concept"
---

# os.Path.rootDir

> Returns the root directory of the current working path.

## Syntax

```
os.Path.rootDir()
  RETURNS STRING
```

## Usage

This function returns the root directory of the current working path.

On UNIX™, it always returns
"`/`".

On Windows® it returns the current working drive as
"`[a-zA-Z]:\`"
