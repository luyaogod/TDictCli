---
title: "fgl_getversion()"
source: "fgl-topics/c_fgl_BuiltInFunctions_FGL_GETVERSION.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > fgl_getversion()"
type: "concept"
---

# fgl_getversion()

> Returns the product version number of Genero.

## Syntax

```
FUNCTION fgl_getversion()
  RETURNS STRING
```

## Usage

The `fgl_getversion()` function returns the product version number of the Genero
Business Development Language runtime system.

> **Important:**
>
> This function is provided for debugging only; do not write business code
> dependent on the build number. The format of the returned value is subject of change in future
> versions.
