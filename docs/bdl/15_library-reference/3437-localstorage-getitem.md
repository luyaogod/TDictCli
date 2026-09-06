---
title: "localStorage.getItem"
source: "fgl-topics/c_fgl_frontcall_localstorage_getitem.html"
breadcrumb: "Library reference > Built-in front calls > Local storage front calls > localStorage.getItem"
type: "concept"
---

# localStorage.getItem

> Returns the current value of local storage key.

## Syntax

```
ui.Interface.frontCall("localStorage", "getItem",
   [key], [value])
```

1. key is the name of the local storage key.
2. value is the current value of the named key.

## Usage

The `getItem` function returns the value of the specified
local storage key.
