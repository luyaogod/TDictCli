---
title: "os.Path.chVolume"
source: "fgl-topics/c_fgl_ext_os_path_chvolume.html"
breadcrumb: "Library reference > Extension packages > The os package > The os.Path class > os.Path methods > os.Path.chVolume"
type: "concept"
---

# os.Path.chVolume

> Changes the current working volume.

## Syntax

```
os.Path.chVolume(
   volume STRING)
  RETURNS INTEGER
```

1. volume is the name of the volume to select as the new current working
   volume.

## Usage

To change the current volume to C:

```
LET result = os.Path.chVolume("C:\\")
```

The function returns `TRUE` if the current working volume is successfully changed,
`FALSE` otherwise.
