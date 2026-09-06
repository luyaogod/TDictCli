---
title: "fgl_dialog_getbufferstart()"
source: "fgl-topics/c_fgl_BuiltInFunctions_FGL_DIALOG_GETBUFFERSTART.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > fgl_dialog_getbufferstart()"
type: "concept"
---

# fgl_dialog_getbufferstart()

> Returns the row offset of the page to feed a paged display array.

## Syntax

```
FUNCTION fgl_dialog_getbufferstart()
  RETURNS INTEGER
```

## Usage

The `FGL_DIALOG_GETBUFFERSTART()` function returns the
record list offset to be used to fill a page of a `DISPLAY
ARRAY` running in [paged mode](../11_user-interface/1959-record-list-display-array.md "The DISPLAY ARRAY instruction provides record list navigation in an application form, with optional record modification actions.").

This function must be called in the context of the `ON FILL
BUFFER` trigger. The returned value is undefined if the function
is used outside this trigger.
