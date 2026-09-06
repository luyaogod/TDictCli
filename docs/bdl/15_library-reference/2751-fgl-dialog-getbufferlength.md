---
title: "fgl_dialog_getbufferlength()"
source: "fgl-topics/c_fgl_BuiltInFunctions_FGL_DIALOG_GETBUFFERLENGTH.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > fgl_dialog_getbufferlength()"
type: "concept"
---

# fgl_dialog_getbufferlength()

> Returns the number of rows to feed a paged DISPLAY ARRAY.

## Syntax

```
FUNCTION fgl_dialog_getbufferlength()
  RETURNS INTEGER
```

## Usage

The `fgl_dialog_getbufferlength()` function returns the number
of rows to be provided by the program to fill a page of a `DISPLAY
ARRAY` running in [paged mode](../11_user-interface/1959-record-list-display-array.md "The DISPLAY ARRAY instruction provides record list navigation in an application form, with optional record modification actions.").

This function must be called in the context of the `ON FILL BUFFER`
trigger. The returned value is undefined if the function is used outside this
trigger.
