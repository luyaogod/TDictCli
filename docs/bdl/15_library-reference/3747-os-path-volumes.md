---
title: "os.Path.volumes"
source: "fgl-topics/c_fgl_ext_os_path_volumes.html"
breadcrumb: "Library reference > Extension packages > The os package > The os.Path class > os.Path methods > os.Path.volumes"
type: "concept"
---

# os.Path.volumes

> Returns the available volumes.

## Syntax

```
os.Path.volumes()
  RETURNS STRING
```

## Usage

The function returns the list of all available volumes separated by
"`|`".

To display the list of available volumes, a volume is identified by its letter,
followed by a colon and a backslash (`:\`).

```
DISPLAY os.Path.volumes()
```

Output
example:

```
C:\|E:\|F:\
```
