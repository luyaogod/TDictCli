---
title: "fgl_window_getoption()"
source: "fgl-topics/c_fgl_BuiltInFunctions_FGL_WINDOW_GETOPTION.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > fgl_window_getoption()"
type: "concept"
---

# fgl_window_getoption()

> Returns attributes of the current window.

## Syntax

```
FUNCTION fgl_window_getoption(
   name STRING )  
  RETURNS STRING
```

1. name is an window attribute name.

## Usage

The `fgl_window_getoption()` function returns the value of the window
attribute passed as parameter.

Possible parameters are: `name`, `x`, `y`,
`width`, `height`, `formline`,
`messageline`.

This function is provided for backward compatibility, do not use this function in
modern GUI applications.
