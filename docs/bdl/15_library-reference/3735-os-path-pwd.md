---
title: "os.Path.pwd"
source: "fgl-topics/c_fgl_ext_os_path_pwd.html"
breadcrumb: "Library reference > Extension packages > The os package > The os.Path class > os.Path methods > os.Path.pwd"
type: "concept"
---

# os.Path.pwd

> Returns the current working directory.

## Syntax

```
os.Path.pwd()
  RETURNS STRING
```

## Usage

This function returns the path of the current working directory.

On a mobile device, this front call returns the current application working directory:

- On Android™, it returns the directory where the program executes.
- On iOS, it returns the "Documents" directory under the application directory.
