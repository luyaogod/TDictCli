---
title: "localStorage.removeItem"
source: "fgl-topics/c_fgl_frontcall_localstorage_removeitem.html"
breadcrumb: "Library reference > Built-in front calls > Local storage front calls > localStorage.removeItem"
type: "concept"
---

# localStorage.removeItem

> Deletes the specified local storage key.

## Syntax

```
ui.Interface.frontCall("localStorage", "removeItem",
   [key], [])
```

1. key is the name of the local storage key.

## Usage

The `removeItem` function deletes the specified local storage key.
