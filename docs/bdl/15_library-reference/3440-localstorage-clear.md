---
title: "localStorage.clear"
source: "fgl-topics/c_fgl_frontcall_localstorage_clear.html"
breadcrumb: "Library reference > Built-in front calls > Local storage front calls > localStorage.clear"
type: "concept"
---

# localStorage.clear

> Removes all local storage key/value pairs.

## Syntax

```
ui.Interface.frontCall("localStorage", "clear",
   [], [])
```

## Usage

The `clear` function removes all local storage keys currently
saved on the front-end side.
