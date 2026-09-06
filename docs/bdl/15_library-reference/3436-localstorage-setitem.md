---
title: "localStorage.setItem"
source: "fgl-topics/c_fgl_frontcall_localstorage_setitem.html"
breadcrumb: "Library reference > Built-in front calls > Local storage front calls > localStorage.setItem"
type: "concept"
---

# localStorage.setItem

> Sets a value for local storage key.

## Syntax

```
ui.Interface.frontCall("localStorage", "setItem",
   [key,value], [])
```

1. key is the name of the local storage key.
2. value is the value to set for the named key.

## Usage

The `setItem` function sets the specified local storage key with
the value passed as parameter.
